package routes

import (
	"log"

	"github.com/cloudwego/hertz/pkg/route"

	"github.com/Done-0/jank/internal/middleware/jwt"
	"github.com/Done-0/jank/pkg/wire"
)

// RegisterPluginRoutes 注册插件路由
func RegisterPluginRoutes(r *route.RouterGroup) {
	pluginController, err := wire.NewPluginController()
	if err != nil {
		log.Fatalf("Failed to initialize plugin controller: %v", err)
	}

	// 插件路由组
	pluginGroup := r.Group("/plugin", jwt.New())
	{
		// POST 方法
		pluginGroup.POST("/register", pluginController.RegisterPlugin)     // 注册插件
		pluginGroup.POST("/unregister", pluginController.UnregisterPlugin) // 注销插件
		pluginGroup.POST("/execute", pluginController.ExecutePlugin)       // 执行插件方法

		// GET 方法
		pluginGroup.GET("/get", pluginController.GetPlugin)    // 获取插件信息 ?plugin_id=xxx
		pluginGroup.GET("/list", pluginController.ListPlugins) // 列举所有插件
	}
}
