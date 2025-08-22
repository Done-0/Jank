package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/hashicorp/go-plugin"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/pkg/plugin/consts"

	pluginUtils "github.com/Done-0/jank/internal/utils/plugin"
	jank "github.com/Done-0/jank/pkg/plugin"
)

// PluginDiscoveryInfo 插件发现信息
type PluginDiscoveryInfo struct {
	*PluginInfo         // 嵌入插件基本信息
	Path         string `json:"path"`          // 插件路径
	IsRegistered bool   `json:"is_registered"` // 是否已注册
}

// PluginManagerImpl 插件管理器实现
type PluginManagerImpl struct {
	plugins map[string]*plugin.Client // 插件客户端映射
	infos   map[string]*PluginInfo    // 插件信息映射
	mu      sync.RWMutex              // 并发安全锁
}

// NewPluginManager 创建插件管理器实例
func NewPluginManager() *PluginManagerImpl {
	return &PluginManagerImpl{
		plugins: make(map[string]*plugin.Client),
		infos:   make(map[string]*PluginInfo),
	}
}

// RegisterPlugin 注册并启动插件
func (m *PluginManagerImpl) RegisterPlugin(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.infos[id]; exists {
		return fmt.Errorf("plugin %s already registered", id)
	}

	// 从配置文件加载插件信息
	cfgs, err := configs.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to get config: %w", err)
	}

	// 扫描所有目录查找匹配的插件ID
	var pluginPath string
	var configData []byte

	entries, err := os.ReadDir(cfgs.PluginConfig.PluginDir)
	if err != nil {
		return fmt.Errorf("failed to scan plugin directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		dirPath := filepath.Join(cfgs.PluginConfig.PluginDir, entry.Name())
		configFile := filepath.Join(dirPath, cfgs.PluginConfig.PluginConfigFile)

		if data, err := os.ReadFile(configFile); err == nil {
			var tempInfo PluginInfo
			if err := json.Unmarshal(data, &tempInfo); err == nil && tempInfo.ID == id {
				pluginPath = dirPath
				configData = data
				break
			}
		}
	}

	if pluginPath == "" {
		return fmt.Errorf("plugin %s not found", id)
	}

	var info PluginInfo
	if err := json.Unmarshal(configData, &info); err != nil {
		return fmt.Errorf("invalid plugin config for %s: %w", id, err)
	}

	// 设置插件工作目录和执行路径
	binaryPath := filepath.Join(pluginPath, info.Binary)

	// 转换为绝对路径，确保路径正确
	absBinaryPath, err := filepath.Abs(binaryPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for plugin binary: %w", err)
	}

	absPluginPath, err := filepath.Abs(pluginPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for plugin directory: %w", err)
	}

	cmd := exec.Command(absBinaryPath)
	cmd.Dir = absPluginPath

	// 创建插件客户端配置
	config := &plugin.ClientConfig{
		HandshakeConfig:  jank.HandshakeConfig,
		Plugins:          jank.PluginMap,
		Cmd:              cmd,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		AutoMTLS:         info.AutoMTLS,
		Managed:          info.Managed,
		StartTimeout:     time.Duration(info.StartTimeout) * time.Millisecond,
		MinPort:          info.MinPort,
		MaxPort:          info.MaxPort,
	}

	client := plugin.NewClient(config)
	if _, err := client.Start(); err != nil {
		client.Kill()
		return fmt.Errorf("failed to start plugin %s: %v", info.ID, err)
	}

	// 更新运行时状态
	info.Status = consts.PluginStatusRunning
	info.StartedAt = time.Now().Unix()
	m.refreshPluginInfo(&info, client)

	// 保存到内存映射
	m.infos[info.ID] = &info
	m.plugins[info.ID] = client

	global.SysLog.Infof("Plugin registered: %s (%s v%s) from %s, PID: %d, Binary: %s, Type: %s, Status: %s",
		info.ID, info.Name, info.Version, info.Repository, info.ProcessPID, info.Binary, info.Type, info.Status)

	return nil
}

// UnregisterPlugin 注销并停止插件
func (m *PluginManagerImpl) UnregisterPlugin(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	client, exists := m.plugins[id]
	if !exists {
		return fmt.Errorf("plugin %s not found", id)
	}

	client.Kill()

	if info, ok := m.infos[id]; ok {
		info.Status = consts.PluginStatusStopped
		global.SysLog.Infof("Plugin unregistered: %s (%s v%s)",
			id, info.Name, info.Version)
	}

	delete(m.plugins, id)
	delete(m.infos, id)

	return nil
}

