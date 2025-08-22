package impl

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/types/consts"
	"github.com/Done-0/jank/internal/utils/theme"
)

// ThemeManagerImpl 主题管理器实现
type ThemeManagerImpl struct {
	themes              map[string]*ThemeInfo // 主题信息映射
	frontendActiveTheme string                // 当前激活的 Frontend 主题 ID
	consoleActiveTheme  string                // 当前激活的 Console 主题 ID
	mu                  sync.RWMutex          // 并发安全锁
}

// NewThemeManager 创建主题管理器实例
func NewThemeManager() *ThemeManagerImpl {
	return &ThemeManagerImpl{
		themes: make(map[string]*ThemeInfo),
	}
}

// SwitchThemeByType 按类型切换主题
func (m *ThemeManagerImpl) SwitchThemeByType(themeID string, themeType string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 1. 检查主题是否已加载，如果未加载则动态加载
	targetTheme, isThemeLoaded := m.themes[themeID]
	if !isThemeLoaded {
		// 1.1 获取应用配置
		cfgs, err := configs.GetConfig()
		if err != nil {
			return fmt.Errorf("failed to load application config: %w", err)
		}

		// 1.2 扫描主题目录
		themeDirectories, err := os.ReadDir(cfgs.ThemeConfig.ThemeDir)
		if err != nil {
			return fmt.Errorf("failed to scan theme directory: %w", err)
		}

		// 1.3 查找目标主题的配置文件
		var foundThemePath string
		var foundThemeConfigData []byte

		for _, directoryEntry := range themeDirectories {
			if !directoryEntry.IsDir() {
				continue
			}

			currentThemeDir := filepath.Join(cfgs.ThemeConfig.ThemeDir, directoryEntry.Name())
			themeConfigFilePath := filepath.Join(currentThemeDir, cfgs.ThemeConfig.ThemeConfigFile)

			configFileData, readErr := os.ReadFile(themeConfigFilePath)
			if readErr != nil {
				continue // 跳过无法读取配置文件的目录
			}

			var candidateThemeInfo ThemeInfo
			if unmarshalErr := json.Unmarshal(configFileData, &candidateThemeInfo); unmarshalErr == nil && candidateThemeInfo.ID == themeID {
				foundThemePath = currentThemeDir
				foundThemeConfigData = configFileData
				break
			}
		}

		// 1.4 验证主题是否找到
		if foundThemePath == "" {
			return fmt.Errorf("theme with ID '%s' not found in theme directory", themeID)
		}

		// 1.5 解析主题配置
		var loadedThemeConfig ThemeInfo
		if err := json.Unmarshal(foundThemeConfigData, &loadedThemeConfig); err != nil {
			return fmt.Errorf("invalid theme configuration for '%s': %w", themeID, err)
		}

		// 1.6 初始化主题信息
		loadedThemeConfig.Path = foundThemePath
		loadedThemeConfig.Status = ""
		loadedThemeConfig.IsActive = false

		// 1.7 执行主题构建脚本
		if err := theme.ExecuteBuildScript(foundThemePath); err != nil {
			return fmt.Errorf("failed to build theme '%s': %w", themeID, err)
		}

		// 1.8 缓存主题信息
		m.themes[loadedThemeConfig.ID] = &loadedThemeConfig
		targetTheme = &loadedThemeConfig
	}

	// 2. 停用当前激活的主题并设置新的激活主题
	switch themeType {
	case consts.ThemeTypeFrontend:
		// 停用当前前端主题
		if m.frontendActiveTheme != "" {
			if currentFrontendTheme, exists := m.themes[m.frontendActiveTheme]; exists {
				currentFrontendTheme.IsActive = false
				currentFrontendTheme.Status = ""
			}
		}
		m.frontendActiveTheme = themeID

	case consts.ThemeTypeConsole:
		// 停用当前控制台主题
		if m.consoleActiveTheme != "" {
			if currentConsoleTheme, exists := m.themes[m.consoleActiveTheme]; exists {
				currentConsoleTheme.IsActive = false
				currentConsoleTheme.Status = ""
			}
		}
		m.consoleActiveTheme = themeID

	default:
		return fmt.Errorf("unsupported theme type: %s", themeType)
	}

	// 3. 激活目标主题
	targetTheme.IsActive = true
	targetTheme.Status = consts.ThemeStatusActive
	targetTheme.LoadedAt = time.Now().Unix()

	// 4. 持久化主题配置到配置文件
	if err := configs.UpdateField(func(config *configs.Config) {
		switch themeType {
		case consts.ThemeTypeFrontend:
			config.ThemeConfig.FrontendLastActiveTheme = themeID
		case consts.ThemeTypeConsole:
			config.ThemeConfig.ConsoleLastActiveTheme = themeID
		}
	}); err != nil {
		global.SysLog.Warnf("Failed to persist %s active theme configuration: %v", themeType, err)
	}

	global.SysLog.Infof("Theme switched successfully: %s (%s v%s) from %s, Author: %s, Path: %s, Status: %s",
		targetTheme.ID, targetTheme.Name, targetTheme.Version, targetTheme.Repository,
		targetTheme.Author, targetTheme.Path, targetTheme.Status)
	return nil
}

