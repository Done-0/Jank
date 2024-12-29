package dto

// CreateOneCategoryRequest 创建类目请求
type CreateOneCategoryRequest struct {
	Name        string `json:"name" xml:"name" form:"name" query:"name" validate:"required,min=1,max=255"`
	Description string `json:"description" xml:"description" form:"description" query:"description" validate:"min=0,max=225" default:""`
	ParentID    int64  `json:"parent_id" xml:"parent_id" form:"parent_id" query:"parent_id"`
}
