package cors_middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/internal/global"
)

func InitCORS() echo.MiddlewareFunc {
	return corsWithConfig(defaultCORSConfig())
}

// CORSConfig 定义 CORS 中间件的配置
type corsConfig struct {
	AllowedOrigins   []string // 允许的源
	AllowedMethods   []string // 允许的方法
	AllowedHeaders   []string // 允许的头部
	AllowCredentials bool     // 是否允许携带证书
}

// DefaultCORSConfig 提供了默认的 CORS 配置
func defaultCORSConfig() corsConfig {
	return corsConfig{
		AllowedOrigins:   []string{"*"},                                                                                                   // 默认允许所有域名
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                                                             // 默认允许的请求方法
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Client-Info", "X-Client-Version", "X-Client-Data", "X-Request-Id"}, // 默认允许的请求头
		AllowCredentials: false,                                                                                                           // 默认不允许携带证书
	}
}

// corsWithConfig 返回一个 CORS 中间件函数
func corsWithConfig(config corsConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", strings.Join(config.AllowedOrigins, ","))
			c.Response().Header().Set("Access-Control-Allow-Methods", strings.Join(config.AllowedMethods, ","))
			c.Response().Header().Set("Access-Control-Allow-Headers", strings.Join(config.AllowedHeaders, ","))

			if config.AllowCredentials {
				c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			}

			// 处理预检请求，缓存预检请求结果 24 小时
			if c.Request().Method == "OPTIONS" {
				c.Set("Access-Control-Max-Age", "86400")
				return c.NoContent(http.StatusNoContent)
			}

			// 记录 CORS 请求
			if global.BizLog != nil {
				global.BizLog.Info("CORS request",
					"method", c.Request().Method,
					"path", c.Request().URL.Path,
					"origin", c.Request().Header.Get("Origin"),
				)
			}

			return next(c)
		}
	}
}
