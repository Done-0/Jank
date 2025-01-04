package routes

import (
	"github.com/labstack/echo/v4"
	"jank.com/jank_blog/pkg/serve/controller/verification"
)

func RegisterVerificationRoutes(app *echo.Echo) {
	accountGroup := app.Group("/verification")
	accountGroup.GET("/genImgVerificationCode", verification.GenImgVerificationCode)
	accountGroup.GET("/sendEmailVerificationCode", verification.SendEmailVerificationCode)
}
