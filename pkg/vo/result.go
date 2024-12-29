package vo

import (
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	bizerr "jank.com/jank_blog/internal/error"
)

type Result struct {
	*bizerr.Err
	Data      interface{} `json:"data"`
	RequestId interface{} `json:"requestId"`
	TimeStamp interface{} `json:"timestamp"`
}

// 成功返回
func Success(data interface{}, c echo.Context) Result {
	return Result{
		Err:       nil,
		Data:      data,
		RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
		TimeStamp: time.Now().Unix(),
	}
}

// 失败返回
func Fail(data interface{}, err error, c echo.Context) Result {
	var bizErr *bizerr.Err
	if ok := errors.As(err, &bizErr); ok {
		return Result{
			Err:       bizErr,
			Data:      data,
			RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
			TimeStamp: time.Now().Unix(),
		}
	}

	return Result{
		Err:       bizerr.New(bizerr.ServerError),
		Data:      data,
		RequestId: c.Response().Header().Get(echo.HeaderXRequestID),
		TimeStamp: time.Now().Unix(),
	}
}
