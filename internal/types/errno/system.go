// Package errno 系统级错误码定义
// 创建者：Done-0
// 创建时间：2025-08-05
package errno

import (
	"github.com/Done-0/jank/internal/utils/errorx/code"
)

// 系统级错误码: 10000 ~ 19999
const (
	ErrInternalServer     = 10001 // 内部服务器错误
	ErrInvalidParams      = 10002 // 参数验证失败
	ErrUnauthorized       = 10003 // 身份认证失败
	ErrForbidden          = 10004 // 权限不足
	ErrResourceNotFound   = 10005 // 资源不存在
	ErrResourceConflict   = 10006 // 资源冲突
	ErrTooManyRequests    = 10007 // 请求频率超限
	ErrServiceUnavailable = 10008 // 服务暂不可用
)

func init() {
	code.Register(ErrInternalServer, "internal server error: {msg}")
	code.Register(ErrInvalidParams, "invalid parameter: {msg}")
	code.Register(ErrUnauthorized, "unauthorized access: {msg}")
	code.Register(ErrForbidden, "permission denied: {resource}")
	code.Register(ErrResourceNotFound, "{resource} not found: {id}")
	code.Register(ErrResourceConflict, "{resource} already exists: {id}")
	code.Register(ErrTooManyRequests, "too many requests: {limit} per {period}")
	code.Register(ErrServiceUnavailable, "service unavailable: {service}")
}
