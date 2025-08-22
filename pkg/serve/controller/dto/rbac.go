// Package dto 提供 RBAC 权限管理相关的数据传输对象
package dto

// CreatePermissionRequest 创建权限策略请求
type CreatePermissionRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=100"`                 // 权限名称
	Description string `json:"description" validate:"omitempty,max=500"`               // 权限描述（可选）
	Role        string `json:"role" validate:"required,min=1,max=100"`                 // 角色名称
	Resource    string `json:"resource" validate:"required,min=1,max=255"`             // 资源路径
	Action      string `json:"action" validate:"required,oneof=* GET POST PUT DELETE"` // 操作方法
}

// DeletePermissionRequest 删除权限策略请求
type DeletePermissionRequest struct {
	Role     string `json:"role" validate:"required,min=1,max=100"`                 // 角色名称
	Resource string `json:"resource" validate:"required,min=1,max=255"`             // 资源路径
	Action   string `json:"action" validate:"required,oneof=* GET POST PUT DELETE"` // 操作方法
}

// AssignPermissionRequest 为角色分配权限请求
type AssignPermissionRequest struct {
	Role     string `json:"role" validate:"required,min=1,max=100"`                 // 角色名称
	Resource string `json:"resource" validate:"required,min=1,max=255"`             // 资源路径
	Action   string `json:"action" validate:"required,oneof=* GET POST PUT DELETE"` // 操作方法
}

// RevokePermissionRequest 撤销角色权限请求
type RevokePermissionRequest struct {
	Role     string `json:"role" validate:"required,min=1,max=100"`                 // 角色名称
	Resource string `json:"resource" validate:"required,min=1,max=255"`             // 资源路径
	Action   string `json:"action" validate:"required,oneof=* GET POST PUT DELETE"` // 操作方法
}

// CreateRoleRequest 创建角色请求
type CreateRoleRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=100"`                 // 角色名称
	Description string `json:"description" validate:"omitempty,max=500"`               // 角色描述（可选）
	Role        string `json:"role" validate:"required,min=1,max=100"`                 // 角色标识符
	Resource    string `json:"resource" validate:"required,min=1,max=255"`             // 资源路径
	Action      string `json:"action" validate:"required,oneof=* GET POST PUT DELETE"` // 操作方法
}

// DeleteRoleRequest 删除角色请求
type DeleteRoleRequest struct {
	Role string `json:"role" validate:"required,min=1,max=100"` // 角色名称
}

// GetRolePermissionsRequest 获取角色权限请求
type GetRolePermissionsRequest struct {
	Role string `query:"role" validate:"required,min=1,max=100"` // 角色名称
}

// AssignRoleRequest 为用户分配角色请求
type AssignRoleRequest struct {
	UserID string `json:"user_id" validate:"required,min=1,max=100"` // 用户 ID
	Role   string `json:"role" validate:"required,min=1,max=100"`    // 角色名称
}

// RevokeRoleRequest 撤销用户角色请求
type RevokeRoleRequest struct {
	UserID string `json:"user_id" validate:"required,min=1,max=100"` // 用户 ID
	Role   string `json:"role" validate:"required,min=1,max=100"`    // 角色名称
}

// GetUserRolesRequest 获取用户角色请求
type GetUserRolesRequest struct {
	UserID string `query:"user_id" validate:"required,min=1,max=100"` // 用户 ID
}

// CheckPermissionRequest 权限检查请求
type CheckPermissionRequest struct {
	UserID   string `json:"user_id" validate:"required,min=1,max=100"`              // 用户 ID
	Resource string `json:"resource" validate:"required,min=1,max=255"`             // 资源路径
	Action   string `json:"action" validate:"required,oneof=* GET POST PUT DELETE"` // 操作方法
}
