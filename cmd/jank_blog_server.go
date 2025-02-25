package cmd

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/banner"
	"jank.com/jank_blog/internal/db"
	"jank.com/jank_blog/internal/middleware"
	"jank.com/jank_blog/internal/redis"
	"jank.com/jank_blog/pkg/router"
)

// Start 启动服务
func Start() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("程序启动时加载配置失败: %v", err)
		return
	}

	// 初始化 echo 实例
	app := echo.New()
	app.HideBanner = true
	banner.InitBanner()

	// 初始化中间件
	middleware.InitMiddleware(app)

	// 初始化数据库连接并自动迁移模型
	db.New(config)

	// 初始化 Redis 连接
	redis.New(config)

	// 注册路由
	router.RegisterRoutes(app)

	// 启动服务
	app.Logger.Fatal(app.Start(fmt.Sprintf("%s:%s", config.AppConfig.AppHost, config.AppConfig.AppPort)))
}
