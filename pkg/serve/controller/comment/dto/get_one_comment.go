package dto

// GetOneCommentRequest 获取评论请求
// @Param comment_id path int true "评论ID"
type GetOneCommentRequest struct {
	CommentID int64 `json:"comment_id" xml:"comment_id" form:"comment_id" query:"comment_id" validate:"required,gt=0"`
}
