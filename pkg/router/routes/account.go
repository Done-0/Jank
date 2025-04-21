package routes

import (
	"github.com/labstack/echo/v4"

	auth_middleware "jank.com/jank_blog/internal/middleware/auth"
	"jank.com/jank_blog/pkg/serve/controller/account"
)

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
