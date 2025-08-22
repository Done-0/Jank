// Package dto 提供分类相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-08-13
package dto

// CreateCategoryRequest 创建分类请求
type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=100"`   // 分类名称
	Description string `json:"description" validate:"omitempty,max=500"` // 分类描述
	ParentID    string `json:"parent_id" validate:"omitempty"`           // 父分类 ID，为空表示顶级分类
	Sort        int64  `json:"sort" validate:"omitempty,min=0"`          // 排序权重，数字越大越靠前
	IsActive    bool   `json:"is_active" validate:"omitempty"`           // 是否启用，默认为true
}

// DeleteCategoryRequest 删除分类请求
type DeleteCategoryRequest struct {
	ID string `json:"id" validate:"required"` // 分类 ID
}

// GetCategoryRequest 获取分类请求
type GetCategoryRequest struct {
	ID string `query:"id" validate:"required"` // 分类 ID
}

// UpdateCategoryRequest 更新分类请求
type UpdateCategoryRequest struct {
	ID          string `json:"id" validate:"required"`                   // 分类 ID
	Name        string `json:"name" validate:"omitempty,min=1,max=100"`  // 分类名称
	Description string `json:"description" validate:"omitempty,max=500"` // 分类描述
	ParentID    string `json:"parent_id" validate:"omitempty"`           // 父分类 ID，为空表示顶级分类
	Sort        int64  `json:"sort" validate:"omitempty,min=0"`          // 排序权重，数字越大越靠前
	IsActive    bool   `json:"is_active" validate:"omitempty"`           // 是否启用
}

// ListCategoriesRequest 获取分类列表请求
type ListCategoriesRequest struct {
	PageNo   int64  `query:"page_no" validate:"required,min=1"`           // 页码
	PageSize int64  `query:"page_size" validate:"required,min=1,max=100"` // 每页数量
	ParentID string `query:"parent_id" validate:"omitempty"`              // 父分类 ID，为空时获取顶级分类
	IsActive *bool  `query:"is_active" validate:"omitempty"`              // 是否启用，为空时获取所有分类
}
