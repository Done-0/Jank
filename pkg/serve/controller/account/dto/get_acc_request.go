package dto

// GetAccountRequest            获取账户信息请求体
// @Description	请求获取账户信息时所需参数
// @Param			email	    body	string	true	"用户邮箱"
type GetAccountRequest struct {
	Email string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
}
