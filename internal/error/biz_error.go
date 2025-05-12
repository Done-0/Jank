// Package biz_err 提供业务错误码和错误信息定义
// 创建者：Done-0
// 创建时间：2025-05-10
package biz_err

// Err 业务错误结构体
type Err struct {
	Code int    `json:"code"` // 错误码
	Msg  string `json:"msg"`  // 错误信息
}

// Error 实现 error 接口的方法
// 返回值：
//   - string: 错误信息
func (b *Err) Error() string {
	return b.Msg
}

// New 创建一个 Err 实例，基于提供的错误代码和可选的错误信息
// 参数：
//   - code: 错误码
//   - msg: 可选的错误信息，不提供则使用错误码对应的默认信息
//
// 返回值：
//   - *Err: 业务错误实例
func New(code int, msg ...string) *Err {
	message := ""

	if len(msg) <= 0 {
		message = GetMessage(code)
	} else {
		message = msg[0]
	}

	return &Err{
		Code: code,
		Msg:  message,
	}
}
