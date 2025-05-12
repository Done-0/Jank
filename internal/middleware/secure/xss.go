// Package secure_middleware 提供安全相关中间件
// 创建者：Done-0
// 创建时间：2025-05-10
package secure_middleware

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// InitXss 返回一个 XSS 防护中间件，使用默认配置
// 返回值：
//   - echo.MiddlewareFunc: Echo 框架中间件函数
func InitXss() echo.MiddlewareFunc {
	return xssWithConfig(defaultXSSConfig)
}

// xssConfig 用于配置 XSS 防护中间件
type xssConfig struct {
	Skipper               func(echo.Context) bool // 用于跳过中间件的配置
	XSSPrevention         string                  // X-XSS-Protection 头部配置
	ContentTypeNosniff    string                  // X-Content-Type-Options 头部配置
	XFrameOptions         string                  // X-Frame-Options 头部配置
	HSTSMaxAge            int                     // Strict-Transport-Security 头部配置
	HSTSExcludeSubdomains bool                    // 是否排除子域名的 HSTS 配置
	ContentSecurityPolicy string                  // Content-Security-Policy 头部配置
}

// defaultXSSConfig 默认的 XSS 防护配置
var defaultXSSConfig = xssConfig{
	Skipper:               func(c echo.Context) bool { return false }, // 默认不跳过
	XSSPrevention:         "1; mode=block",                            // 开启 XSS 防护
	ContentTypeNosniff:    "nosniff",                                  // 禁止浏览器自动猜测内容类型
	XFrameOptions:         "SAMEORIGIN",                               // 允许来自同一来源的嵌入式框架
	HSTSMaxAge:            0,                                          // 只能通过HTTPS来访问的时间(单位秒)
	HSTSExcludeSubdomains: false,                                      // 是否排除子域名的 HSTS 配置
	ContentSecurityPolicy: "",                                         // Content-Security-Policy 头部配置
}

// xssWithConfig 返回一个 XSS 防护中间件函数
// 参数：
//   - config: XSS 防护配置
//
// 返回值：
//   - echo.MiddlewareFunc: Echo 框架中间件函数
func xssWithConfig(config xssConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = defaultXSSConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			// 设置 X-XSS-Protection 头部
			c.Response().Header().Set("X-XSS-Protection", config.XSSPrevention)
			// 设置 X-Content-Type-Options 头部
			c.Response().Header().Set("X-Content-Type-Options", config.ContentTypeNosniff)
			// 设置 X-Frame-Options 头部
			c.Response().Header().Set("X-Frame-Options", config.XFrameOptions)
			// 设置 Strict-Transport-Security 头部
			if config.HSTSMaxAge > 0 {
				hstsHeader := "max-age=" + strconv.Itoa(config.HSTSMaxAge)
				if config.HSTSExcludeSubdomains {
					hstsHeader += "; includeSubDomains"
				}
				c.Response().Header().Set("Strict-Transport-Security", hstsHeader)
			}
			// 设置 Content-Security-Policy 头部
			if config.ContentSecurityPolicy != "" {
				c.Response().Header().Set("Content-Security-Policy", config.ContentSecurityPolicy)
			}

			return next(c)
		}
	}
}
