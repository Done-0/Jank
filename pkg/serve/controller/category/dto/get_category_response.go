package dto

// GetOneCategoryResponse 获取类目响应
type GetOneCategoryResponse struct {
	ID          int64                     `json:"id" xml:"id" form:"id" query:"id"`
	Name        string                    `json:"name" xml:"name" form:"name" query:"name"`
	Description string                    `json:"description" xml:"description" form:"description" query:"description"`
	ParentID    int64                     `json:"parent_id" xml:"parent_id" form:"parent_id" query:"parent_id"`
	IsActive    bool                      `json:"is_active" xml:"is_active" form:"is_active" query:"is_active"`
	Path        string                    `json:"path" xml:"path" form:"path" query:"path"`
	Children    []*GetOneCategoryResponse `json:"children" xml:"children" form:"children" query:"children"`
}
