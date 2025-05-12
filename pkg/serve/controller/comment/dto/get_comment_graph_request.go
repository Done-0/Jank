// Package dto 提供评论相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// GetCommentGraphRequest 获取评论请求
// @Param post_id path int true "帖子ID"
type GetCommentGraphRequest struct {
	PostID int64 `json:"post_id" xml:"post_id" form:"post_id" query:"post_id" validate:"required"`
}