// GetActiveThemeByType 按类型获取当前激活的主题
func (m *ThemeManagerImpl) GetActiveThemeByType(themeType string) (*ThemeInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var activeThemeID string
	switch themeType {
	case consts.ThemeTypeFrontend:
		activeThemeID = m.frontendActiveTheme
	case consts.ThemeTypeConsole:
		activeThemeID = m.consoleActiveTheme
	default:
		return nil, fmt.Errorf("unsupported theme type: %s", themeType)
	}

	if activeThemeID == "" {
		return nil, fmt.Errorf("no active %s theme found", themeType)
	}

	theme, exists := m.themes[activeThemeID]
	if !exists {
		return nil, fmt.Errorf("active %s theme %s not found", themeType, activeThemeID)
	}

	themeCopy := *theme
	return &themeCopy, nil
}

// ListThemes 列出所有主题
func (m *ThemeManagerImpl) ListThemes() ([]*ThemeInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	cfgs, err := configs.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	entries, err := os.ReadDir(cfgs.ThemeConfig.ThemeDir)
	if err != nil {
		return nil, fmt.Errorf("failed to scan theme directory: %w", err)
	}

	var themes []*ThemeInfo
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		themePath := filepath.Join(cfgs.ThemeConfig.ThemeDir, entry.Name())
		configPath := filepath.Join(themePath, cfgs.ThemeConfig.ThemeConfigFile)
		if configData, err := os.ReadFile(configPath); err == nil {
			var themeInfo ThemeInfo
			if err := json.Unmarshal(configData, &themeInfo); err == nil {
				if loadedTheme, exists := m.themes[themeInfo.ID]; exists {
					themes = append(themes, loadedTheme)
				} else {
					themeInfo.Path = themePath
					themeInfo.Status = ""
					themeInfo.IsActive = (themeInfo.ID == m.frontendActiveTheme)
					themes = append(themes, &themeInfo)
				}
			}
		}
	}
	return themes, nil
}

