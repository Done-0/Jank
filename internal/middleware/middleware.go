package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"jank.com/jank_blog/internal/error/error_handler"
	"jank.com/jank_blog/internal/logger"
	cors_middleware "jank.com/jank_blog/internal/middleware/cors"
	recover_middleware "jank.com/jank_blog/internal/middleware/recover"
	"jank.com/jank_blog/internal/middleware/secure"
	swagger_middleware "jank.com/jank_blog/internal/middleware/swagger"
)

func InitMiddleware(app *echo.Echo) {
	// 设置全局错误处理函数
	app.HTTPErrorHandler = func(err error, c echo.Context) {
		error_handler.HandleGlobalError(err, c)
	}
	// 配置 CORS 中间件
	app.Use(cors_middleware.CORS())
	// 初始化 Swagger 中间件
	app.Use(swagger_middleware.InitSwagger())
	// 全局请求 ID 中间件
	app.Use(middleware.RequestID())
	// 日志中间件
	app.Use(logger.New())
	// 配置 xss 防御中间件
	app.Use(secure.InitXss())
	// 配置 csrf 防御中间件
	app.Use(secure.InitCSRF())
	// 全局异常恢复中间件
	app.Use(recover_middleware.InitRecover())
}
