package dto

// GetPermissionsByRoleRequest 获取角色权限请求
// @Description 获取角色权限请求所需参数
// @Param role_id path int true "角色ID"
type GetPermissionsByRoleRequest struct {
	RoleID int64 `json:"role_id" xml:"role_id" form:"role_id" query:"role_id" validate:"required,gt=0"`
}
