package auth_middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/account"
)

// JWTConfig 用于配置 JWT 中间件
type JWTConfig struct {
	Authorization   string
	TokenPrefix     string
	RefreshToken    string
	RedisPrefix     string
	LocalsUserIdKey string
	LocalsEmailKey  string
}

// DefaultJWTConfig 提供默认的 JWT 配置
var DefaultJWTConfig = JWTConfig{
	Authorization:   "Authorization",
	TokenPrefix:     "Bearer ",
	RefreshToken:    "Refresh_Token",
	RedisPrefix:     "ACC_AUTH_TOKEN_CACHE_PREFIX",
	LocalsUserIdKey: "Locals_User_Id",
	LocalsEmailKey:  "Locals_Email",
}

func JWTMiddleware() echo.MiddlewareFunc {
	config := DefaultJWTConfig

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get(DefaultJWTConfig.Authorization)
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "缺少 Authorization Header 身份验证请求头")
			}

			tokenString := strings.TrimPrefix(authHeader, config.TokenPrefix)

			// 验证 Access Token
			token, err := utils.ValidateJWTToken(tokenString, false)
			if err != nil {
				return handleTokenRefresh(c, config)
			}

			// 将验证后的用户信息设置到上下文中
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userId := int64(claims[DefaultJWTConfig.LocalsUserIdKey].(float64))
				email := claims[DefaultJWTConfig.LocalsEmailKey].(string)
				c.Set(account.LocalsUserIdKey, userId)
				c.Set(account.LocalsEmailKey, email)

				if err := processToken(c, token, tokenString, config); err != nil {
					return err
				}
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "Access Token 无效")
			}

			return next(c)
		}
	}
}

// handleTokenRefresh 尝试使用 Refresh Token 刷新 Access Token
func handleTokenRefresh(c echo.Context, config JWTConfig) error {
	refreshHeader := c.Request().Header.Get(DefaultJWTConfig.RefreshToken)
	if refreshHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "无效的 Access Token")
	}

	refreshTokenString := strings.TrimPrefix(refreshHeader, config.TokenPrefix)
	newTokens, refreshErr := RefreshTokenLogic(refreshTokenString)
	if refreshErr != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "无效的 Access 和 Refresh Token")
	}

	c.Response().Header().Set(DefaultJWTConfig.Authorization, config.TokenPrefix+newTokens["accessToken"])
	c.Response().Header().Set(DefaultJWTConfig.RefreshToken, config.TokenPrefix+newTokens["refreshToken"])
	return nil
}

// processToken 处理有效的 Access Token
func processToken(c echo.Context, token *jwt.Token, tokenString string, config JWTConfig) error {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := int64(claims[DefaultJWTConfig.LocalsUserIdKey].(float64))
		c.Set(account.LocalsUserIdKey, userId)

		redisKey := config.RedisPrefix + strconv.FormatInt(userId, 10)
		exp := claims["exp"].(float64)
		expireTime := time.Until(time.Unix(int64(exp), 0))

		err := global.Redis.Set(context.Background(), redisKey, tokenString, expireTime).Err()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "无法存储 Token 到 Redis")
		}
	} else {
		utils.BizLogger(c).Errorf("Access Token 无效或 claims 解析失败")
		return echo.NewHTTPError(http.StatusUnauthorized, "Access Token 验证失败")
	}

	return nil
}

// RefreshToken 处理刷新 JWT 的请求
func RefreshToken(c echo.Context) error {
	authHeader := c.Request().Header.Get(DefaultJWTConfig.Authorization)
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Refresh Token 缺失")
	}

	refreshTokenString := strings.TrimPrefix(authHeader, DefaultJWTConfig.TokenPrefix)

	tokens, err := RefreshTokenLogic(refreshTokenString)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Refresh Token 验证失败")
	}

	return c.JSON(http.StatusOK, tokens)
}

// RefreshTokenLogic 负责刷新 Token
func RefreshTokenLogic(refreshTokenString string) (map[string]string, error) {
	token, err := utils.ValidateJWTToken(refreshTokenString, true)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := uint(claims[DefaultJWTConfig.LocalsUserIdKey].(float64))
		email := claims[DefaultJWTConfig.LocalsEmailKey].(string)

		newAccessToken, newRefreshToken, err := utils.GenerateJWT(userId, email)
		if err != nil {
			return nil, err
		}

		return map[string]string{
			"accessToken":  newAccessToken,
			"refreshToken": newRefreshToken,
		}, nil
	}

	return nil, fmt.Errorf("token 验证失败")
}
