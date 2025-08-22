package internal

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// StackTracer 堆栈跟踪接口
type StackTracer interface {
	StackTrace() string // 获取堆栈跟踪信息
}

// withStack 带堆栈信息的错误包装器
type withStack struct {
	cause error  // 原因错误
	stack string // 堆栈信息
}

// Unwrap 返回被包装的原始错误
//
// 返回值：
//
//	error: 原始错误
func (w *withStack) Unwrap() error {
	return w.cause
}

// StackTrace 获取堆栈跟踪信息
//
// 返回值：
//
//	string: 堆栈跟踪字符串
func (w *withStack) StackTrace() string {
	return w.stack
}

// Error 错误字符串表示
//
// 返回值：
//
//	string: 包含堆栈信息的错误字符串
func (w *withStack) Error() string {
	return fmt.Sprintf("%s\nstack=%s", w.cause.Error(), w.stack)
}

// stack 获取当前调用栈信息
//
// 返回值：
//
//	string: 格式化的堆栈跟踪字符串
func stack() string {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(2, pcs[:])

	b := strings.Builder{}
	for i := 0; i < n; i++ {
		fn := runtime.FuncForPC(pcs[i])

		file, line := fn.FileLine(pcs[i])
		name := trimPathPrefix(fn.Name())
		b.WriteString(fmt.Sprintf("%s:%d %s\n", file, line, name))
	}

	return b.String()
}

// trimPathPrefix 去除路径前缀，只保留函数名
//
// 参数：
//
//	s: 完整函数名
//
// 返回值：
//
//	string: 简化后的函数名
func trimPathPrefix(s string) string {
	i := strings.LastIndex(s, "/")
	s = s[i+1:]
	i = strings.Index(s, ".")
	return s[i+1:]
}

// withStackTraceIfNotExists 如果错误没有堆栈信息则添加
//
// 参数：
//
//	err: 原始错误
//
// 返回值：
//
//	error: 带堆栈信息的错误
func withStackTraceIfNotExists(err error) error {
	if err == nil {
		return nil
	}

	// skip if stack has already exist
	var stackTracer StackTracer
	if errors.As(err, &stackTracer) {
		return err
	}

	return &withStack{
		err,
		stack(),
	}
}
