package dto

// DeleteOnePostRequest    文章删除请求
// @Param id path int true "文章 ID"
type DeleteOnePostRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
}
