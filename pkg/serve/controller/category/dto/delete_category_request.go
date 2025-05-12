// Package dto 提供类目相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// DeleteOneCategoryRequest  删除类目请求
// @Param id path int64 true "类目ID"
type DeleteOneCategoryRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
}
