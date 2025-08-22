// Package errno 文章模块错误码定义
// 创建者：Done-0
// 创建时间：2025-08-13
package errno

import (
	"github.com/Done-0/jank/internal/utils/errorx/code"
)

// 文章模块错误码: 40000 ~ 49999
const (
	ErrPostCreateFailed = 40001 // 创建文章失败
	ErrPostGetFailed    = 40002 // 获取文章失败
	ErrPostUpdateFailed = 40003 // 更新文章失败
	ErrPostDeleteFailed = 40004 // 删除文章失败
	ErrPostListFailed   = 40005 // 获取文章列表失败
)

func init() {
	code.Register(ErrPostCreateFailed, "create post failed: {title}")
	code.Register(ErrPostGetFailed, "get post failed: {id}")
	code.Register(ErrPostUpdateFailed, "update post failed: {id}")
	code.Register(ErrPostDeleteFailed, "delete post failed: {id}")
	code.Register(ErrPostListFailed, "list posts failed: {msg}")
}
