// Package dto 提供评论相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// DeleteCommentRequest 删除评论请求
// @Param id path int64 true "评论ID"
type DeleteCommentRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
}
