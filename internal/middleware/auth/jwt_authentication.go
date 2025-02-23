package authMiddleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/utils"
)

// JWTConfig 用于配置 JWT 中间件
type JWTConfig struct {
	Authorization string
	TokenPrefix   string
	RefreshToken  string
	UserCache     string
}

// DefaultJWTConfig 提供默认的 JWT 配置
var DefaultJWTConfig = JWTConfig{
	Authorization: "Authorization",
	TokenPrefix:   "Bearer ",
	RefreshToken:  "Refresh_Token",
	UserCache:     "User_Cache",
}

// JWTMiddleware 用于处理请求的 JWT 验证和 Token 缓存管理
func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get(DefaultJWTConfig.Authorization)
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "缺少 Authorization 请求头")
			}

			// 提取 token
			tokenString := strings.TrimPrefix(authHeader, DefaultJWTConfig.TokenPrefix)
			_, err := utils.ValidateJWTToken(tokenString, false)
			if err != nil {
				// 尝试刷新 access token
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
				return next(c)
			}

			// 从 token 中提取 accountID 和 roleID
			accountID, roleID, err := utils.ParseAccountAndRoleIDFromJWT(tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "无效的 Access Token，请重新登录")
			}

			// 检查 Redis 缓存中的会话状态
			cacheKey := fmt.Sprintf("%s:%d:%d", DefaultJWTConfig.UserCache, accountID, roleID)
			val, err := global.RedisClient.Get(c.Request().Context(), cacheKey).Result()

			if err != nil && err.Error() != "redis: nil" {
				return echo.NewHTTPError(http.StatusUnauthorized, "无效的会话，请重新登录")
			} else if val == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "无效的会话，请重新登录")
			}

			return next(c)
		}
	}
}
