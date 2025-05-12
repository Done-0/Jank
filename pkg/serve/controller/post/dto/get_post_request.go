// Package dto 提供文章相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// GetOnePostRequest        获取文章的请求结构体
// @Param	id		path	string	true	"文章 ID"
type GetOnePostRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"omitempty" default:"0"`
}
