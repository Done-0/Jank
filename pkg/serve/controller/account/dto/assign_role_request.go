package dto

// AssignRoleRequest 分配角色请求结构
// @Description 用户分配角色的请求结构
// @Param   RoleID       int64  "角色ID"
// @Param   PermissionID int64  "权限ID"
type AssignRoleRequest struct {
	AccountID int64 `json:"account_id" xml:"account_id" form:"account_id,gt=0"`
	RoleID    int64 `json:"role_id" xml:"role_id" form:"role_id" query:"role_id" validate:"required,gt=0"`
}
