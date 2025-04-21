package error_middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	bizerr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/pkg/vo"
)

// InitError 全局错误处理中间件
func InitError() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				code := http.StatusInternalServerError
				var e *bizerr.Err
				if errors.As(err, &e) {
					code = e.Code
				}

				// 捕获请求信息：请求方法、请求URI、客户端IP、User-Agent
				requestMethod := c.Request().Method
				requestURI := c.Request().RequestURI
				clientIP := c.Request().RemoteAddr
				userAgent := c.Request().UserAgent()

				// 构建日志消息
				logMessage := fmt.Sprintf("请求异常: %v | Method: %s | URI: %s | IP: %s | User-Agent: %s", err, requestMethod, requestURI, clientIP, userAgent)
				global.SysLog.Error(logMessage)

				return c.JSON(code, vo.Fail(nil, bizerr.New(code, err.Error()), c))
			}
			return nil
		}
	}
}
