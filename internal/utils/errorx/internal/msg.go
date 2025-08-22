package internal

import (
	"fmt"
)

// withMessage 带消息的错误包装器
type withMessage struct {
	cause error  // 原因错误
	msg   string // 消息
}

// Unwrap 返回被包装的原始错误
//
// 返回值：
//
//	error: 原始错误
func (w *withMessage) Unwrap() error {
	return w.cause
}

// Error 错误字符串表示
//
// 返回值：
//
//	string: 包含消息和原因的错误字符串
func (w *withMessage) Error() string {
	return fmt.Sprintf("%s\ncause=%s", w.msg, w.cause.Error())
}

// wrapf 使用格式化消息包装错误（内部函数）
//
// 参数：
//
//	err: 原始错误
//	format: 格式化字符串
//	args: 格式化参数
//
// 返回值：
//
//	error: 包装后的错误
func wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	err = &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}

	return err
}

// Wrapf 使用格式化消息包装错误并添加堆栈信息
//
// 参数：
//
//	err: 原始错误
//	format: 格式化字符串
//	args: 格式化参数
//
// 返回值：
//
//	error: 带堆栈信息的包装错误
func Wrapf(err error, format string, args ...any) error {
	return withStackTraceIfNotExists(wrapf(err, format, args...))
}
