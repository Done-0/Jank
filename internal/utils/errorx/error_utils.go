package errorx

import (
	"fmt"
	"strings"

	"github.com/Done-0/jank/internal/utils/errorx/internal"
)

// StatusError 带状态码的错误接口
type StatusError interface {
	error
	Code() int32              // 获取错误码
	Msg() string              // 获取错误消息
	Extra() map[string]string // 获取额外信息
}

// Option StatusError 配置选项
type Option = internal.Option

// KV 创建键值对参数选项
// 参数：
//
//	k: 键名
//	v: 值
//
// 返回值：
//
//	Option: 配置选项
func KV(k, v string) Option {
	return internal.Param(k, v)
}

// KVf 创建格式化键值对参数选项
// 参数：
//
//	k: 键名
//	v: 格式化字符串
//	a: 格式化参数
//
// 返回值：
//
//	Option: 配置选项
func KVf(k, v string, a ...any) Option {
	formatValue := fmt.Sprintf(v, a...)
	return internal.Param(k, formatValue)
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
	return internal.Extra(k, v)
}

// New 根据错误码创建新的错误
// 参数：
//
//	code: 状态码
//	options: 配置选项
//
// 返回值：
//
//	error: 创建的错误
func New(code int32, options ...Option) error {
	return internal.NewByCode(code, options...)
}

// WrapByCode 使用状态码包装现有错误
// 参数：
//
//	err: 原始错误
//	statusCode: 状态码
//	options: 配置选项
//
// 返回值：
//
//	error: 包装后的错误
func WrapByCode(err error, statusCode int32, options ...Option) error {
	if err == nil {
		return nil
	}

	return internal.WrapByCode(err, statusCode, options...)
}

// Wrapf 使用格式化消息包装错误
// 参数：
//
//	err: 原始错误
//	format: 格式化字符串
//	args: 格式化参数
//
// 返回值：
//
//	error: 包装后的错误
func Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	return internal.Wrapf(err, format, args...)
}

// ErrorWithoutStack 获取不带堆栈信息的错误消息
// 参数：
//
//	err: 错误对象
//
// 返回值：
//
//	string: 不带堆栈的错误消息
func ErrorWithoutStack(err error) string {
	if err == nil {
		return ""
	}
	errMsg := err.Error()
	index := strings.Index(errMsg, "stack=")
	if index != -1 {
		errMsg = errMsg[:index]
	}
	return errMsg
}

// Register 注册错误码定义
// 参数：
//
//	code: 错误码
//	msg: 错误消息模板
//	opts: 注册选项
func Register(code int32, msg string, opts ...internal.RegisterOption) {
	internal.Register(code, msg, opts...)
}

// SetDefaultErrorCode 设置默认错误码
// 参数：
//
//	code: 默认错误码
func SetDefaultErrorCode(code int32) {
	internal.SetDefaultErrorCode(code)
}

// RegisterOption 注册选项类型
type RegisterOption = internal.RegisterOption
