package dto

// DeleteOneCategoryRequest  删除类目请求
// @Param id path int64 true "类目ID"
type DeleteOneCategoryRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
}
