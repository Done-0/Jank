// Package errno RBAC权限模块错误码定义
// 创建者：Done-0
// 创建时间：2025-08-05
package errno

import (
	"github.com/Done-0/jank/internal/utils/errorx/code"
)

// RBAC模块错误码: 50000 ~ 59999
const (
	ErrUserRoleNotFound = 50001 // 用户角色不存在
)

func init() {
	code.Register(ErrUserRoleNotFound, "user role not found: user {user} role {role}")
}
