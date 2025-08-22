// Package theme 提供主题系统核心接口定义
// 创建者：Done-0
// 创建时间：2025-08-09
package theme

import (
	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/theme/impl"
)

// ThemeManager 主题管理器接口
type ThemeManager interface {
	// SwitchThemeByType 按类型切换主题（frontend/console）
	SwitchThemeByType(id string, themeType string) error
	// GetActiveThemeByType 按类型获取当前激活的主题（frontend/console）
	GetActiveThemeByType(themeType string) (*impl.ThemeInfo, error)
	// ListThemes 列举所有主题
	ListThemes() ([]*impl.ThemeInfo, error)
	// InitializeTheme 初始化主题系统，加载上次激活的主题或默认主题
	InitializeTheme() error
	// Shutdown 关闭主题系统
	Shutdown()
}

// 全局主题管理器实例
var GlobalThemeManager ThemeManager

// New 初始化主题管理器
func New(config *configs.Config) {
	GlobalThemeManager = impl.NewThemeManager()
	global.SysLog.Info("Theme system initialized")

	// 初始化主题系统
	if err := GlobalThemeManager.InitializeTheme(); err != nil {
		global.SysLog.Errorf("Failed to initialize theme system: %v", err)
	}
}
