package biz_err

type Err struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (b *Err) Error() string {
	return b.Msg
}

// New 创建一个 Err 实例，基于提供的错误代码和可选的错误信息。
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
