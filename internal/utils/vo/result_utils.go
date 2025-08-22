// Package vo 提供通用值对象
package vo

import (
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/requestid"

	"github.com/Done-0/jank/internal/types/errno"
	"github.com/Done-0/jank/internal/utils/errorx"
)

// Result 通用 API 响应结构体
type Result struct {
	Error     *Error `json:"error,omitempty"`
	Data      any    `json:"data,omitempty"`
	RequestId string `json:"requestId"`
	TimeStamp int64  `json:"timeStamp"`
}

// Error Microsoft API 风格错误信息
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Success 成功返回
// 参数：
//
//	c: Hertz 上下文
//	data: 响应数据
//
// 返回值：
//
//	Result: 通用 API 响应结构体
func Success(c *app.RequestContext, data any) Result {
	return Result{
		Data:      data,
		RequestId: requestid.Get(c),
		TimeStamp: time.Now().Unix(),
	}
}

// Fail 失败返回
func Fail(c *app.RequestContext, data any, err error) Result {
	if errData, ok := data.(error); ok {
		data = errData.Error()
	}

	if statusErr, ok := err.(errorx.StatusError); ok {
		return Result{
			Error:     &Error{Code: fmt.Sprintf("%d", statusErr.Code()), Message: statusErr.Msg()},
			Data:      data,
			RequestId: requestid.Get(c),
			TimeStamp: time.Now().Unix(),
		}
	}

	return Result{
		Error:     &Error{Code: fmt.Sprintf("%d", errno.ErrInternalServer), Message: err.Error()},
		Data:      data,
		RequestId: requestid.Get(c),
		TimeStamp: time.Now().Unix(),
	}
}
