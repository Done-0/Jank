// Package router 提供应用程序路由注册功能
// 创建者：Done-0
// 创建时间：2025-08-05
package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/Done-0/jank/pkg/router/routes"
)

// New 函数用于注册应用程序的路由
// 参数：
//
//	app: Hertz 路由引擎
func New(app *server.Hertz) {
	api := app.Group("/api/v1")

	// 注册用户相关的路由
	routes.RegisterUserRoutes(api)

	// 注册RBAC相关的路由
	routes.RegisterRBACRoutes(api)

	// 注册验证码相关的路由
	routes.RegisterVerificationRoutes(api)

	// 注册分类相关的路由
	routes.RegisterCategoryRoutes(api)

	// 注册文章相关的路由
	routes.RegisterPostRoutes(api)

	// 注册插件相关的路由
	routes.RegisterPluginRoutes(api)

	// 注册主题相关的路由
	routes.RegisterThemeRoutes(app, api)
}
