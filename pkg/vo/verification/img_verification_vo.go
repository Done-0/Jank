// Package verification 提供验证码相关的数据传输对象
// 创建者：Done-0
// 创建时间：2025-05-10
package verification

// ImgVerificationVO        图片验证码
// @Description             图片验证码
// @Property		img	body	string	true	"图片的base64编码"
type ImgVerificationVO struct {
	ImgBase64 string `json:"imgBase64"`
}
