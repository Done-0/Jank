// Package dto 提供主题相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-08-09
package dto

// SwitchThemeRequest 切换主题请求
type SwitchThemeRequest struct {
	ID        string `json:"id" form:"id" binding:"required"`                      // 主题 ID
	ThemeType string `json:"theme_type" binding:"required,oneof=frontend console"` // 主题类型：frontend/console
	Rebuild   bool   `json:"rebuild,omitempty"`                                    // 是否重新构建主题，可选参数
}

// ListThemesRequest 列举主题请求
type ListThemesRequest struct {
	Status   string `query:"status" validate:"omitempty"`        // 主题状态筛选
	PageNo   int64  `query:"page_no" validate:"min=1"`           // 页码
	PageSize int64  `query:"page_size" validate:"min=1,max=100"` // 每页数量
}
