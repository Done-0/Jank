// Package dto 提供评论相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// CreateCommentRequest 创建评论请求
// @Param content     body string  true  "评论内容"
// @Param user_id     body int64   true  "用户ID"
// @Param post_id     body int64   true  "文章ID"
// @Param reply_to_comment_id body int64 false "回复的评论ID"
type CreateCommentRequest struct {
	Content          string `json:"content" xml:"content" form:"content" query:"content" validate:"required,min=1,max=1024"`
	UserId           int64  `json:"user_id" xml:"user_id" form:"user_id" query:"user_id" validate:"required"`
	PostId           int64  `json:"post_id" xml:"post_id" form:"post_id" query:"post_id" validate:"required"`
	ReplyToCommentId int64  `json:"reply_to_comment_id" xml:"reply_to_comment_id" form:"reply_to_comment_id" query:"reply_to_comment_id" validate:"omitempty"`
}
