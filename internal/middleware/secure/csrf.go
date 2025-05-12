// Package secure_middleware 提供安全相关中间件
// 创建者：Done-0
// 创建时间：2025-05-10
package secure_middleware

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

// InitCSRF 初始化 CSRF 中间件，使用默认配置
// 返回值：
//   - echo.MiddlewareFunc: Echo 框架中间件函数
func InitCSRF() echo.MiddlewareFunc {
	return csrfWithConfig(defaultCSRFConfig)
}

// csrfConfig 定义了 CSRF 中间件的配置
type csrfConfig struct {
	Skipper        func(echo.Context) bool // 用于跳过中间件的配置
	TokenLength    uint8                   // Token 的长度
	TokenLookup    string                  // Token 查找方式，默认 "header:X-CSRF-Token"
	ContextKey     string                  // 上下文存储 CSRF Token 的键
	CookieName     string                  // Cookie 名称
	CookiePath     string                  // Cookie 路径
	CookieDomain   string                  // Cookie 域
	CookieSecure   bool                    // 是否启用 Secure Cookie
	CookieHTTPOnly bool                    // 是否启用 HttpOnly Cookie
	CookieSameSite http.SameSite           // Cookie 的 SameSite 设置
	CookieMaxAge   int                     // Cookie 有效期，单位为秒
}

// defaultCSRFConfig 提供默认的 CSRF 配置
var defaultCSRFConfig = csrfConfig{
	Skipper:        func(c echo.Context) bool { return false },
	TokenLength:    32,                                // Token 默认长度 32 字节
	TokenLookup:    "header:" + echo.HeaderXCSRFToken, // 默认从 Header 查找 X-CSRF-Token
	ContextKey:     "csrf",                            // 上下文中的 CSRF Token 键
	CookieName:     "_csrf",                           // 默认 CSRF Cookie 名称
	CookiePath:     "/",                               // Cookie 默认路径
	CookieDomain:   "",                                // 默认不设置 Cookie 域
	CookieSecure:   false,                             // 默认不开启 Secure
	CookieHTTPOnly: true,                              // 默认启用 HttpOnly
	CookieSameSite: http.SameSiteLaxMode,              // 默认 SameSite 设置为 Lax
	CookieMaxAge:   86400,                             // Cookie 默认 24 小时有效期
}

// csrfWithConfig 使用传入的配置生成 CSRF 中间件
// 参数：
//   - config: CSRF 配置
//
// 返回值：
//   - echo.MiddlewareFunc: Echo 框架中间件函数
func csrfWithConfig(config csrfConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			token, err := getTokenFromRequest(c, config.TokenLookup)
			if err != nil || token == "" {
				token = generateCSRFToken(config.TokenLength)
				setCSRFCookie(c, config, token)
			} else {
				csrfCookie, err := c.Cookie(config.CookieName)
				if err != nil || csrfCookie.Value != token {
					return echo.NewHTTPError(http.StatusForbidden, "CSRF token 验证失败")
				}
			}

			c.Set(config.ContextKey, token)

			return next(c)
		}
	}
}

// getTokenFromRequest 从请求中获取 CSRF token
// 参数：
//   - c: Echo 上下文
//   - lookup: Token 查找方式
//
// 返回值：
//   - string: CSRF Token
//   - error: 获取过程中的错误
func getTokenFromRequest(c echo.Context, lookup string) (string, error) {
	parts := strings.Split(lookup, ":")
	if len(parts) != 2 {
		return "", errors.New("无效的 Token 查找方式")
	}
	switch parts[0] {
	case "header":
		return c.Request().Header.Get(parts[1]), nil
	case "form":
		return c.FormValue(parts[1]), nil
	case "query":
		return c.QueryParam(parts[1]), nil
	default:
		return "", errors.New("不支持的 Token 查找类型")
	}
}

// generateCSRFToken 生成随机的 CSRF token
// 参数：
//   - length: Token长度
//
// 返回值：
//   - string: 生成的 CSRF Token
func generateCSRFToken(length uint8) string {
	token := make([]byte, length)
	rand.Read(token)
	return base64.StdEncoding.EncodeToString(token)
}

// setCSRFCookie 设置 CSRF Token 到 Cookie
// 参数：
//   - c: Echo 上下文
//   - config: CSRF 配置
//   - token: CSRF Token
func setCSRFCookie(c echo.Context, config csrfConfig, token string) {
	cookie := &http.Cookie{
		Name:     config.CookieName,
		Value:    token,
		Path:     config.CookiePath,
		Domain:   config.CookieDomain,
		Secure:   config.CookieSecure,
		HttpOnly: config.CookieHTTPOnly,
		SameSite: config.CookieSameSite,
		Expires:  time.Now().Add(time.Duration(config.CookieMaxAge) * time.Second),
	}
	c.SetCookie(cookie)
}
