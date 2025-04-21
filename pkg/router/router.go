package router

import (
	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/pkg/router/routes"
)

// @title		Jank Blog API
// @version		1.0
// @description	This is the API documentation for Jank Blog.
// @host		localhost:9010
// @BasePath	/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description 输入格式: Bearer {token}
// New 函数用于注册应用程序的路由
func New(app *echo.Echo) {
	// 创建多版本 API 路由组
	api1 := app.Group("/api/v1")
	api2 := app.Group("/api/v2")

	// 注册测试相关的路由
	routes.RegisterTestRoutes(api1, api2)
	// 注册账户相关的路由
	routes.RegisterAccountRoutes(api1)
	// 注册验证相关的路由
	routes.RegisterVerificationRoutes(api1)
	// 注册文章相关的路由
	routes.RegisterPostRoutes(api1)
	// 注册类目相关的路由
	routes.RegisterCategoryRoutes(api1)
	// 注册评论相关的路由
	routes.RegisterCommentRoutes(api1)
}
