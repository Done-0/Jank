package dto

// GetCommentGraphRequest 获取评论请求
// @Param post_id path int true "帖子ID"
type GetCommentGraphRequest struct {
	PostID int64 `json:"post_id" xml:"post_id" form:"post_id" query:"post_id" validate:"required,gt=0"`
}
