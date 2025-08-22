// Package middleware 提供通用中间件
// 创建者：Done-0
// 创建时间：2025-08-05
package middleware

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/Done-0/jank/internal/middleware/cors"
	"github.com/Done-0/jank/internal/middleware/logger"
	"github.com/Done-0/jank/internal/middleware/requestID"
)

// New 初始化并注册所有中间件
// 参数
//
// h *server.Hertz: Hertz 服务器实例
func New(h *server.Hertz) {
	// 请求 ID 中间件
	h.Use(requestID.New())

	// 访问日志中间件
	h.Use(logger.New())

	// CORS 中间件
	h.Use(cors.New())
}
