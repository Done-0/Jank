package dto

// GetRolesByAccRequest 获取角色请求
// @Description 获取角色请求所需参数
// @Param account_id path int true "账户ID"
type GetRolesByAccRequest struct {
	AccountID int64 `json:"account_id" xml:"account_id" form:"account_id" query:"account_id" validate:"required,gt=0"`
}
