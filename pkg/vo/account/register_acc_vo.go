package account

// RegisterAccountVO     获取账户信息请求体
// @Description	请求获取账户信息时所需参数
// @Property			email	    body	string	true	"用户邮箱"
// @Property			nickname	body	string	true	"用户昵称"
// @Property			role_code	body	string	true	"用户角色编码"
type RegisterAccountVO struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}
