// Package internal 错误状态处理内部实现
// 创建者：Done-0
// 创建时间：2025-08-05
package internal

import (
	"errors"
	"fmt"
	"strings"
)

// StatusError 状态错误接口
type StatusError interface {
	error
	Code() int32 // 获取状态码
}

// statusError 状态错误实现
type statusError struct {
	statusCode int32     // 状态码
	message    string    // 错误消息
	ext        Extension // 扩展信息
}

// withStatus 带状态的错误包装器
type withStatus struct {
	status *statusError // 状态错误
	stack  string       // 堆栈信息
	cause  error        // 原因错误
}

// Extension 扩展信息
type Extension struct {
	Extra map[string]string // 额外信息
}

// Code 获取状态码
// 返回值：
//
//	int32: 状态码
func (w *statusError) Code() int32 {
	return w.statusCode
}

// Msg 获取错误消息
// 返回值：
//
//	string: 错误消息
func (w *statusError) Msg() string {
	return w.message
}

// Error 错误字符串表示
// 返回值：
//
//	string: 错误字符串
func (w *statusError) Error() string {
	return fmt.Sprintf("code=%d message=%s", w.statusCode, w.message)
}

// Extra 获取额外信息
// 返回值：
//
//	map[string]string: 额外信息映射
func (w *statusError) Extra() map[string]string {
	return w.ext.Extra
}

// Code 获取状态码
func (w *withStatus) Code() int32 {
	return w.status.Code()
}

// Msg 获取错误消息
func (w *withStatus) Msg() string {
	return w.status.Msg()
}

// Extra 获取额外信息
func (w *withStatus) Extra() map[string]string {
	return w.status.Extra()
}

// Unwrap 支持 Go errors.Unwrap()
// 返回值：
//
//	error: 被包装的原始错误
func (w *withStatus) Unwrap() error {
	return w.cause
}

// Is 支持 Go errors.Is()
// 参数：
//
//	target: 目标错误
//
// 返回值：
//
//	bool: 是否匹配
func (w *withStatus) Is(target error) bool {
	var ws StatusError
	if errors.As(target, &ws) && w.status.Code() == ws.Code() {
		return true
	}
	return false
}

// As 支持 Go errors.As()
// 参数：
//
//	target: 目标接口
//
// 返回值：
//
//	bool: 是否匹配
func (w *withStatus) As(target any) bool {
	return errors.As(w.status, target)
}

// StackTrace 获取堆栈跟踪信息
// 返回值：
//
//	string: 堆栈跟踪字符串
func (w *withStatus) StackTrace() string {
	return w.stack
}

// Error 错误字符串表示
// 返回值：
//
//	string: 包含原因和堆栈的完整错误信息
func (w *withStatus) Error() string {
	b := strings.Builder{}
	b.WriteString(w.status.Error())

	if w.cause != nil {
		b.WriteString("\n")
		b.WriteString(fmt.Sprintf("cause=%s", w.cause))
	}

	if w.stack != "" {
		b.WriteString("\n")
		b.WriteString(fmt.Sprintf("stack=%s", w.stack))
	}

	return b.String()
}

// Option withStatus 配置选项函数
type Option func(ws *withStatus)

// Param 创建参数替换选项
// 参数：
//
//	k: 参数键
//	v: 参数值
//
// 返回值：
//
//	Option: 配置选项
func Param(k, v string) Option {
	return func(ws *withStatus) {
		if ws == nil || ws.status == nil {
			return
		}
		ws.status.message = strings.Replace(ws.status.message, fmt.Sprintf("{%s}", k), v, -1)
	}
}

// Extra 创建额外信息选项
// 参数：
//
//	k: 键名
//	v: 值
//
// 返回值：
//
//	Option: 配置选项
func Extra(k, v string) Option {
	return func(ws *withStatus) {
		if ws == nil || ws.status == nil {
			return
		}
		if ws.status.ext.Extra == nil {
			ws.status.ext.Extra = make(map[string]string)
		}
		ws.status.ext.Extra[k] = v
	}
}

// NewByCode 根据错误码创建新错误
// 参数：
//
//	code: 错误码
//	options: 配置选项
//
// 返回值：
//
//	error: 创建的错误
func NewByCode(code int32, options ...Option) error {
	ws := &withStatus{
		status: getStatusByCode(code),
		cause:  nil,
		stack:  stack(),
	}

	for _, opt := range options {
		opt(ws)
	}

	return ws
}

// WrapByCode 使用错误码包装现有错误
// 参数：
//
//	err: 原始错误
//	code: 错误码
//	options: 配置选项
//
// 返回值：
//
//	error: 包装后的错误
func WrapByCode(err error, code int32, options ...Option) error {
	if err == nil {
		return nil
	}

	ws := &withStatus{
		status: getStatusByCode(code),
		cause:  err,
	}

	for _, opt := range options {
		opt(ws)
	}

	// skip if stack has already exist
	var stackTracer StackTracer
	if errors.As(err, &stackTracer) {
		return ws
	}

	ws.stack = stack()

	return ws
}

// getStatusByCode 根据错误码获取状态错误
// 参数：
//
//	code: 错误码
//
// 返回值：
//
//	*statusError: 状态错误实例
func getStatusByCode(code int32) *statusError {
	codeDefinition, ok := CodeDefinitions[code]
	if ok {
		// predefined err code
		return &statusError{
			statusCode: code,
			message:    codeDefinition.Message,
			ext:        Extension{},
		}
	}

	return &statusError{
		statusCode: code,
		message:    DefaultErrorMsg,
		ext:        Extension{},
	}
}
