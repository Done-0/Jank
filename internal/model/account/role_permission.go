package model

import "jank.com/jank_blog/internal/model/base"

// RolePermission 角色和权限关联表
type RolePermission struct {
	base.Base
	RoleID       int64 `gorm:"index" json:"role_id"`       // 角色ID
	PermissionID int64 `gorm:"index" json:"permission_id"` // 权限ID
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
