package account

// ImgVerificationVO        图片验证码
// @Description             图片验证码
// @Param		img	body	string	true	"图片的base64编码"
type ImgVerificationVO struct {
	Img string `json:"img"`
}
