package routes

import (
	"github.com/labstack/echo/v4"
	auth_middleware "jank.com/jank_blog/internal/middleware/auth"
	"jank.com/jank_blog/pkg/serve/controller/account"
)

func RegisterAccountRoutes(app *echo.Echo) {
	accountGroup := app.Group("/account")
	accountGroup.POST("/getAccount", account.GetAccount)
	accountGroup.POST("/registerAccount", account.RegisterAcc)
	accountGroup.POST("/loginAccount", account.LoginAccount)
	accountGroup.POST("/getUserProfile", account.GetUserProfile, auth_middleware.JWTMiddleware())
	accountGroup.POST("/logoutAccount", account.LogoutAccount, auth_middleware.JWTMiddleware())
}
