package error_handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	bizerr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/pkg/vo"
)

// HandleGlobalError 全局错误处理函数
func HandleGlobalError(err error, c echo.Context) error {
	code := http.StatusInternalServerError

	var e *bizerr.Err
	if errors.As(err, &e) {
		code = e.Code
	}

	// 捕获请求信息：请求方法、请求URI、客户端IP、User-Agent
	requestMethod := c.Request().Method
	requestURI := c.Request().RequestURI
	clientIP := c.Request().RemoteAddr
	userAgent := c.Request().UserAgent()

	logMessage := fmt.Sprintf("请求异常: %v | Method: %s | URI: %s | IP: %s | User-Agent: %s", err, requestMethod, requestURI, clientIP, userAgent)
	global.SysLog.Errorf(logMessage)

	return c.JSON(code, vo.Fail(nil, bizerr.New(code, err.Error()), c))
}