// InitializeTheme 初始化主题系统，加载上次激活的主题或默认主题
func (m *ThemeManagerImpl) InitializeTheme() error {
	cfgs, err := configs.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load application config: %w", err)
	}

	// 扫描主题目录
	themeDirectories, err := os.ReadDir(cfgs.ThemeConfig.ThemeDir)
	if err != nil {
		return fmt.Errorf("failed to scan theme directory: %w", err)
	}

	// 初始化两种类型的主题
	themeTypes := []struct {
		themeType    string
		lastActiveID string
		defaultID    string
		isActive     bool
	}{
		{
			themeType:    consts.ThemeTypeFrontend,
			lastActiveID: cfgs.ThemeConfig.FrontendLastActiveTheme,
			defaultID:    cfgs.ThemeConfig.FrontendDefaultTheme,
			isActive:     true,
		},
		{
			themeType:    consts.ThemeTypeConsole,
			lastActiveID: cfgs.ThemeConfig.ConsoleLastActiveTheme,
			defaultID:    cfgs.ThemeConfig.ConsoleDefaultTheme,
			isActive:     false,
		},
	}

	for _, themeTypeConfig := range themeTypes {
		// 1. 确定要加载的主题 ID
		targetThemeID := themeTypeConfig.lastActiveID
		switch {
		case targetThemeID == "":
			targetThemeID = themeTypeConfig.defaultID
		}

		// 2. 尝试加载目标主题
		var loadedTheme *ThemeInfo
		for _, directoryEntry := range themeDirectories {
			if !directoryEntry.IsDir() {
				continue
			}

			currentThemeDir := filepath.Join(cfgs.ThemeConfig.ThemeDir, directoryEntry.Name())
			themeConfigFilePath := filepath.Join(currentThemeDir, cfgs.ThemeConfig.ThemeConfigFile)

			configFileData, readErr := os.ReadFile(themeConfigFilePath)
			if readErr != nil {
				continue
			}

			var candidateThemeInfo ThemeInfo
			if unmarshalErr := json.Unmarshal(configFileData, &candidateThemeInfo); unmarshalErr == nil && candidateThemeInfo.ID == targetThemeID {
				candidateThemeInfo.Path = currentThemeDir

				// 尝试构建主题
				switch buildErr := theme.ExecuteBuildScript(currentThemeDir); buildErr {
				case nil:
					loadedTheme = &candidateThemeInfo
				default:
					// 对于前端主题，构建失败是致命错误
					switch themeTypeConfig.themeType {
					case consts.ThemeTypeFrontend:
						return fmt.Errorf("failed to build frontend theme %s: %w", targetThemeID, buildErr)
					case consts.ThemeTypeConsole:
						global.SysLog.Warnf("Failed to build console theme %s: %v", targetThemeID, buildErr)
						continue
					}
				}
				break
			}
		}

		// 3. 如果目标主题加载失败，尝试回退到默认主题
		switch {
		case loadedTheme == nil && targetThemeID != themeTypeConfig.defaultID:
			global.SysLog.Warnf("Failed to load %s theme %s, falling back to default theme %s",
				themeTypeConfig.themeType, targetThemeID, themeTypeConfig.defaultID)

			for _, directoryEntry := range themeDirectories {
				if !directoryEntry.IsDir() {
					continue
				}

				currentThemeDir := filepath.Join(cfgs.ThemeConfig.ThemeDir, directoryEntry.Name())
				themeConfigFilePath := filepath.Join(currentThemeDir, cfgs.ThemeConfig.ThemeConfigFile)

				configFileData, readErr := os.ReadFile(themeConfigFilePath)
				if readErr != nil {
					continue
				}

				var candidateThemeInfo ThemeInfo
				if unmarshalErr := json.Unmarshal(configFileData, &candidateThemeInfo); unmarshalErr == nil && candidateThemeInfo.ID == themeTypeConfig.defaultID {
					candidateThemeInfo.Path = currentThemeDir

					switch buildErr := theme.ExecuteBuildScript(currentThemeDir); buildErr {
					case nil:
						loadedTheme = &candidateThemeInfo
					default:
						switch themeTypeConfig.themeType {
						case consts.ThemeTypeFrontend:
							return fmt.Errorf("failed to build default frontend theme: %w", buildErr)
						case consts.ThemeTypeConsole:
							global.SysLog.Warnf("Failed to build default console theme: %v", buildErr)
							continue
						}
					}
					break
				}
			}
		}

		// 4. 处理主题加载结果
		switch themeTypeConfig.themeType {
		case consts.ThemeTypeFrontend:
			switch loadedTheme {
			case nil:
				return fmt.Errorf("frontend theme %s not found", targetThemeID)
			default:
				// 激活前端主题
				loadedTheme.Status = consts.ThemeStatusActive
				loadedTheme.LoadedAt = time.Now().Unix()
				loadedTheme.IsActive = themeTypeConfig.isActive
				m.themes[loadedTheme.ID] = loadedTheme
				m.frontendActiveTheme = loadedTheme.ID

				// 持久化配置
				if err := configs.UpdateField(func(config *configs.Config) {
					config.ThemeConfig.FrontendLastActiveTheme = loadedTheme.ID
				}); err != nil {
					global.SysLog.Warnf("Failed to persist frontend active theme configuration: %v", err)
				}

				global.SysLog.Infof("Frontend theme system initialized: %s (%s v%s) from %s, Author: %s, Path: %s, Status: %s",
					loadedTheme.ID, loadedTheme.Name, loadedTheme.Version, loadedTheme.Repository,
					loadedTheme.Author, loadedTheme.Path, loadedTheme.Status)
			}

		case consts.ThemeTypeConsole:
			switch {
			case loadedTheme != nil:
				// 激活控制台主题
				loadedTheme.Status = consts.ThemeStatusActive
				loadedTheme.LoadedAt = time.Now().Unix()
				loadedTheme.IsActive = themeTypeConfig.isActive
				m.themes[loadedTheme.ID] = loadedTheme
				m.consoleActiveTheme = loadedTheme.ID

				// 持久化配置
				if err := configs.UpdateField(func(config *configs.Config) {
					config.ThemeConfig.ConsoleLastActiveTheme = loadedTheme.ID
				}); err != nil {
					global.SysLog.Warnf("Failed to persist console active theme configuration: %v", err)
				}

				global.SysLog.Infof("Console theme system initialized: %s (%s v%s) from %s, Author: %s, Path: %s, Status: %s",
					loadedTheme.ID, loadedTheme.Name, loadedTheme.Version, loadedTheme.Repository,
					loadedTheme.Author, loadedTheme.Path, loadedTheme.Status)
			default:
				global.SysLog.Warnf("Console theme %s not found, console theme system not initialized", targetThemeID)
			}
		}
	}

	return nil
}

// Shutdown 关闭主题系统
func (m *ThemeManagerImpl) Shutdown() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, theme := range m.themes {
		global.SysLog.Infof("Theme shutdown: %s (%s v%s) from %s, Author: %s, Path: %s, Status: %s",
			theme.ID, theme.Name, theme.Version, theme.Repository, theme.Author, theme.Path, theme.Status)
	}

	m.themes = make(map[string]*ThemeInfo)
	m.frontendActiveTheme = ""
	m.consoleActiveTheme = ""
}
