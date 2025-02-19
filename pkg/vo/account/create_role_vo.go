package account

// CreateRoleVo 表示创建角色后的响应结构
// @Description 角色创建成功后的返回数据
// @Property   ID           int64  "角色ID"
// @Property   Code         string "角色编码"
// @Property   Description  string "角色描述"
type CreateRoleVo struct {
	ID          int64  `json:"id" xml:"id" form:"id" query:"id"`
	Code        string `json:"code" xml:"code" form:"code" query:"code"`
	Description string `json:"description" xml:"description" form:"description" query:"description"`
}
