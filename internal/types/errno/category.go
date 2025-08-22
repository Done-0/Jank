// Package errno 分类模块错误码定义
// 创建者：Done-0
// 创建时间：2025-08-13
package errno

import (
	"github.com/Done-0/jank/internal/utils/errorx/code"
)

// 分类模块错误码: 70000 ~ 79999
const (
	ErrCategoryCreateFailed = 70001 // 创建分类失败
	ErrCategoryGetFailed    = 70002 // 获取分类失败
	ErrCategoryUpdateFailed = 70003 // 更新分类失败
	ErrCategoryDeleteFailed = 70004 // 删除分类失败
	ErrCategoryListFailed   = 70005 // 获取分类列表失败
)

func init() {
	code.Register(ErrCategoryCreateFailed, "create category failed: {name}")
	code.Register(ErrCategoryGetFailed, "get category failed: {id}")
	code.Register(ErrCategoryUpdateFailed, "update category failed: {id}")
	code.Register(ErrCategoryDeleteFailed, "delete category failed: {id}")
	code.Register(ErrCategoryListFailed, "list categories failed: {msg}")
}
