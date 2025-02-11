package dto

// DeleteCommentRequest 删除评论请求
// @Param id path int64 true "评论ID"
type DeleteCommentRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
}
