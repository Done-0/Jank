package dto

// AssignPermissionRequest 分配权限请求结构
// @Description 分配权限给角色的请求结构
// @Param   RoleID       int64 "角色ID"
// @Param   PermissionID int64 "权限ID"
type AssignPermissionRequest struct {
	RoleID       int64 `json:"role_id" xml:"role_id" form:"role_id" query:"role_id" validate:"required"`
	PermissionID int64 `json:"permission_id" xml:"permission_id" form:"permission_id" query:"permission_id" validate:"required"`
}
