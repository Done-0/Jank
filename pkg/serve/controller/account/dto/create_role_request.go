package dto

// CreateRoleRequest 创建角色请求结构
// @Description 创建角色时的请求结构
// @Param   Code        string "角色编码"
// @Param   Description string "角色描述"
type CreateRoleRequest struct {
	Code        string `json:"code" xml:"code" form:"code" query:"code" validate:"required"`
	Description string `json:"description" xml:"description" form:"description" query:"description" validate:"required,default:''"`
}
