// Package vo 分类相关值对象
// 创建者：Done-0
// 创建时间：2025-08-13
package vo

// CreateCategoryResponse 创建分类响应
type CreateCategoryResponse struct {
	ID          string `json:"id"`          // 分类 ID
	Name        string `json:"name"`        // 分类名称
	Description string `json:"description"` // 分类描述
	ParentID    string `json:"parent_id"`   // 父分类 ID
	Sort        int64  `json:"sort"`        // 排序权重
	IsActive    bool   `json:"is_active"`   // 是否启用
	Message     string `json:"message"`     // 创建结果消息
}

// GetCategoryResponse 获取分类响应
type GetCategoryResponse struct {
	ID          string `json:"id"`          // 分类 ID
	Name        string `json:"name"`        // 分类名称
	Description string `json:"description"` // 分类描述
	ParentID    string `json:"parent_id"`   // 父分类 ID
	Sort        int64  `json:"sort"`        // 排序权重
	IsActive    bool   `json:"is_active"`   // 是否启用
	CreatedAt   string `json:"created_at"`  // 创建时间
	UpdatedAt   string `json:"updated_at"`  // 更新时间
}

// UpdateCategoryResponse 更新分类响应
type UpdateCategoryResponse struct {
	ID          string `json:"id"`          // 分类 ID
	Name        string `json:"name"`        // 分类名称
	Description string `json:"description"` // 分类描述
	ParentID    string `json:"parent_id"`   // 父分类 ID
	Sort        int64  `json:"sort"`        // 排序权重
	IsActive    bool   `json:"is_active"`   // 是否启用
	Message     string `json:"message"`     // 更新结果消息
}

// DeleteCategoryResponse 删除分类响应
type DeleteCategoryResponse struct {
	Message string `json:"message"` // 删除结果消息
}

// CategoryItem 分类列表项
type CategoryItem struct {
	ID          string `json:"id"`          // 分类 ID
	Name        string `json:"name"`        // 分类名称
	Description string `json:"description"` // 分类描述
	ParentID    string `json:"parent_id"`   // 父分类 ID
	Sort        int64  `json:"sort"`        // 排序权重
	IsActive    bool   `json:"is_active"`   // 是否启用
	CreatedAt   string `json:"created_at"`  // 创建时间
	UpdatedAt   string `json:"updated_at"`  // 更新时间
}

// ListCategoriesResponse 分类列表响应
type ListCategoriesResponse struct {
	Total    int64           `json:"total"`     // 总数量
	PageNo   int64           `json:"page_no"`   // 当前页码
	PageSize int64           `json:"page_size"` // 每页数量
	List     []*CategoryItem `json:"list"`      // 分类列表
}
