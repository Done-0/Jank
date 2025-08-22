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

// RegisterUserRoutes 注册用户相关路由
// 参数：
//
//	r: Hertz 路由组，API v1 版本组
func RegisterUserRoutes(r *route.RouterGroup) {
	userController, err := wire.NewUserController()
	if err != nil {
		log.Fatalf("Failed to initialize user controller: %v", err)
	}

	// 用户路由组
	userGroup := r.Group("/user")
	{
		// 公开接口（无需认证）
		userGroup.POST("/register", userController.Register)          // 用户注册
		userGroup.POST("/login", userController.Login)                // 用户登录
		userGroup.POST("/refresh-token", userController.RefreshToken) // 刷新 token

		// 需要认证的接口
		userGroup.POST("/logout", jwt.New(), userController.Logout)                // 用户登出
		userGroup.POST("/update", jwt.New(), userController.Update)                // 更新用户信息
		userGroup.POST("/reset-password", jwt.New(), userController.ResetPassword) // 重置密码

		userGroup.GET("/profile", jwt.New(), userController.GetProfile) // 获取用户资料
		userGroup.GET("/list", userController.ListUsers)                // 获取用户列表（管理员）

		userGroup.POST("/role", jwt.New(), userController.UpdateUserRole) // 更新用户角色（管理员）
	}
}
