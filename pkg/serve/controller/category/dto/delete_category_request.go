package dto

// DeleteOneCategoryRequest 删除类目请求
type DeleteOneCategoryRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
}
