package biz_err

const (
	SUCCESS     = 200
	UNKNOWN_ERR = 00000
	SERVER_ERR  = 10000
	BAD_REQUEST = 20000

	SEND_IMG_VERIFICATION_CODE_FAIL   = 10001
	SEND_EMAIL_VERIFICATION_CODE_FAIL = 10002
)

var CodeMsg = map[int]string{
	SUCCESS:     "请求成功",
	UNKNOWN_ERR: "未知业务异常",
	SERVER_ERR:  "服务端异常",
	BAD_REQUEST: "错误请求",

	SEND_IMG_VERIFICATION_CODE_FAIL:   "图形验证码发送失败",
	SEND_EMAIL_VERIFICATION_CODE_FAIL: "邮箱验证码发送失败",
}

func GetMessage(code int) string {
	if msg, ok := CodeMsg[code]; ok {
		return msg
	}
	return CodeMsg[UNKNOWN_ERR]
}
