// Package biz_err 提供业务错误码和错误信息定义
// 创建者：Done-0
// 创建时间：2025-05-10
package biz_err

// 错误码常量定义
const (
	SUCCESS     = 200
	UNKNOWN_ERR = 00000
	SERVER_ERR  = 10000
	BAD_REQUEST = 20000

	SEND_IMG_VERIFICATION_CODE_FAIL   = 10001
	SEND_EMAIL_VERIFICATION_CODE_FAIL = 10002
)

// CodeMsg 错误码对应的错误信息
var CodeMsg = map[int]string{
	SUCCESS:     "请求成功",
	UNKNOWN_ERR: "未知业务异常",
	SERVER_ERR:  "服务端异常",
	BAD_REQUEST: "错误请求",

	SEND_IMG_VERIFICATION_CODE_FAIL:   "图形验证码发送失败",
	SEND_EMAIL_VERIFICATION_CODE_FAIL: "邮箱验证码发送失败",
}

// GetMessage 根据错误码获取对应的错误信息
// 参数：
//   - code: 错误码
//
// 返回值：
//   - string: 错误信息
func GetMessage(code int) string {
	if msg, ok := CodeMsg[code]; ok {
		return msg
	}
	return CodeMsg[UNKNOWN_ERR]
}
