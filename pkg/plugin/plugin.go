// Package plugin 提供插件系统核心接口定义
// 创建者：Done-0
// 创建时间：2025-08-05
package plugin

import (
	"context"

	"github.com/hashicorp/go-plugin"

	"github.com/Done-0/jank/pkg/plugin/consts"
)

// Plugin 插件接口定义
type Plugin interface {
	Execute(ctx context.Context, method string, args map[string]any) (map[string]any, error)
	HealthCheck(ctx context.Context) error
}

// HandshakeConfig 插件握手配置
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "JANK_PLUGIN",
	MagicCookieValue: "jank-plugin",
}

// PluginMap 定义支持的插件接口类型映射（按功能分类）
var PluginMap = map[string]plugin.Plugin{
	consts.PluginTypeProvider: &GRPCPlugin{}, // 数据提供者插件
	consts.PluginTypeFilter:   &GRPCPlugin{}, // 数据过滤插件
	consts.PluginTypeHandler:  &GRPCPlugin{}, // 业务处理插件
	consts.PluginTypeNotifier: &GRPCPlugin{}, // 通知插件
}
