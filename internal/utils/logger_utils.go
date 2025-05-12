// Package utils 提供日志记录工具
// 创建者：Done-0
// 创建时间：2025-05-10
package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"jank.com/jank_blog/internal/global"
)

const (
	BIZLOG = "Bizlog" // 业务日志键名
)

// BizLogger 业务日志记录器
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - *logrus.Entry: 日志条目
func BizLogger(c echo.Context) *logrus.Entry {
	if bizLog, ok := c.Get(BIZLOG).(*logrus.Entry); ok {
		return bizLog
	}

	return logrus.NewEntry(global.SysLog)
}
