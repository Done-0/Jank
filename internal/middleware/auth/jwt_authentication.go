package authMiddleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/mapper"
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
	RefreshToken:  "Refresh_Token",
	UserCache:     "User_Cache",
}

// RBACConfig 定义了权限缓存前缀的配置
type RBACConfig struct {
	CachePrefix string
}

// DefaultRBACConfig 默认配置
var DefaultRBACConfig = RBACConfig{
	CachePrefix: "RBAC_Permission",
}

// AuthMiddleware 处理 JWT 认证和权限校验
// requiredPermissionIDs 参数 :
//   - 若传入权限 ID，则在 JWT 认证通过后，校验当前角色是否拥有【至少一个】对应权限；
//   - 若未传入权限 ID，则仅进行 JWT 认证，不校验权限。
func AuthMiddleware(requiredPermissionIDs ...int64) echo.MiddlewareFunc {
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

			// 从 Token 中解析 accountID 和 roleID
			accountID, roleID, err := utils.ParseAccountAndRoleIDFromJWT(tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "无效的 Access Token，请重新登录")
			}

			//  验证数据库中的用户和角色是否匹配
			accountRole, err := mapper.GetRoleByAccountID(accountID)
			if err != nil || accountRole.RoleID != roleID {
				userCacheKey := fmt.Sprintf("%s:%d:%d", DefaultJWTConfig.UserCache, accountID, roleID)
				if delErr := global.RedisClient.Del(c.Request().Context(), userCacheKey).Err(); delErr != nil {
					global.BizLog.Errorf("删除用户缓存失败 [%s]: %v", userCacheKey, delErr)
				}
				return echo.NewHTTPError(http.StatusUnauthorized, "用户角色发生变更，请重新登录")
			}

			sessionCacheKey := fmt.Sprintf("%s:%d:%d", DefaultJWTConfig.UserCache, accountID, roleID)
			if sessionVal, err := global.RedisClient.Get(c.Request().Context(), sessionCacheKey).Result(); err != nil || sessionVal == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "无效会话，请重新登录")
			}

			// 如果未传入权限 ID，则仅进行 JWT 认证
			if len(requiredPermissionIDs) == 0 {
				return next(c)
			}

			// 校验当前角色是否拥有【至少一个】对应权限
			rolePermissions, err := mapper.GetPermissionsByRole(fmt.Sprint(roleID))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "获取权限信息失败")
			}

			hasPermission := false
			for _, rp := range rolePermissions {
				for _, reqPerm := range requiredPermissionIDs {
					if rp.PermissionID == reqPerm {
						hasPermission = true
						break
					}
				}
				if hasPermission {
					break
				}
			}
			if !hasPermission {
				permCacheKey := fmt.Sprintf("%s:%d", DefaultRBACConfig.CachePrefix, roleID)
				if delErr := global.RedisClient.Del(c.Request().Context(), permCacheKey).Err(); delErr != nil {
					global.BizLog.Errorf("删除权限缓存失败 [%s]: %v", permCacheKey, delErr)
				}
				return echo.NewHTTPError(http.StatusForbidden, "权限不足，请联系管理员")
			}

			return next(c)
		}
	}
}
