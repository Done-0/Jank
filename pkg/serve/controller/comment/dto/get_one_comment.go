// Package dto 提供评论相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// GetOneCommentRequest 获取评论请求
// @Param comment_id path int true "评论ID"
type GetOneCommentRequest struct {
	CommentID int64 `json:"comment_id" xml:"comment_id" form:"comment_id" query:"comment_id" validate:"required"`
}
