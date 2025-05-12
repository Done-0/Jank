// Package dto 提供账户相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// RegisterRequest  用户注册请求体
// @Description	用户注册所需参数
// @Param			email		body	string	true	"用户邮箱"
// @Param			phone		body	string	true	"用户手机号"
// @Param			nickname	body	string	true	"用户昵称"
// @Param			password	body	string	true	"用户密码"
// @Param			email_verification_code	body	string	true	"用户邮箱验证码"
// @Param			img_verification_code	body	string	true	"用户图片验证码"
type RegisterRequest struct {
	Email                 string `json:"email" xml:"email" form:"email" query:"email" validate:"required"`
	Phone                 string `json:"phone" xml:"phone" form:"phone" query:"phone" default:""`
	Nickname              string `json:"nickname" xml:"nickname" form:"nickname" query:"nickname" validate:"required,min=1,max=20"`
	Password              string `json:"password" xml:"password" form:"password" query:"password" validate:"required,min=6,max=20"`
	EmailVerificationCode string `json:"email_verification_code" xml:"email_verification_code" form:"email_verification_code" query:"email_verification_code" validate:"required"`
	ImgVerificationCode   string `json:"img_verification_code" xml:"img_verification_code" form:"img_verification_code" query:"img_verification_code" validate:"required"`
}
