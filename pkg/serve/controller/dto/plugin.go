// Package dto 提供插件相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-08-05
package dto

// RegisterPluginRequest 注册插件请求
type RegisterPluginRequest struct {
	ID      string `json:"id" validate:"required,min=1,max=100"` // 插件 ID
	Rebuild bool   `json:"rebuild,omitempty"`                    // 强制重新编译
}

// UnregisterPluginRequest 注销插件请求
type UnregisterPluginRequest struct {
	ID string `json:"id" validate:"required"` // 插件 ID
}

// GetPluginRequest 获取插件信息请求
type GetPluginRequest struct {
	ID string `query:"id" validate:"required"` // 插件 ID
}

// ListPluginsRequest 列举插件请求
type ListPluginsRequest struct {
	Status   string `query:"status" validate:"omitempty"`                 // 插件状态筛选
	PageNo   int64  `query:"page_no" validate:"required,min=1"`           // 页码
	PageSize int64  `query:"page_size" validate:"required,min=1,max=100"` // 每页数量
}

// ExecutePluginRequest 执行插件方法请求
type ExecutePluginRequest struct {
	ID     string         `json:"id" validate:"required"`     // 插件 ID
	Method string         `json:"method" validate:"required"` // 方法名
	Args   map[string]any `json:"args" validate:"omitempty"`  // 方法参数
}
