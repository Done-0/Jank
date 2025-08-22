// Package vo 用户相关值对象
// 创建者：Done-0
// 创建时间：2025-08-05
package vo

// RegisterResponse 用户注册响应
type RegisterResponse struct {
	ID       string   `json:"id"`       // 用户 ID
	Email    string   `json:"email"`    // 用户邮箱
	Nickname string   `json:"nickname"` // 用户昵称
	Roles    []string `json:"roles"`    // 用户角色列表
}

// LoginResponse 用户登录响应
type LoginResponse struct {
	AccessToken  string `json:"access_token"`  // 访问令牌
	RefreshToken string `json:"refresh_token"` // 刷新令牌
}

// LogoutResponse 用户登出响应
type LogoutResponse struct {
	Message string `json:"message"` // 登出结果消息
}

// RefreshTokenResponse 刷新token响应
type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`  // 新的访问令牌
	RefreshToken string `json:"refresh_token"` // 新的刷新令牌
}

// GetProfileResponse 获取用户资料响应
type GetProfileResponse struct {
	ID       string   `json:"id"`       // 用户 ID
	Email    string   `json:"email"`    // 用户邮箱
	Nickname string   `json:"nickname"` // 用户昵称
	Avatar   string   `json:"avatar"`   // 用户头像
	Roles    []string `json:"roles"`    // 用户角色列表
}

// UpdateResponse 更新用户信息响应
type UpdateResponse struct {
	ID       string   `json:"id"`       // 用户 ID
	Email    string   `json:"email"`    // 用户邮箱
	Nickname string   `json:"nickname"` // 用户昵称
	Avatar   string   `json:"avatar"`   // 用户头像
	Roles    []string `json:"roles"`    // 用户角色列表
}

// ResetPasswordResponse 重置密码响应
type ResetPasswordResponse struct {
	Message string `json:"message"` // 重置结果消息
}

// UpdateUserRoleResponse 管理员更新用户角色响应
type UpdateUserRoleResponse struct {
	ID       string   `json:"id"`       // 用户 ID
	Email    string   `json:"email"`    // 用户邮箱
	Nickname string   `json:"nickname"` // 用户昵称
	Roles    []string `json:"roles"`    // 更新后的角色列表
	Message  string   `json:"message"`  // 操作结果消息
}

// UserItem 用户列表项
type UserItem struct {
	ID       string   `json:"id"`       // 用户 ID
	Email    string   `json:"email"`    // 用户邮箱
	Nickname string   `json:"nickname"` // 用户昵称
	Avatar   string   `json:"avatar"`   // 用户头像
	Roles    []string `json:"roles"`    // 用户角色列表
}

// ListUsersResponse 用户列表响应
type ListUsersResponse struct {
	Total    int64       `json:"total"`     // 总数量
	PageNo   int64       `json:"page_no"`   // 当前页码
	PageSize int64       `json:"page_size"` // 每页数量
	List     []*UserItem `json:"list"`      // 用户列表
}
