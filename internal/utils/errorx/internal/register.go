// Package internal 错误码注册内部实现
// 创建者：Done-0
// 创建时间：2025-08-05
package internal

const (
	DefaultErrorMsg = "Service Internal Error" // 默认错误消息
)

var (
	ServiceInternalErrorCode int32 = 1                               // 服务内部错误码
	CodeDefinitions                = make(map[int32]*CodeDefinition) // 错误码定义映射
)

// CodeDefinition 错误码定义
type CodeDefinition struct {
	Code    int32  // 错误码
	Message string // 错误消息模板
}

// RegisterOption 注册选项函数
type RegisterOption func(definition *CodeDefinition)

// Register 注册错误码定义
// 参数：
//
//	code: 错误码
//	msg: 错误消息模板
//	opts: 注册选项
func Register(code int32, msg string, opts ...RegisterOption) {
	definition := &CodeDefinition{
		Code:    code,
		Message: msg,
	}

	for _, opt := range opts {
		opt(definition)
	}

	CodeDefinitions[code] = definition
}

// SetDefaultErrorCode 设置默认错误码
// 参数：
//
//	code: 默认错误码
func SetDefaultErrorCode(code int32) {
	ServiceInternalErrorCode = code
}
