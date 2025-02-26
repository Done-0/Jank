package account

// RoleVo 角色返回值对象
// @Description 角色信息的返回结构
// @Property   ID           int64  "角色ID"
// @Property   Code         string "角色编码"
// @Property   Description  string "角色描述"
type RoleVo struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
