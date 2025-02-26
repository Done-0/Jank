package dto

// DeleteRoleRequest 删除角色请求结构
// @Description 删除角色时的请求结构
// @Param   RoleID int64 "角色ID"
type DeleteRoleRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required,gt=0"`
}
