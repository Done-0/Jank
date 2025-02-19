package dto

// DeletePermissionRequest 删除权限请求结构
// @Description 删除权限时的请求结构
// @Param   ID int64 "权限ID"
type DeletePermissionRequest struct {
	ID int64 `json:"id" xml:"id" form:"id" query:"id" validate:"required"`
}
