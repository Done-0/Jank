// Package code 错误码注册适配器
// 创建者：Done-0
// 创建时间：2025-08-05
package code

import (
	"github.com/Done-0/jank/internal/utils/errorx"
)

// RegisterOptionFn 注册选项函数类型
type RegisterOptionFn = errorx.RegisterOption

// Register 注册预定义的错误码信息
// 参数：
//
//	code: 错误码
//	msg: 错误消息模板
//	opts: 注册选项
func Register(code int32, msg string, opts ...RegisterOptionFn) {
	errorx.Register(code, msg, opts...)
}

// SetDefaultErrorCode 设置默认错误码
// 参数：
//
//	code: 默认错误码
func SetDefaultErrorCode(code int32) {
	errorx.SetDefaultErrorCode(code)
}
