package account

// LoginVO           返回给前端的登录信息
// @Description	登录成功后返回的访问令牌和刷新令牌
// @Property			access_token	body	string	true	"访问令牌"
// @Property			refresh_token	body	string	true	"刷新令牌"
type LoginVO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
