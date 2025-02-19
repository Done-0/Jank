package dto

// AssignRoleRequest 分配角色请求结构
// @Description 用户分配角色的请求结构
// @Param   RoleID       int64  "角色ID"
// @Param   PermissionID int64  "权限ID"
type AssignRoleRequest struct {
	RoleID       int64 `json:"role_id" xml:"role_id" form:"role_id" query:"role_id" validate:"required"`
	PermissionID int64 `json:"permission_ids" xml:"permission_ids" form:"permission_ids" query:"permission_ids" validate:"required"`
}
