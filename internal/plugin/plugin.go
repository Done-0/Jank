package plugin

import (
	"context"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/plugin/impl"
)

// PluginManager 插件管理器接口
type PluginManager interface {
	// RegisterPlugin 注册插件
	RegisterPlugin(id string) error
	// UnregisterPlugin 注销插件
	UnregisterPlugin(id string) error
	// ExecutePlugin 执行插件
	ExecutePlugin(ctx context.Context, id, method string, args map[string]any) (map[string]any, error)
	// GetPlugin 获取插件信息
	GetPlugin(id string) (*impl.PluginInfo, error)
	// ListPlugins 列举所有插件（包括未注册的）
	ListPlugins() ([]*impl.PluginDiscoveryInfo, error)
	// StartAutoPlugins 启动自动启动的插件
	StartAutoPlugins() error
	// Shutdown 关闭插件系统
	Shutdown()
}

// 全局插件管理器实例
var GlobalPluginManager PluginManager

// New 初始化插件管理器
func New(config *configs.Config) {
	GlobalPluginManager = impl.NewPluginManager()
	global.SysLog.Info("Plugin system initialized")

	// 启动自动启动插件
	if err := GlobalPluginManager.StartAutoPlugins(); err != nil {
		global.SysLog.Errorf("Failed to start auto plugins: %v", err)
	}
}
