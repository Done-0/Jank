package routes

import (
	"github.com/labstack/echo/v4"
	authMiddleware "jank.com/jank_blog/internal/middleware/auth"
	"jank.com/jank_blog/pkg/serve/controller/account"
)

func RegisterAccountRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	accountGroupV1 := apiV1.Group("/account")
	accountGroupV1.POST("/getAccount", account.GetAccount)
	accountGroupV1.POST("/registerAccount", account.RegisterAcc)
	accountGroupV1.POST("/loginAccount", account.LoginAccount)
	accountGroupV1.POST("/logoutAccount", account.LogoutAccount, authMiddleware.JWTMiddleware())
	accountGroupV1.POST("/resetPassword", account.ResetPassword, authMiddleware.JWTMiddleware())
}

func RegisterRolePermissionRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	// 角色管理
	roleGroup := apiV1.Group("/role")
	roleGroup.POST("/createOneRole", account.CreateRole)
	roleGroup.POST("/updateOneRole", account.UpdateRole)
	roleGroup.POST("/deleteOneRole", account.DeleteRole)
	roleGroup.GET("/listAllRoles", account.ListRoles)

	// 权限管理
	permissionGroup := apiV1.Group("/permission")
	permissionGroup.POST("/createOnePermission", account.CreatePermission)
	permissionGroup.POST("/updateOnePermission", account.UpdatePermission)
	permissionGroup.POST("/deleteOnePermission", account.DeletePermission)
	permissionGroup.GET("/listAllPermissions", account.ListPermissions)

	// 用户角色管理
	accRoleGroup := apiV1.Group("/acc-role")
	accRoleGroup.POST("/assignRoleToAcc", account.AssignRoleToAcc)
	accRoleGroup.POST("/updateRoleForAcc", account.UpdateRoleForAcc)
	accRoleGroup.POST("/deleteRoleFromAcc", account.DeleteRoleFromAcc)
	accRoleGroup.POST("/getRolesByAcc", account.GetRolesByAcc)

	// 角色权限管理
	rolePermissionGroup := apiV1.Group("/role-permission")
	rolePermissionGroup.POST("/assignPermissionToRole", account.AssignPermissionToRole)
	rolePermissionGroup.POST("/updatePermissionForRole", account.UpdatePermissionForRole)
	rolePermissionGroup.POST("/deletePermissionFromRole", account.DeletePermissionFromRole)
	rolePermissionGroup.POST("/getPermissionsByRole", account.GetPermissionsByRole)
}
