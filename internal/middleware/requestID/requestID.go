// Package requestID 提供请求 ID 中间件
// 创建者：Done-0
// 创建时间：2025-08-08
package requestID

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/hertz-contrib/requestid"

	"github.com/Done-0/jank/internal/types/consts"
)

// New 创建访问日志中间件
// 返回值：
//
// app.HandlerFunc: 访问日志中间件
func New() app.HandlerFunc {
	return requestid.New(
		requestid.WithGenerator(func(ctx context.Context, c *app.RequestContext) string {
			return uuid.New().String()
		}),
		requestid.WithCustomHeaderStrKey(consts.HeaderRequestID),
	)
}
