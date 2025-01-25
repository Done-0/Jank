package account

// GetAccountVo     获取账户信息请求体
//
//	@Description	请求获取账户信息时所需参数
//	@Param			email	    body	string	true	"用户邮箱"
//	@Param			nickname	body	string	true	"用户昵称"
//	@Param			phone	    body	string	true	"用户手机号"
//	@Param			role_code	body	string	true	"用户角色编码"
type GetAccountVo struct {
	Email    string `json:"email" xml:"email" form:"email" query:"email"`
	Nickname string `json:"nickname" xml:"nickname" form:"nickname" query:"nickname"`
	Phone    string `json:"phone" xml:"phone" form:"phone" query:"phone"`
	RoleCode string `json:"role_code" xml:"role_code" form:"role_code" query:"role_code"`
}
