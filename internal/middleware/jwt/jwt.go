// Package jwt 提供 JWT 认证中间件
// 创建者：Done-0
// 创建时间：2025-08-10
package jwt

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/types/errno"
	"github.com/Done-0/jank/internal/utils/errorx"
	"github.com/Done-0/jank/internal/utils/vo"

	constants "github.com/Done-0/jank/internal/types/consts"
)

// New 创建 JWT 认证中间件
func New() app.HandlerFunc {
	cfgs, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}
	jwtConfig := cfgs.AppConfig.JWT

	// 创建 JWT 中间件配置
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       constants.JWTRealm,
		Key:         []byte(jwtConfig.Secret),
		Timeout:     time.Duration(jwtConfig.ExpireTime) * time.Hour,
		MaxRefresh:  time.Duration(jwtConfig.RefreshExpire) * time.Hour,
		IdentityKey: constants.JWTSubjectClaim,
		PayloadFunc: func(data any) jwt.MapClaims {
			if userID, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.JWTSubjectClaim: userID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) any {
			claims := jwt.ExtractClaims(ctx, c)
			if userID, exists := claims[constants.JWTSubjectClaim]; exists {
				if id, ok := userID.(float64); ok {
					return int64(id)
				}
			}
			return nil
		},
		Authorizator: func(data any, ctx context.Context, c *app.RequestContext) bool {
			if userID, ok := data.(int64); ok && userID > 0 {
				authHeader := string(c.GetHeader(constants.HeaderAuthorization))
				if !strings.HasPrefix(authHeader, constants.JWTBearerPrefix) {
					return false
				}
				currentToken := strings.TrimPrefix(authHeader, constants.JWTBearerPrefix)

				cacheKey := fmt.Sprintf("%s:%d", constants.AuthAccessTokenKeyPrefix, userID)
				cachedToken := global.RedisClient.Get(ctx, cacheKey).Val()
				if cachedToken == "" || cachedToken != currentToken {
					return false
				}

				c.Set(constants.JWTSubjectClaim, userID)
				return true
			}
			return false
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(consts.StatusUnauthorized, vo.Fail(c, nil, errorx.New(errno.ErrUnauthorized, errorx.KV("msg", message))))
		},
		TokenLookup:   constants.JWTTokenLookup,
		TokenHeadName: constants.JWTTokenHeadName,
		TimeFunc:      time.Now,
	})
	if err != nil {
		panic(fmt.Sprintf("JWT 中间件初始化失败: %v", err))
	}

	return authMiddleware.MiddlewareFunc()
}
