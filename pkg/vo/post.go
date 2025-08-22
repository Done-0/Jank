// Package vo 文章相关值对象
// 创建者：Done-0
// 创建时间：2025-08-13
package vo

// CreatePostResponse 创建文章响应
type CreatePostResponse struct {
	ID           string `json:"id"`            // 文章 ID
	Title        string `json:"title"`         // 文章标题
	Description  string `json:"description"`   // 文章描述/摘要
	Image        string `json:"image"`         // 文章封面图片
	Status       string `json:"status"`        // 文章状态
	CategoryID   string `json:"category_id"`   // 分类 ID
	CategoryName string `json:"category_name"` // 分类名称
	Markdown     string `json:"markdown"`      // Markdown内容
	Message      string `json:"message"`       // 创建结果消息
}

// GetPostResponse 获取文章响应
type GetPostResponse struct {
	ID           string `json:"id"`            // 文章 ID
	Title        string `json:"title"`         // 文章标题
	Description  string `json:"description"`   // 文章描述/摘要
	Image        string `json:"image"`         // 文章封面图片
	Status       string `json:"status"`        // 文章状态
	CategoryID   string `json:"category_id"`   // 分类 ID
	CategoryName string `json:"category_name"` // 分类名称
	Markdown     string `json:"markdown"`      // Markdown 内容
	HTML         string `json:"html"`          // 渲染后的 HTML
	CreatedAt    string `json:"created_at"`    // 创建时间
	UpdatedAt    string `json:"updated_at"`    // 更新时间
}

// UpdatePostResponse 更新文章响应
type UpdatePostResponse struct {
	ID           string `json:"id"`            // 文章 ID
	Title        string `json:"title"`         // 文章标题
	Description  string `json:"description"`   // 文章描述/摘要
	Image        string `json:"image"`         // 文章封面图片
	Status       string `json:"status"`        // 文章状态
	CategoryID   string `json:"category_id"`   // 分类 ID
	CategoryName string `json:"category_name"` // 分类名称
	Markdown     string `json:"markdown"`      // Markdown内容
	Message      string `json:"message"`       // 更新结果消息
}

// DeletePostResponse 删除文章响应
type DeletePostResponse struct {
	Message string `json:"message"` // 删除结果消息
}

// PostItem 文章列表项
type PostItem struct {
	ID           string `json:"id"`            // 文章 ID
	Title        string `json:"title"`         // 文章标题
	Description  string `json:"description"`   // 文章描述/摘要
	Image        string `json:"image"`         // 文章封面图片
	Status       string `json:"status"`        // 文章状态
	CategoryID   string `json:"category_id"`   // 分类 ID
	CategoryName string `json:"category_name"` // 分类名称
	CreatedAt    string `json:"created_at"`    // 创建时间
	UpdatedAt    string `json:"updated_at"`    // 更新时间
}

// ListPostsResponse 文章列表响应
type ListPostsResponse struct {
	Total    int64       `json:"total"`     // 总数量
	PageNo   int64       `json:"page_no"`   // 当前页码
	PageSize int64       `json:"page_size"` // 每页数量
	List     []*PostItem `json:"list"`      // 文章列表
}
