package dto

// GetRolesByUserRequest 获取角色请求
// @Description 获取角色请求所需参数
// @Param user_id path int64 true "用户ID"
type GetRolesByUserRequest struct {
	UserID int64 `json:"user_id" xml:"user_id" form:"user_id" query:"user_id" validate:"required,gt=0"`
}
