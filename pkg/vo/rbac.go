// Package vo RBAC 相关值对象
// 创建者：Done-0
// 创建时间：2025-08-08
package vo

// PolicyResponse 权限策略响应 (p, role, resource, action)
type PolicyResponse struct {
	Role     string `json:"role"`     // 角色名称
	Resource string `json:"resource"` // 资源路径
	Action   string `json:"action"`   // 操作方法
}

// PermissionResponse 权限响应
type PermissionResponse struct {
	Name        string `json:"name"`        // 权限名称
	Description string `json:"description"` // 权限描述
	Resource    string `json:"resource"`    // 资源路径
	Action      string `json:"action"`      // 操作方法
}

// PermissionListResponse 权限列表响应
type PermissionListResponse struct {
	Total int64                `json:"total"` // 总条数
	List  []PermissionResponse `json:"list"`  // 权限列表
}

// PolicyOpResponse 策略操作响应
type PolicyOpResponse struct {
	Success     bool   `json:"success"`     // 操作是否成功
	Name        string `json:"name"`        // 权限名称
	Description string `json:"description"` // 权限描述
	Role        string `json:"role"`        // 角色名称
	Resource    string `json:"resource"`    // 资源路径
	Action      string `json:"action"`      // 操作方法
	Message     string `json:"message"`     // 操作消息
}

// PermissionOpResponse 权限分配操作响应
type PermissionOpResponse struct {
	Success    bool   `json:"success"`    // 操作是否成功
	Role       string `json:"role"`       // 角色名称
	Permission string `json:"permission"` // 权限名称
	Message    string `json:"message"`    // 操作结果消息
}

// RoleListResponse 角色列表响应
type RoleListResponse struct {
	Total int64    `json:"total"` // 总条数
	List  []string `json:"list"`  // 角色列表
}

// RolePermissionsResponse 角色权限响应
type RolePermissionsResponse struct {
	Role        string               `json:"role"`        // 角色名称
	Total       int64                `json:"total"`       // 权限总数
	Permissions []PermissionResponse `json:"permissions"` // 权限列表
}

// UserRolesResponse 用户角色响应
type UserRolesResponse struct {
	UserID string   `json:"user_id"` // 用户 ID
	Roles  []string `json:"roles"`   // 角色列表
}

// ListUserResponse 用户列表响应
type ListUserResponse struct {
	Total int64      `json:"total"` // 总条数
	List  []UserInfo `json:"list"`  // 用户列表
}

// UserInfo 用户信息
type UserInfo struct {
	ID       string   `json:"id"`       // 用户 ID
	Nickname string   `json:"nickname"` // 用户昵称
	Email    string   `json:"email"`    // 用户邮箱
	Roles    []string `json:"roles"`    // 用户角色列表
}

// RoleOpResponse 用户角色操作响应
type RoleOpResponse struct {
	Success bool     `json:"success"` // 操作是否成功
	UserID  string   `json:"user_id"` // 用户 ID
	Role    string   `json:"role"`    // 角色名称
	Roles   []string `json:"roles"`   // 用户所有角色列表（成功时返回）
	Message string   `json:"message"` // 操作结果消息
}

// CheckResponse 权限检查响应
type CheckResponse struct {
	Allowed bool   `json:"allowed"` // 是否允许
	Reason  string `json:"reason"`  // 原因说明
}

// CreateRoleResponse 创建角色响应
type CreateRoleResponse struct {
	Success     bool   `json:"success"`     // 操作是否成功
	Name        string `json:"name"`        // 角色名称
	Description string `json:"description"` // 角色描述
	Role        string `json:"role"`        // 角色标识符
	Resource    string `json:"resource"`    // 资源路径
	Action      string `json:"action"`      // 操作方法
	Message     string `json:"message"`     // 操作消息
}

// DeleteRoleResponse 删除角色响应
type DeleteRoleResponse struct {
	Success bool   `json:"success"` // 删除是否成功
	Role    string `json:"role"`    // 角色名称
	Message string `json:"message"` // 操作结果消息
}
