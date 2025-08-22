package routes

import (
	"log"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"

	"github.com/Done-0/jank/internal/middleware/jwt"
	"github.com/Done-0/jank/pkg/wire"
)

// RegisterThemeRoutes 注册主题相关路由（包含前端和API路由）
func RegisterThemeRoutes(h *server.Hertz, apiGroup *route.RouterGroup) {
	themeController, err := wire.NewThemeController()
	if err != nil {
		log.Fatalf("Failed to initialize theme controller: %v", err)
	}

	// === 前端路由 ===
	// Frontend 主题首页处理器
	h.GET("/", themeController.ServeHomePage)

	// Console 主题路由处理器
	h.GET("/console", themeController.ServeHomePage)
	h.GET("/console/*filepath", themeController.ServeStaticResource)

	// 静态资源处理器 - 处理所有未匹配的路径（跳过API路径）
	h.NoRoute(themeController.ServeStaticResource)

	// 主题路由组
	themeGroup := apiGroup.Group("/theme", jwt.New())
	{
		// POST 方法
		themeGroup.POST("/switch", themeController.SwitchTheme) // 切换主题

		// GET 方法
		themeGroup.GET("/get", themeController.GetActiveTheme) // 获取当前激活主题
		themeGroup.GET("/list", themeController.ListThemes)    // 列举所有主题
	}
}
