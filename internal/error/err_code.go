package biz_err

const (
	Success     = 200
	UnKnowErr   = 00000
	ServerError = 10000
	BadRequest  = 20000

	SendImgVerificationCodeFail   = 10001
	SendEmailVerificationCodeFail = 10002
)

var CodeMsg = map[int]string{
	Success:     "请求成功",
	UnKnowErr:   "未知业务异常",
	ServerError: "服务端异常",
	BadRequest:  "错误请求",

	SendImgVerificationCodeFail:   "图形验证码发送失败",
	SendEmailVerificationCodeFail: "邮箱验证码发送失败",
}

func GetMessage(code int) string {
	if msg, ok := CodeMsg[code]; ok {
		return msg
	}
	return CodeMsg[UnKnowErr]
}
