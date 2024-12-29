package dto

// UpdateOneCategoryRequest 更新类目请求
type UpdateOneCategoryRequest struct {
	ID          int64                     `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
	Name        string                    `json:"name" xml:"name" form:"name" query:"name" validate:"required,min=1,max=255"`
	Description string                    `json:"description" xml:"description" form:"description" query:"description" validate:"required,min=0,max=255" default:""`
	ParentID    int64                     `json:"parent_id" xml:"parent_id" form:"parent_id" query:"parent_id"`
	IsActive    bool                      `json:"is_active" xml:"is_active" form:"is_active" query:"is_active"`
	Path        string                    `json:"path" xml:"path" form:"path" query:"path"`
	Children    []*GetOneCategoryResponse `json:"children" xml:"children" form:"children" query:"children"`
}
