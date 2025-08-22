// Package logger 提供访问日志中间件配置
// 创建者：Done-0
// 创建时间：2025-08-05
package logger

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/logger/accesslog"
)

// New 创建访问日志中间件
// 返回值：
//
// app.HandlerFunc: 访问日志中间件
func New() app.HandlerFunc {
	return accesslog.New(accesslog.WithTimeFormat(time.RFC3339))
}
