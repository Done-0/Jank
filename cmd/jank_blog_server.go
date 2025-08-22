// Package cmd 提供应用程序的启动和运行入口
// 创建者：Done-0
// 创建时间：2025-08-05
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/casbin"
	"github.com/Done-0/jank/internal/db"
	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/logger"
	"github.com/Done-0/jank/internal/middleware"
	"github.com/Done-0/jank/internal/plugin"
	"github.com/Done-0/jank/internal/redis"
	"github.com/Done-0/jank/internal/theme"
	"github.com/Done-0/jank/pkg/router"
)

// Start 启动服务
func Start() {
	// 初始化配置
	if err := configs.New(configs.DefaultConfigPath); err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	cfgs, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	// 初始化日志
	logger.New(cfgs)

	// 初始化数据库
	db.New(cfgs)

	// 初始化 Redis
	redis.New(cfgs)

	// 初始化 Casbin 权限系统
	casbin.New(cfgs)

	// 初始化插件系统
	plugin.New(cfgs)

	// 初始化主题系统
	theme.New(cfgs)

	// 创建 Hertz 服务器实例
	addr := fmt.Sprintf("%s:%s", cfgs.AppConfig.AppHost, cfgs.AppConfig.AppPort)
	h := server.Default(
		server.WithHostPorts(addr),
		server.WithExitWaitTime(10*time.Second),
	)

	// 注册中间件
	middleware.New(h)

	// 注册路由
	router.New(h)

	// 注册优雅关闭钩子
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		plugin.GlobalPluginManager.Shutdown()
		theme.GlobalThemeManager.Shutdown()
	})

	// 启动信息
	global.SysLog.Infof("⇨ Hertz server starting on %s", addr)

	// 优雅关闭
	h.Spin()
}
