// Package dto 提供类目相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// GetOneCategoryRequest 更新类目请求
// @Param id path int true "类目ID"
type GetOneCategoryRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
}
