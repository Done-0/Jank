// Package rbac 提供 RBAC 权限控制中间件
// 创建者：Done-0
// 创建时间：2025-08-10
package rbac

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/casbin"

	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/types/consts"
)

// RequirePermission 创建需要特定权限的中间件
func RequirePermission(permissions string, opts ...casbin.Option) app.HandlerFunc {
	authMiddleware, err := casbin.NewCasbinMiddlewareFromEnforcer(global.Enforcer, func(ctx context.Context, c *app.RequestContext) string {
		if userID, exists := c.Get(consts.JWTSubjectClaim); exists {
			if id, ok := userID.(int64); ok {
				return fmt.Sprintf("%d", id)
			}
		}
		return ""
	})
	if err != nil {
		return nil
	}
	return authMiddleware.RequiresPermissions(permissions, opts...)
}

// RequireRole 创建需要特定角色的中间件
func RequireRole(roles string, opts ...casbin.Option) app.HandlerFunc {
	authMiddleware, err := casbin.NewCasbinMiddlewareFromEnforcer(global.Enforcer, func(ctx context.Context, c *app.RequestContext) string {
		if userID, exists := c.Get(consts.JWTSubjectClaim); exists {
			if id, ok := userID.(int64); ok {
				return fmt.Sprintf("%d", id)
			}
		}
		return ""
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create RBAC middleware: %v", err))
	}
	return authMiddleware.RequiresRoles(roles, opts...)
}
