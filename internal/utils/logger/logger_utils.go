// Package logger 提供日志记录工具
// 创建者：Done-0
// 创建时间：2025-08-08
package logger

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/sirupsen/logrus"

	"github.com/Done-0/jank/internal/global"
)

const (
	BIZLOG = "Bizlog" // 业务日志键名
)

// BizLogger 业务日志记录器
// 参数：
//
//	c: Hertz 上下文
//
// 返回值：
//
//	*logrus.Entry: 日志条目
func BizLogger(c *app.RequestContext) *logrus.Entry {
	if bizLog, ok := c.Get(BIZLOG); ok {
		if entry, ok := bizLog.(*logrus.Entry); ok {
			return entry
		}
	}

	return logrus.NewEntry(global.SysLog)
}
