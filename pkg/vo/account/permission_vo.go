package account

// PermissionVo 权限返回值对象
// @Description 权限信息的返回结构
// @Property   ID           int64  "权限ID"
// @Property   Code         string "权限编码"
// @Property   Description  string "权限描述"
type PermissionVo struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
