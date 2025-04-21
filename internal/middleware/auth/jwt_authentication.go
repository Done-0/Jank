package auth_middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/utils"
)

// JWTConfig 定义了 Token 相关的配置
type JWTConfig struct {
	Authorization string
	TokenPrefix   string
	RefreshToken  string
	UserCache     string
}

// DefaultJWTConfig 默认配置
var DefaultJWTConfig = JWTConfig{
	Authorization: "Authorization",
	TokenPrefix:   "Bearer ",
	RefreshToken:  "REFRESH_TOKEN",
	UserCache:     "USER_CACHE",
}

// AuthMiddleware 处理 JWT 认证中间件
func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 从请求头中提取 Access Token
			authHeader := c.Request().Header.Get(DefaultJWTConfig.Authorization)
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "缺少 Authorization 请求头")
			}
			tokenString := strings.TrimPrefix(authHeader, DefaultJWTConfig.TokenPrefix)

			// 验证 JWT Token；若验证失败则尝试使用 Refresh Token 刷新
			_, err := utils.ValidateJWTToken(tokenString, false)
			if err != nil {
				refreshHeader := c.Request().Header.Get(DefaultJWTConfig.RefreshToken)
				if refreshHeader == "" {
					return echo.NewHTTPError(http.StatusUnauthorized, "无效 Access Token，请重新登录")
				}
				refreshTokenString := strings.TrimPrefix(refreshHeader, DefaultJWTConfig.TokenPrefix)
				newTokens, refreshErr := utils.RefreshTokenLogic(refreshTokenString)
				if refreshErr != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, "无效 Access 和 Refresh Token，请重新登录")
				}
				c.Response().Header().Set(DefaultJWTConfig.Authorization, DefaultJWTConfig.TokenPrefix+newTokens["accessToken"])
				c.Response().Header().Set(DefaultJWTConfig.RefreshToken, DefaultJWTConfig.TokenPrefix+newTokens["refreshToken"])
				tokenString = newTokens["accessToken"]
			}

			// 从 Token 中解析 accountID
			accountID, err := utils.ParseAccountAndRoleIDFromJWT(tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "无效的 Access Token，请重新登录")
			}

			sessionCacheKey := fmt.Sprintf("%s:%d", DefaultJWTConfig.UserCache, accountID)
			if sessionVal, err := global.RedisClient.Get(c.Request().Context(), sessionCacheKey).Result(); err != nil || sessionVal == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "无效会话，请重新登录")
			}

			return next(c)
		}
	}
}
