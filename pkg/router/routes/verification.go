package routes

import (
	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/pkg/serve/controller/verification"
)

func RegisterVerificationRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	accountGroupV1 := apiV1.Group("/verification")
	accountGroupV1.GET("/sendImgVerificationCode", verification.SendImgVerificationCode)
	accountGroupV1.GET("/sendEmailVerificationCode", verification.SendEmailVerificationCode)
}
