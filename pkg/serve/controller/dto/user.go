// Package dto 提供用户相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-08-05
package dto

// RegisterRequest 用户注册请求
type RegisterRequest struct {
	Email                 string `json:"email" validate:"required,email"`                   // 邮箱
	Password              string `json:"password" validate:"required,min=6,max=20"`         // 密码
	Nickname              string `json:"nickname" validate:"required,min=2,max=20"`         // 昵称
	EmailVerificationCode string `json:"email_verification_code" validate:"required,len=6"` // 邮箱验证码
}

// LoginRequest 用户登录请求
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`           // 邮箱
	Password string `json:"password" validate:"required,min=6,max=20"` // 密码
}

// RefreshTokenRequest 刷新token请求
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"` // refresh token
}

// UpdateRequest 更新用户信息请求
type UpdateRequest struct {
	Nickname string `json:"nickname" validate:"omitempty,min=2,max=20"` // 昵称
	Avatar   string `json:"avatar" validate:"omitempty,url"`            // 头像
}

// ResetPasswordRequest 重置密码请求
type ResetPasswordRequest struct {
	OldPassword           string `json:"old_password" validate:"required,min=6,max=20"`     // 原密码
	NewPassword           string `json:"new_password" validate:"required,min=6,max=20"`     // 新密码
	EmailVerificationCode string `json:"email_verification_code" validate:"required,len=6"` // 邮箱验证码
}

// ListUsersRequest 获取用户列表请求
type ListUsersRequest struct {
	PageNo   int64  `query:"page_no" validate:"required,min=1"`           // 页码
	PageSize int64  `query:"page_size" validate:"required,min=1,max=100"` // 每页数量
	Keyword  string `query:"keyword" validate:"omitempty"`                // 搜索关键词（邮箱、昵称）
	Role     string `query:"role" validate:"omitempty"`                   // 角色筛选
}

// UpdateUserRoleRequest 管理员更新用户角色请求
type UpdateUserRoleRequest struct {
	ID   string `json:"id" validate:"required,min=1"` // 目标用户 ID
	Role string `json:"role" validate:"required"`     // 新角色
}
