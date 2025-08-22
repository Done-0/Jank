// Package consts 提供应用程序常量定义
// 创建者：Done-0
// 创建时间：2025-08-05
package consts

const (
	// 插件状态
	PluginStatusRunning    = "running"     // 插件正在运行
	PluginStatusStopped    = "stopped"     // 插件已停止
	PluginStatusError      = "error"       // 插件错误
	PluginStatusAvailable  = "available"   // 未注册但可用（有二进制文件）
	PluginStatusSourceOnly = "source_only" // 未注册但有源码（无二进制文件）
)

const (
	// 插件类型标识符
	PluginTypeProvider = "provider" // 数据提供者插件
	PluginTypeFilter   = "filter"   // 数据过滤插件
	PluginTypeHandler  = "handler"  // 业务处理插件
	PluginTypeNotifier = "notifier" // 通知插件
)
