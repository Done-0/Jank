package service

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/vo"
)

// RBACService 权限管理服务接口
type RBACService interface {
	// 权限策略管理
	CreatePermission(c *app.RequestContext, req *dto.CreatePermissionRequest) (*vo.PolicyOpResponse, error) // 创建权限策略
	DeletePermission(c *app.RequestContext, req *dto.DeletePermissionRequest) (*vo.PolicyOpResponse, error) // 删除权限策略
	ListPermissions(c *app.RequestContext) (*vo.PermissionListResponse, error)                              // 获取所有权限
	AssignPermission(c *app.RequestContext, req *dto.AssignPermissionRequest) (*vo.PolicyOpResponse, error) // 为角色分配权限
	RevokePermission(c *app.RequestContext, req *dto.RevokePermissionRequest) (*vo.PolicyOpResponse, error) // 撤销角色权限

	// 角色管理
	CreateRole(c *app.RequestContext, req *dto.CreateRoleRequest) (*vo.CreateRoleResponse, error)                      // 创建角色
	DeleteRole(c *app.RequestContext, req *dto.DeleteRoleRequest) (*vo.DeleteRoleResponse, error)                      // 删除角色
	ListRoles(c *app.RequestContext) (*vo.RoleListResponse, error)                                                     // 获取所有角色
	GetRolePermissions(c *app.RequestContext, req *dto.GetRolePermissionsRequest) (*vo.RolePermissionsResponse, error) // 获取角色权限

	// 用户角色管理
	AssignRole(c *app.RequestContext, req *dto.AssignRoleRequest) (*vo.RoleOpResponse, error)        // 为用户分配角色
	RevokeRole(c *app.RequestContext, req *dto.RevokeRoleRequest) (*vo.RoleOpResponse, error)        // 撤销用户角色
	GetUserRoles(c *app.RequestContext, req *dto.GetUserRolesRequest) (*vo.UserRolesResponse, error) // 获取用户角色

	// 权限检查
	CheckPermission(c *app.RequestContext, req *dto.CheckPermissionRequest) (*vo.CheckResponse, error) // 权限检查
}
