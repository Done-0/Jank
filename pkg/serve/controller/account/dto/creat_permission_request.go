package dto

// CreatePermissionRequest 创建权限请求结构
// @Description 创建权限时的请求结构
// @Param   Code        string "权限编码"
// @Param   Description string "权限描述"
type CreatePermissionRequest struct {
	Code        string `json:"code" xml:"code" form:"code" query:"code" validate:"required"`
	Description string `json:"description" xml:"description" form:"description" query:"description" validate:"required,default:''"`
}
