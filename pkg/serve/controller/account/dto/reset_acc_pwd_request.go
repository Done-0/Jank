package dto

// ResetPwdRequest  重置密码请求体
// @Description	用户重置密码所需参数
// @Param			email					body	string	true	"用户邮箱"
// @Param			new_password			body	string	true	"新密码"
// @Param			again_new_password		body	string	true	"再次输入新密码"
// @Param			email_verification_code	body	string	true	"邮箱验证码"
type ResetPwdRequest struct {
	Email                 string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
	NewPassword           string `json:"new_password" xml:"new_password" form:"new_password" query:"new_password" validate:"required,min=6,max=20"`
	AgainNewPassword      string `json:"again_new_password" xml:"again_new_password" form:"again_new_password" query:"again_new_password" validate:"required,min=6,max=20"`
	EmailVerificationCode string `json:"email_verification_code" xml:"email_verification_code" form:"email_verification_code" query:"email_verification_code" validate:"required"`
}
