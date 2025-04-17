package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"jank.com/jank_blog/internal/global"
)

// InitLogger 创建日志中间件
func InitLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqId := c.Response().Header().Get(echo.HeaderXRequestID)
			if reqId == "" {
				reqId = middleware.DefaultRequestIDConfig.Generator()
				c.Response().Header().Set(echo.HeaderXRequestID, reqId)
			}
			bizLog := global.SysLog.WithFields(logrus.Fields{
				"requestId": reqId,
				"requestIp": c.RealIP(),
			})
			// 将 BizLog 存储到当前请求上下文中
			c.Set("BizLog", bizLog)
			return next(c)
		}
	}
}