// ExecutePlugin 执行插件方法
func (m *PluginManagerImpl) ExecutePlugin(ctx context.Context, id, method string, args map[string]any) (map[string]any, error) {
	m.mu.RLock()
	client, exists := m.plugins[id]
	info, infoExists := m.infos[id]
	m.mu.RUnlock()

	if !exists || !infoExists {
		return nil, fmt.Errorf("plugin %s not found", id)
	}

	// 获取 RPC 客户端并执行
	rpcClient, err := client.Client()
	if err != nil {
		m.mu.Lock()
		info.Status = consts.PluginStatusError
		m.refreshPluginInfo(info, client)
		m.mu.Unlock()
		return nil, err
	}

	raw, err := rpcClient.Dispense(info.Type)
	if err != nil {
		m.mu.Lock()
		info.Status = consts.PluginStatusError
		m.refreshPluginInfo(info, client)
		m.mu.Unlock()
		return nil, err
	}

	// 执行插件方法
	result, err := raw.(jank.Plugin).Execute(ctx, method, args)

	// 更新执行结果状态
	m.mu.Lock()
	if err != nil {
		info.Status = consts.PluginStatusError
	} else {
		// 执行成功，插件保持运行状态
		info.Status = consts.PluginStatusRunning
	}
	m.refreshPluginInfo(info, client)
	m.mu.Unlock()

	return result, err
}

// GetPlugin 获取插件信息
func (m *PluginManagerImpl) GetPlugin(id string) (*PluginInfo, error) {
	m.mu.RLock()
	info, exists := m.infos[id]
	client, clientExists := m.plugins[id]
	m.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("plugin %s not found", id)
	}

	// 刷新运行时信息
	if clientExists {
		m.mu.Lock()
		m.refreshPluginInfo(info, client)
		m.mu.Unlock()
	}

	return m.createPluginCopy(info), nil
}

// ListPlugins 列出所有插件（包括已注册和未注册的）
func (m *PluginManagerImpl) ListPlugins() ([]*PluginDiscoveryInfo, error) {
	cfgs, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	entries, err := os.ReadDir(cfgs.PluginConfig.PluginDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read plugin directory: %w", err)
	}

	var discoveredPlugins []*PluginDiscoveryInfo

	// 获取已注册插件的ID列表
	m.mu.RLock()
	registeredPluginIDs := make(map[string]bool)
	registeredPluginsMap := make(map[string]*PluginInfo)
	for id, info := range m.infos {
		registeredPluginIDs[id] = true
		// 刷新运行时信息
		if client, exists := m.plugins[id]; exists {
			m.refreshPluginInfo(info, client)
		}
		registeredPluginsMap[id] = info
	}
	m.mu.RUnlock()

	// 遍历所有插件目录
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		pluginPath := filepath.Join(cfgs.PluginConfig.PluginDir, entry.Name())
		configFile := filepath.Join(pluginPath, cfgs.PluginConfig.PluginConfigFile)

		configData, err := os.ReadFile(configFile)
		if err != nil {
			global.SysLog.Warnf("Skipping %s: cannot read %s", entry.Name(), cfgs.PluginConfig.PluginConfigFile)
			continue
		}

		var config PluginInfo
		if err := json.Unmarshal(configData, &config); err != nil {
			global.SysLog.Warnf("Skipping %s: invalid %s", entry.Name(), cfgs.PluginConfig.PluginConfigFile)
			continue
		}

		// 检查插件状态和运行时信息
		isRegistered := registeredPluginIDs[config.ID]
		var finalInfo *PluginInfo

		if isRegistered {
			// 使用已注册插件的运行时信息
			finalInfo = m.createPluginCopy(registeredPluginsMap[config.ID])
		} else {
			// 检查是否有二进制文件或源码
			binaryPath := pluginUtils.GenerateBinaryPath(pluginPath, config.ID, config.Binary)
			var status string
			if pluginUtils.CheckBinaryExists(binaryPath) {
				status = consts.PluginStatusAvailable
			} else {
				status = consts.PluginStatusSourceOnly
			}
			// 使用配置文件信息
			finalInfo = &config
			finalInfo.Status = status
		}

		discoveryInfo := &PluginDiscoveryInfo{
			PluginInfo:   finalInfo,
			Path:         pluginPath,
			IsRegistered: isRegistered,
		}

		discoveredPlugins = append(discoveredPlugins, discoveryInfo)
	}

	return discoveredPlugins, nil
}

