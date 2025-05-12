// Package dto 提供账户相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-05-10
package dto

// LoginRequest 用户登录请求体
// @Description	用户登录请求所需参数
// @Param			email		body	string	true	"用户邮箱"
// @Param			password	body	string	true	"用户密码"
// @Param			img_verification_code	body	string	true	"图片验证码"
type LoginRequest struct {
	Email               string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
	Password            string `json:"password" xml:"password" form:"password" query:"password" validate:"required"`
	ImgVerificationCode string `json:"img_verification_code" xml:"img_verification_code" form:"img_verification_code" query:"img_verification_code" validate:"required"`
}
