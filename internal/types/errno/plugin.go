// Package errno 插件模块错误码定义
// 创建者：Done-0
// 创建时间：2025-08-06
package errno

import (
	"github.com/Done-0/jank/internal/utils/errorx/code"
)

// 插件模块错误码: 20000 ~ 29999
const (
	ErrPluginNotFound         = 20001 // 插件不存在
	ErrPluginRegisterFailed   = 20002 // 插件注册失败
	ErrPluginUnregisterFailed = 20003 // 插件注销失败
	ErrExecutePluginFailed    = 20004 // 执行插件失败
	ErrListPluginsFailed      = 20005 // 列举插件失败
)

func init() {
	code.Register(ErrPluginNotFound, "plugin not found: {plugin_id}")
	code.Register(ErrPluginRegisterFailed, "failed to register plugin: {plugin_id}")
	code.Register(ErrPluginUnregisterFailed, "failed to unregister plugin: {plugin_id}")
	code.Register(ErrExecutePluginFailed, "failed to execute plugin: {msg}")
	code.Register(ErrListPluginsFailed, "failed to list plugins: {msg}")
}
