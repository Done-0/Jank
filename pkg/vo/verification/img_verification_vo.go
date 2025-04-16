package verification

// ImgVerificationVO        图片验证码
// @Description             图片验证码
// @Property		img	body	string	true	"图片的base64编码"
type ImgVerificationVO struct {
	ImgBase64 string `json:"imgBase64"`
}
