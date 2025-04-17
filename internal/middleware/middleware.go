package middleware

import (
	"github.com/labstack/echo/v4"

	requestMiddleware "github.com/labstack/echo/v4/middleware"
	corsMiddleware "jank.com/jank_blog/internal/middleware/cors"
	errorMiddleware "jank.com/jank_blog/internal/middleware/error"
	loggerMiddleware "jank.com/jank_blog/internal/middleware/logger"
	recoverMiddleware "jank.com/jank_blog/internal/middleware/recover"
	secureMiddleware "jank.com/jank_blog/internal/middleware/secure"
)

func New(app *echo.Echo) {
	// 设置全局错误处理
	app.Use(errorMiddleware.InitError())
	// 配置 CORS 中间件
	app.Use(corsMiddleware.InitCORS())
	// 全局请求 ID 中间件
	app.Use(requestMiddleware.RequestID())
	// 日志中间件
	app.Use(loggerMiddleware.InitLogger())
	// 配置 xss 防御中间件
	app.Use(secureMiddleware.InitXss())
	// 配置 csrf 防御中间件
	app.Use(secureMiddleware.InitCSRF())
	// 全局异常恢复中间件
	app.Use(recoverMiddleware.InitRecover())
	// 初始化 Swagger 中间件
	//app.Use(swaggerMiddleware.InitSwagger())
}