// Shutdown 关闭所有插件
func (m *PluginManagerImpl) Shutdown() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for id, client := range m.plugins {
		client.Kill()
		if info, ok := m.infos[id]; ok {
			info.Status = consts.PluginStatusStopped
			info.IsExited = true
			global.SysLog.Infof("Plugin shutdown: %s (%s v%s) from %s, Author: %s, Binary: %s, Type: %s, Status: %s",
				info.ID, info.Name, info.Version, info.Repository, info.Author, info.Binary, info.Type, info.Status)
		}
	}

	m.plugins = make(map[string]*plugin.Client)
	m.infos = make(map[string]*PluginInfo)
}

// StartAutoPlugins 扫描并启动配置为自动启动的插件
func (m *PluginManagerImpl) StartAutoPlugins() error {
	cfgs, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	entries, err := os.ReadDir(cfgs.PluginConfig.PluginDir)
	if err != nil {
		if os.IsNotExist(err) {
			global.SysLog.Info("Plugin directory not found")
			return nil
		}
		return fmt.Errorf("failed to read plugin directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		// 构建插件路径
		pluginPath := filepath.Join(cfgs.PluginConfig.PluginDir, entry.Name())
		configFile := filepath.Join(pluginPath, cfgs.PluginConfig.PluginConfigFile)

		// 读取配置文件
		configData, err := os.ReadFile(configFile)
		if err != nil {
			global.SysLog.Warnf("Skipping %s: no %s found", entry.Name(), cfgs.PluginConfig.PluginConfigFile)
			continue
		}

		var config PluginInfo
		if err := json.Unmarshal(configData, &config); err != nil {
			global.SysLog.Warnf("Skipping %s: invalid %s", entry.Name(), cfgs.PluginConfig.PluginConfigFile)
			continue
		}

		// 只处理标记为自动启动的插件
		if !config.AutoStart {
			continue
		}

		// 生成二进制路径
		binaryPath := pluginUtils.GenerateBinaryPath(pluginPath, config.ID, config.Binary)

		// 检查二进制文件是否存在，不存在则尝试编译
		if !pluginUtils.CheckBinaryExists(binaryPath) {
			if err := pluginUtils.ExecuteBuildScript(pluginPath); err != nil {
				global.SysLog.Warnf("Plugin %s: build failed: %v", config.ID, err)
				continue
			}

			global.SysLog.Infof("Plugin built successfully: %s", config.ID)
		}

		// 更新配置中的二进制路径
		config.Binary = binaryPath

		// 注册并启动插件
		if err := m.RegisterPlugin(config.ID); err != nil {
			global.SysLog.Errorf("Failed to auto-start plugin %s: %v", config.ID, err)
		}
	}

	return nil
}

// refreshPluginInfo 刷新插件运行时信息
func (m *PluginManagerImpl) refreshPluginInfo(info *PluginInfo, client *plugin.Client) {
	info.ProcessID = client.ID()
	info.Protocol = string(client.Protocol())
	info.IsExited = client.Exited()
	info.NegotiatedVersion = client.NegotiatedVersion()

	if config := client.ReattachConfig(); config != nil {
		info.ProcessPID = config.Pid
		info.ProtocolVersion = config.ProtocolVersion
		if config.Addr != nil {
			info.NetworkAddr = config.Addr.String()
		}
	}
}

// createPluginCopy 创建插件信息的深拷贝
func (m *PluginManagerImpl) createPluginCopy(info *PluginInfo) *PluginInfo {
	return &PluginInfo{
		ID: info.ID, Name: info.Name, Version: info.Version, Author: info.Author,
		Description: info.Description, Repository: info.Repository, Binary: info.Binary,
		Type:      info.Type,
		AutoStart: info.AutoStart, StartTimeout: info.StartTimeout,
		MinPort: info.MinPort, MaxPort: info.MaxPort,
		AutoMTLS: info.AutoMTLS, Managed: info.Managed,
		Status: info.Status, StartedAt: info.StartedAt,
		ProcessID: info.ProcessID, Protocol: info.Protocol,
		IsExited: info.IsExited, NegotiatedVersion: info.NegotiatedVersion,
		ProcessPID: info.ProcessPID, ProtocolVersion: info.ProtocolVersion,
		NetworkAddr: info.NetworkAddr,
	}
}
