// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-08-05
package routes

import (
	"log"

	"github.com/cloudwego/hertz/pkg/route"

	"github.com/Done-0/jank/internal/middleware/jwt"
	"github.com/Done-0/jank/pkg/wire"
)

// RegisterRBACRoutes 注册RBAC相关路由
func RegisterRBACRoutes(r *route.RouterGroup) {
	rbacController, err := wire.NewRBACController()
	if err != nil {
		log.Fatalf("Failed to initialize RBAC controller: %v", err)
	}

	// RBAC 路由组
	rbacGroup := r.Group("/rbac", jwt.New())
	{
		// 权限管理
		rbacGroup.POST("/create-permission", rbacController.CreatePermission) // 创建权限
		rbacGroup.POST("/delete-permission", rbacController.DeletePermission) // 删除权限
		rbacGroup.POST("/assign-permission", rbacController.AssignPermission) // 为角色分配权限
		rbacGroup.POST("/revoke-permission", rbacController.RevokePermission) // 撤销角色权限
		rbacGroup.GET("/list-permissions", rbacController.ListPermissions)    // 获取所有权限

		// 角色管理
		rbacGroup.POST("/create-role", rbacController.CreateRole)                 // 创建角色
		rbacGroup.POST("/delete-role", rbacController.DeleteRole)                 // 删除角色
		rbacGroup.GET("/list-roles", rbacController.ListRoles)                    // 获取所有角色
		rbacGroup.GET("/get-role-permissions", rbacController.GetRolePermissions) // 获取角色权限

		// 用户角色管理
		rbacGroup.POST("/assign-role", rbacController.AssignRole)     // 为用户分配角色
		rbacGroup.POST("/revoke-role", rbacController.RevokeRole)     // 撤销用户角色
		rbacGroup.GET("/get-user-roles", rbacController.GetUserRoles) // 获取用户角色

		// 权限检查
		rbacGroup.POST("/check-permission", rbacController.CheckPermission) // 权限检查
	}
}
