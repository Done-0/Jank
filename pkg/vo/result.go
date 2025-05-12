// Package vo 提供视图对象定义和响应结果包装
// 创建者：Done-0
// 创建时间：2025-05-10
package vo

import (
	"errors"
	"time"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
)

// Result 通用 API 响应结果结构体
type Result struct {
	*bizErr.Err             // 错误信息
	Data        interface{} `json:"data"`      // 响应数据
	RequestId   interface{} `json:"requestId"` // 请求ID
	TimeStamp   interface{} `json:"timeStamp"` // 响应时间戳
}

// Success 成功返回
// 参数：
//   - c: Echo 上下文
//   - data: 响应数据
//
// 返回值：
//   - Result: 成功响应结果
func Success(c echo.Context, data interface{}) Result {
	return Result{
		Err:       nil,
		Data:      data,
		RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
		TimeStamp: time.Now().Unix(),
	}
}

// Fail 失败返回
// 参数：
//   - c: Echo 上下文
//   - data: 错误相关数据
//   - err: 错误对象
//
// 返回值：
//   - Result: 失败响应结果
func Fail(c echo.Context, data interface{}, err error) Result {
	var newBizErr *bizErr.Err
	if ok := errors.As(err, &newBizErr); ok {
		return Result{
			Err:       newBizErr,
			Data:      data,
			RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
			TimeStamp: time.Now().Unix(),
		}
	}

	return Result{
		Err:       bizErr.New(bizErr.SERVER_ERR),
		Data:      data,
		RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
		TimeStamp: time.Now().Unix(),
	}
}
