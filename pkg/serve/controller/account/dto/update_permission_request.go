package dto

// UpdatePermissionRequest 更新权限请求结构
// @Description 更新权限时的请求结构
// @Param   ID          int64  "权限ID"
// @Param   Code        string "权限编码"
// @Param   Description string "权限描述"
type UpdatePermissionRequest struct {
	ID          int64  `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
	Code        string `json:"code" xml:"code" form:"code" query:"code" validate:"required"`
	Description string `json:"description" xml:"description" form:"description" query:"description" validate:"required"`
}
