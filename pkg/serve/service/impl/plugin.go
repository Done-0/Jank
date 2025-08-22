package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/plugin"
	"github.com/Done-0/jank/internal/plugin/impl"
	"github.com/Done-0/jank/internal/utils/logger"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/service"
	"github.com/Done-0/jank/pkg/vo"

	pluginUtils "github.com/Done-0/jank/internal/utils/plugin"
)

// PluginServiceImpl 插件服务实现
type PluginServiceImpl struct{}

// NewPluginService 创建插件服务实例
func NewPluginService() service.PluginService {
	return &PluginServiceImpl{}
}

// RegisterPlugin 注册插件逻辑
func (s *PluginServiceImpl) RegisterPlugin(c *app.RequestContext, req *dto.RegisterPluginRequest) (*vo.RegisterPluginResponse, error) {
	cfgs, err := configs.GetConfig()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get config: %v", err)
		log.Fatalf("failed to get config: %v", err)
	}

	entries, err := os.ReadDir(cfgs.PluginConfig.PluginDir)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to read plugin directory: %v", err)
		return &vo.RegisterPluginResponse{Message: err.Error()}, fmt.Errorf("failed to read plugin directory: %w", err)
	}

	var pluginInfo *impl.PluginInfo
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		pluginPath := filepath.Join(cfgs.PluginConfig.PluginDir, entry.Name())
		configFile := filepath.Join(pluginPath, cfgs.PluginConfig.PluginConfigFile)

		configData, err := os.ReadFile(configFile)
		if err != nil {
			continue
		}

		var config impl.PluginInfo
		if err := json.Unmarshal(configData, &config); err != nil {
			continue
		}

		if config.ID == req.ID {
			binaryPath := pluginUtils.GenerateBinaryPath(pluginPath, config.ID, config.Binary)

			// 检查是否需要构建：强制重建 或 二进制文件不存在
			isRebuild := req.Rebuild || !pluginUtils.CheckBinaryExists(binaryPath)
			if isRebuild {
				logger.BizLogger(c).Infof("building plugin: %s", req.ID)
				if err := pluginUtils.ExecuteBuildScript(pluginPath); err != nil {
					logger.BizLogger(c).Errorf("failed to build plugin %s: %v", req.ID, err)
					return &vo.RegisterPluginResponse{Message: err.Error()},
						fmt.Errorf("build failed: %w", err)
				}
			}

			config.Binary = binaryPath
			pluginInfo = &config
			break
		}
	}

	if pluginInfo == nil {
		logger.BizLogger(c).Errorf("plugin not found: %s", req.ID)
		return &vo.RegisterPluginResponse{Message: "plugin not found"},
			fmt.Errorf("plugin with ID %s not found", req.ID)
	}

	if err := plugin.GlobalPluginManager.RegisterPlugin(req.ID); err != nil {
		logger.BizLogger(c).Errorf("failed to register plugin %s: %v", req.ID, err)
		return &vo.RegisterPluginResponse{Message: err.Error()}, fmt.Errorf("failed to register plugin %s: %v", req.ID, err)
	}

	return &vo.RegisterPluginResponse{Message: "Plugin registered successfully"}, nil
}

// UnregisterPlugin 注销插件逻辑
func (s *PluginServiceImpl) UnregisterPlugin(c *app.RequestContext, req *dto.UnregisterPluginRequest) (*vo.UnregisterPluginResponse, error) {
	if err := plugin.GlobalPluginManager.UnregisterPlugin(req.ID); err != nil {
		logger.BizLogger(c).Errorf("failed to unregister plugin %s: %v", req.ID, err)
		return &vo.UnregisterPluginResponse{Message: err.Error()}, err
	}
	return &vo.UnregisterPluginResponse{Message: "Plugin unregistered successfully"}, nil
}

// ExecutePlugin 执行插件方法逻辑
func (s *PluginServiceImpl) ExecutePlugin(c *app.RequestContext, req *dto.ExecutePluginRequest) (*vo.ExecutePluginResponse, error) {
	result, err := plugin.GlobalPluginManager.ExecutePlugin(context.Background(), req.ID, req.Method, req.Args)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to execute plugin %s method %s: %v", req.ID, req.Method, err)
		return &vo.ExecutePluginResponse{}, fmt.Errorf("failed to execute plugin %s method %s: %v", req.ID, req.Method, err)
	}

	resultMap := make(map[string]any)
	for k, v := range result {
		resultMap[k] = v
	}

	return &vo.ExecutePluginResponse{
		Method: req.Method,
		Data:   resultMap,
	}, nil
}

