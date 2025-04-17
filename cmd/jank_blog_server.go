package cmd

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/banner"
	"jank.com/jank_blog/internal/db"
	"jank.com/jank_blog/internal/logger"
	"jank.com/jank_blog/internal/middleware"
	"jank.com/jank_blog/internal/redis"
	"jank.com/jank_blog/pkg/router"
)

// Start 启动服务
func Start() {
	if err := configs.Init(configs.DefaultConfigPath); err != nil {
		log.Fatalf("配置初始化失败: %v", err)
		return
	}

	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("获取配置失败: %v", err)
		return
	}

	// 初始化 Logger
	logger.New()

	// 初始化 echo 实例
	app := echo.New()

	// 初始化 Banner
	banner.New(app)

	// 初始化中间件
	middleware.New(app)

	// 初始化数据库连接并自动迁移模型
	db.New(config)

	// 初始化 Redis 连接
	redis.New(config)

	// 注册路由
	router.New(app)

	// 启动服务
	app.Logger.Fatal(app.Start(fmt.Sprintf("%s:%s", config.AppConfig.AppHost, config.AppConfig.AppPort)))
}
