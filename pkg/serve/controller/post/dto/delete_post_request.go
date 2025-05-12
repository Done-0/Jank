// Package dto 提供文章相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// DeleteOnePostRequest    文章删除请求
// @Param id path int true "文章 ID"
type DeleteOnePostRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
}
