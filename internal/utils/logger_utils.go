package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"jank.com/jank_blog/internal/global"
)

const (
	bizlog = "Bizlog"
)

// BizLogger 业务日志记录器
func BizLogger(c echo.Context) *logrus.Entry {
	if bizLog, ok := c.Get(bizlog).(*logrus.Entry); ok {
		return bizLog
	}

	return logrus.NewEntry(global.SysLog)
}
