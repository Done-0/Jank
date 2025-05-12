// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-05-10
package routes

import (
	"github.com/labstack/echo/v4"

	auth_middleware "jank.com/jank_blog/internal/middleware/auth"
	"jank.com/jank_blog/pkg/serve/controller/account"
)

// RegisterAccountRoutes 注册账户相关路由
// 参数：
//   - r: Echo 路由组数组，r[0] 为 API v1 版本组
func RegisterAccountRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	accountGroupV1 := apiV1.Group("/account")
	accountGroupV1.POST("/getAccount", account.GetAccount, auth_middleware.AuthMiddleware())
	accountGroupV1.POST("/registerAccount", account.RegisterAcc)
	accountGroupV1.POST("/loginAccount", account.LoginAccount)
	accountGroupV1.POST("/logoutAccount", account.LogoutAccount, auth_middleware.AuthMiddleware())
	accountGroupV1.POST("/resetPassword", account.ResetPassword, auth_middleware.AuthMiddleware())
}
