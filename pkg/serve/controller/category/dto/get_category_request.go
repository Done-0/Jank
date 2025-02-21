package dto

// GetOneCategoryRequest 更新类目请求
// @Param id path int true "类目ID"
type GetOneCategoryRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required,gte=0"`
}
