// Package mapper 提供RBAC权限相关的数据访问接口
// 创建者：Done-0
// 创建时间：2025-08-11
package mapper

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/model/rbac"
)

// RBACMapper RBAC数据访问接口
type RBACMapper interface {
	// 权限管理
	CreatePermission(c *app.RequestContext, name, description, role, resource, action string) (*rbac.Policy, error) // 创建权限
	DeletePermission(c *app.RequestContext, role, resource, action string) (bool, error)                            // 删除权限
	ListPermissions(c *app.RequestContext) ([]*rbac.Policy, error)                                                  // 获取所有权限
	PermissionExists(c *app.RequestContext, resource, action string) (bool, error)                                  // 检查权限是否存在

	// 角色管理
	ListRoles(c *app.RequestContext) ([]*rbac.Policy, error)                       // 获取所有角色
	RoleExists(c *app.RequestContext, role string) (bool, error)                   // 检查角色是否存在
	GetRolePermissions(c *app.RequestContext, role string) ([]*rbac.Policy, error) // 获取角色权限

	// 用户角色管理
	AssignRole(c *app.RequestContext, user, role string) (*rbac.Policy, error) // 分配角色
	RevokeRole(c *app.RequestContext, user, role string) (bool, error)         // 撤销角色
	GetUserRoles(c *app.RequestContext, user string) ([]*rbac.Policy, error)   // 获取用户角色
	UserHasRole(c *app.RequestContext, user, role string) (bool, error)        // 检查用户是否有指定角色
	ListUsers(c *app.RequestContext) ([]*rbac.Policy, error)                   // 获取所有用户

	// 权限检查
	CheckPermission(c *app.RequestContext, user, resource, action string) (bool, error) // 权限检查
}
