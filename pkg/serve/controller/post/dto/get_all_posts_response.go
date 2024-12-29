package dto

// GetAllPostsResponse 获取全部文章响应
// @Description 获取全部文章响应
// @Summary 获取全部文章响应
// @Produce json
// @Success 200 {object} GetAllPostsResponse
// @Router /api/v1/posts [get]
type GetAllPostsResponse struct {
	ID              int64   `json:"id" xml:"id" form:"id" query:"id"`
	Title           string  `json:"title" xml:"title" form:"title" query:"title"`
	Image           string  `json:"image" xml:"image" form:"image" query:"image"`
	Visibility      string  `json:"visibility" xml:"visibility" form:"visibility" query:"visibility"`
	ContentMarkdown string  `json:"contentMarkdown" xml:"contentMarkdown" form:"contentMarkdown" query:"contentMarkdown"`
	ContentHTML     string  `json:"contentHtml" xml:"contentHtml" form:"contentHtml" query:"contentHtml"`
	CategoryIDs     []int64 `json:"category_ids" xml:"category_ids" form:"category_ids" query:"category_ids"`
}
