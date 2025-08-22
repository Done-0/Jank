// Package dto 提供文章相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-08-13
package dto

// CreatePostRequest 创建文章请求
type CreatePostRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`                            // 文章标题
	Description string `json:"description" validate:"omitempty,max=500"`                           // 文章描述/摘要
	Image       string `json:"image" validate:"omitempty,url"`                                     // 文章封面图片
	Status      string `json:"status" validate:"omitempty,oneof=draft published private archived"` // 文章状态
	CategoryID  string `json:"category_id" validate:"omitempty"`                                   // 分类 ID
	Markdown    string `json:"markdown" validate:"omitempty,max=100000"`                           // Markdown 内容
}

// DeletePostRequest 删除文章请求
type DeletePostRequest struct {
	ID string `json:"id" validate:"required"` // 文章 ID
}

// GetPostRequest 获取文章请求
type GetPostRequest struct {
	ID string `query:"id" validate:"required"` // 文章 ID
}

// UpdatePostRequest 更新文章请求
type UpdatePostRequest struct {
	ID          string `json:"id" validate:"required"`                                             // 文章 ID
	Title       string `json:"title" validate:"omitempty,min=1,max=255"`                           // 文章标题
	Description string `json:"description" validate:"omitempty,max=500"`                           // 文章描述/摘要
	Image       string `json:"image" validate:"omitempty,url"`                                     // 文章封面图片
	Status      string `json:"status" validate:"omitempty,oneof=draft published private archived"` // 文章状态
	CategoryID  string `json:"category_id" validate:"omitempty"`                                   // 分类 ID
	Markdown    string `json:"markdown" validate:"omitempty,max=100000"`                           // Markdown内容
}

// ListPublishedPostsRequest 获取文章列表请求
type ListPublishedPostsRequest struct {
	PageNo     int64  `query:"page_no" validate:"required,min=1"`           // 页码
	PageSize   int64  `query:"page_size" validate:"required,min=1,max=100"` // 每页数量
	CategoryID *int64 `query:"category_id" validate:"omitempty"`            // 分类ID，为空时不按分类筛选
}

// ListPostsByStatusRequest 根据状态获取文章列表请求
type ListPostsByStatusRequest struct {
	PageNo     int64  `query:"page_no" validate:"required,min=1"`                                  // 页码
	PageSize   int64  `query:"page_size" validate:"required,min=1,max=100"`                        // 每页数量
	Status     string `query:"status" validate:"omitempty,oneof=draft published private archived"` // 文章状态，为空时获取所有文章
	CategoryID *int64 `query:"category_id" validate:"omitempty"`                                   // 分类ID，为空时不按分类筛选，有值时必须大于0
}
