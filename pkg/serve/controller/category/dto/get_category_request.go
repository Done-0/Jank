package dto

// GetOneCategoryRequest 更新类目请求
type GetOneCategoryRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
}
