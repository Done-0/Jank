package dto

// CheckPermissionRequest 检查用户权限请求结构
// @Description 检查用户是否拥有某项权限
// @Param   AccountID      int64  "用户ID"
// @Param   PermissionCode string "权限编码"
type CheckPermissionRequest struct {
	AccountID      int64  `json:"account_id" xml:"account_id" form:"account_id" query:"account_id" validate:"required"`
	PermissionCode string `json:"permission_code" xml:"permission_code" form:"permission_code" query:"permission_code" validate:"required"`
}
