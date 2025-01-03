package middleware

import (
	"github.com/labstack/echo/v4"
	request_middleware "github.com/labstack/echo/v4/middleware"

	"jank.com/jank_blog/internal/logger"
	cors_middleware "jank.com/jank_blog/internal/middleware/cors"
	error_middleware "jank.com/jank_blog/internal/middleware/error"
	recover_middleware "jank.com/jank_blog/internal/middleware/recover"
	secure_middleware "jank.com/jank_blog/internal/middleware/secure"
	swagger_middleware "jank.com/jank_blog/internal/middleware/swagger"
)

func InitMiddleware(app *echo.Echo) {
	// 设置全局错误处理
	app.Use(error_middleware.InitGlobalError())
	// 配置 CORS 中间件
	app.Use(cors_middleware.InitCORS())
	// 初始化 Swagger 中间件
	app.Use(swagger_middleware.InitSwagger())
	// 全局请求 ID 中间件
	app.Use(request_middleware.RequestID())
	// 日志中间件
	app.Use(logger.New())
	// 配置 xss 防御中间件
	app.Use(secure_middleware.InitXss())
	// 配置 csrf 防御中间件
	app.Use(secure_middleware.InitCSRF())
	// 全局异常恢复中间件
	app.Use(recover_middleware.InitRecover())
}
