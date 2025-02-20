package dto

// UpdateRoleRequest 更新角色请求结构
// @Description 更新角色时的请求结构
// @Param   ID          int64  "角色ID"
// @Param   Code        string "角色编码"
// @Param   Description string "角色描述"
type UpdateRoleRequest struct {
	ID          int64  `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
	Code        string `json:"code" xml:"code" form:"code" query:"code" validate:"required"`
	Description string `json:"description" xml:"description" form:"description" query:"description" validate:"required"`
}