// GetPlugin 获取插件信息逻辑
func (s *PluginServiceImpl) GetPlugin(c *app.RequestContext, req *dto.GetPluginRequest) (*vo.GetPluginResponse, error) {
	info, err := plugin.GlobalPluginManager.GetPlugin(req.ID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get plugin %s: %v", req.ID, err)
		return &vo.GetPluginResponse{}, fmt.Errorf("failed to get plugin %s: %v", req.ID, err)
	}

	response := vo.GetPluginResponse{
		// 基本信息
		ID:          info.ID,
		Name:        info.Name,
		Version:     info.Version,
		Author:      info.Author,
		Description: info.Description,
		Repository:  info.Repository,
		Binary:      info.Binary,
		Type:        info.Type,

		// 配置信息
		AutoStart:    info.AutoStart,
		StartTimeout: info.StartTimeout,
		MinPort:      info.MinPort,
		MaxPort:      info.MaxPort,
		AutoMTLS:     info.AutoMTLS,
		Managed:      info.Managed,

		// 运行时信息
		Status:            info.Status,
		StartedAt:         info.StartedAt,
		ProcessID:         info.ProcessID,
		Protocol:          info.Protocol,
		IsExited:          info.IsExited,
		NegotiatedVersion: info.NegotiatedVersion,
		ProcessPID:        info.ProcessPID,
		ProtocolVersion:   info.ProtocolVersion,
		NetworkAddr:       info.NetworkAddr,
	}
	return &response, nil
}

// ListPlugins 列举所有插件逻辑（包括未注册的）
func (s *PluginServiceImpl) ListPlugins(c *app.RequestContext, req *dto.ListPluginsRequest) (*vo.ListPluginsResponse, error) {
	discoveredPlugins, err := plugin.GlobalPluginManager.ListPlugins()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to list plugins: %v", err)
		return &vo.ListPluginsResponse{}, fmt.Errorf("failed to list plugins: %w", err)
	}

	var filteredPlugins []vo.GetPluginResponse
	for _, discovered := range discoveredPlugins {
		if req.Status != "" && discovered.Status != req.Status {
			continue
		}

		pluginVO := vo.GetPluginResponse{
			// 基本信息
			ID:          discovered.ID,
			Name:        discovered.Name,
			Version:     discovered.Version,
			Author:      discovered.Author,
			Description: discovered.Description,
			Repository:  discovered.Repository,
			Binary:      discovered.Binary,
			Type:        discovered.Type,

			// 配置信息
			AutoStart:    discovered.AutoStart,
			StartTimeout: discovered.StartTimeout,
			MinPort:      discovered.MinPort,
			MaxPort:      discovered.MaxPort,
			AutoMTLS:     discovered.AutoMTLS,
			Managed:      discovered.Managed,

			// 运行时信息
			Status:            discovered.Status,
			StartedAt:         discovered.StartedAt,
			ProcessID:         discovered.ProcessID,
			Protocol:          discovered.Protocol,
			IsExited:          discovered.IsExited,
			NegotiatedVersion: discovered.NegotiatedVersion,
			ProcessPID:        discovered.ProcessPID,
			ProtocolVersion:   discovered.ProtocolVersion,
			NetworkAddr:       discovered.NetworkAddr,
		}
		filteredPlugins = append(filteredPlugins, pluginVO)
	}

	total := int64(len(filteredPlugins))

	pageNo := req.PageNo
	pageSize := req.PageSize
	start := (pageNo - 1) * pageSize
	end := start + pageSize
	if start >= total {
		filteredPlugins = []vo.GetPluginResponse{}
	} else {
		if end > total {
			end = total
		}
		filteredPlugins = filteredPlugins[start:end]
	}

	return &vo.ListPluginsResponse{
		Total:    total,
		PageNo:   pageNo,
		PageSize: pageSize,
		List:     filteredPlugins,
	}, nil
}
