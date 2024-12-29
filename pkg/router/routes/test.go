package routes

import (
	"github.com/labstack/echo/v4"
	testapi "jank.com/jank_blog/pkg/serve/controller/test"
)

func RegisterTestRoutes(app *echo.Echo) {
	testGroup := app.Group("/test")
	testGroup.GET("/ping", testapi.Ping)
}
