package dto

// GetAllCategoriesResponse 获取所有类目响应
type GetAllCategoriesResponse struct {
	Categories []*GetOneCategoryResponse `json:"categories" xml:"categories" form:"categories" query:"categories"`
}
