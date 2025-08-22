// Package impl 提供RBAC权限相关的数据访问实现
// 创建者：Done-0
// 创建时间：2025-08-11
package impl

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/model/rbac"
	"github.com/Done-0/jank/internal/utils/db"
	"github.com/Done-0/jank/pkg/serve/mapper"
)

// RBACMapperImpl RBAC数据访问实现
type RBACMapperImpl struct{}

// NewRBACMapper 创建RBAC数据访问实例
func NewRBACMapper() mapper.RBACMapper {
	return &RBACMapperImpl{}
}

// CreatePermission 创建权限
func (m *RBACMapperImpl) CreatePermission(c *app.RequestContext, name, description, role, resource, action string) (*rbac.Policy, error) {
	var count int64
	if err := db.GetDBFromContext(c).Model(&rbac.Policy{}).Where("ptype = ? AND v0 = ? AND v1 = ? AND v2 = ? AND deleted = ?", "p", role, resource, action, false).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, nil
	}

	policy := &rbac.Policy{
		Ptype: "p",
		V0:    role,
		V1:    resource,
		V2:    action,
		V3:    name,
		V4:    description,
	}

	if err := db.GetDBFromContext(c).Create(policy).Error; err != nil {
		return nil, err
	}

	return policy, nil
}

// DeletePermission 删除权限（软删除）
func (m *RBACMapperImpl) DeletePermission(c *app.RequestContext, role, resource, action string) (bool, error) {
	result := db.GetDBFromContext(c).Model(&rbac.Policy{}).Where("ptype = ? AND v0 = ? AND v1 = ? AND v2 = ? AND deleted = ?", "p", role, resource, action, false).Update("deleted", true)
	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}

// ListPermissions 获取所有权限
func (m *RBACMapperImpl) ListPermissions(c *app.RequestContext) ([]*rbac.Policy, error) {
	var policies []*rbac.Policy
	err := db.GetDBFromContext(c).
		Where("ptype = ? AND deleted = ?", "p", false).
		Order("id DESC").
		Find(&policies).Error

	if err != nil {
		return nil, err
	}

	permissionMap := make(map[string]*rbac.Policy)
	var result []*rbac.Policy

	for _, policy := range policies {
		key := policy.V1 + ":" + policy.V2
		if _, exists := permissionMap[key]; !exists {
			permissionMap[key] = policy
			result = append(result, policy)
		}
	}

	return result, nil
}

// PermissionExists 检查权限是否存在
func (m *RBACMapperImpl) PermissionExists(c *app.RequestContext, resource, action string) (bool, error) {
	var count int64
	err := db.GetDBFromContext(c).Model(&rbac.Policy{}).Where("ptype = ? AND v1 = ? AND v2 = ? AND deleted = ?", "p", resource, action, false).Count(&count).Error
	return count > 0, err
}

// ListRoles 获取所有角色
func (m *RBACMapperImpl) ListRoles(c *app.RequestContext) ([]*rbac.Policy, error) {
	var policies []*rbac.Policy
	err := db.GetDBFromContext(c).
		Where("ptype = ? AND deleted = ?", "p", false).
		Order("id DESC").
		Find(&policies).Error

	if err != nil {
		return nil, err
	}

	roleMap := make(map[string]*rbac.Policy)
	var result []*rbac.Policy

	for _, policy := range policies {
		if _, exists := roleMap[policy.V0]; !exists {
			roleMap[policy.V0] = policy
			result = append(result, policy)
		}
	}

	return result, nil
}

// RoleExists 检查角色是否存在
func (m *RBACMapperImpl) RoleExists(c *app.RequestContext, role string) (bool, error) {
	var count int64
	err := db.GetDBFromContext(c).Model(&rbac.Policy{}).Where("ptype = ? AND v0 = ? AND deleted = ?", "p", role, false).Count(&count).Error
	return count > 0, err
}

// GetRolePermissions 获取角色权限
func (m *RBACMapperImpl) GetRolePermissions(c *app.RequestContext, role string) ([]*rbac.Policy, error) {
	var policies []*rbac.Policy
	err := db.GetDBFromContext(c).Where("ptype = ? AND v0 = ? AND deleted = ?", "p", role, false).Order("id DESC").Find(&policies).Error
	return policies, err
}

// AssignRole 分配角色
func (m *RBACMapperImpl) AssignRole(c *app.RequestContext, user, role string) (*rbac.Policy, error) {
	var count int64
	if err := db.GetDBFromContext(c).Model(&rbac.Policy{}).Where("ptype = ? AND v0 = ? AND v1 = ? AND deleted = ?", "g", user, role, false).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, nil
	}

	policy := &rbac.Policy{
		Ptype: "g",
		V0:    user,
		V1:    role,
	}

	if err := db.GetDBFromContext(c).Create(policy).Error; err != nil {
		return nil, err
	}

	return policy, nil
}

// RevokeRole 撤销角色（软删除）
func (m *RBACMapperImpl) RevokeRole(c *app.RequestContext, user, role string) (bool, error) {
	result := db.GetDBFromContext(c).Model(&rbac.Policy{}).Where("ptype = ? AND v0 = ? AND v1 = ? AND deleted = ?", "g", user, role, false).Update("deleted", true)
	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}

// GetUserRoles 获取用户角色
func (m *RBACMapperImpl) GetUserRoles(c *app.RequestContext, user string) ([]*rbac.Policy, error) {
	var policies []*rbac.Policy
	err := db.GetDBFromContext(c).Where("ptype = ? AND v0 = ? AND deleted = ?", "g", user, false).Order("id DESC").Find(&policies).Error
	return policies, err
}

// UserHasRole 检查用户是否有指定角色
func (m *RBACMapperImpl) UserHasRole(c *app.RequestContext, user, role string) (bool, error) {
	var count int64
	err := db.GetDBFromContext(c).Model(&rbac.Policy{}).Where("ptype = ? AND v0 = ? AND v1 = ? AND deleted = ?", "g", user, role, false).Count(&count).Error
	return count > 0, err
}

// ListUsers 获取所有用户
func (m *RBACMapperImpl) ListUsers(c *app.RequestContext) ([]*rbac.Policy, error) {
	var policies []*rbac.Policy
	err := db.GetDBFromContext(c).
		Where("ptype = ? AND deleted = ?", "g", false).
		Order("id DESC").
		Find(&policies).Error

	if err != nil {
		return nil, err
	}

	userMap := make(map[string]*rbac.Policy)
	var result []*rbac.Policy

	for _, policy := range policies {
		if _, exists := userMap[policy.V0]; !exists {
			userMap[policy.V0] = policy
			result = append(result, policy)
		}
	}

	return result, nil
}

// CheckPermission 权限检查
func (m *RBACMapperImpl) CheckPermission(c *app.RequestContext, user, resource, action string) (bool, error) {
	var directCount int64
	err := db.GetDBFromContext(c).Model(&rbac.Policy{}).
		Where("ptype = ? AND v0 = ? AND (v1 = ? OR v1 = ?) AND (v2 = ? OR v2 = ?) AND deleted = ?",
			"p", user, resource, "*", action, "*", false).
		Count(&directCount).Error
	if err != nil {
		return false, err
	}
	if directCount > 0 {
		return true, nil
	}

	var userRoles []*rbac.Policy
	err = db.GetDBFromContext(c).Where("ptype = ? AND v0 = ? AND deleted = ?", "g", user, false).Find(&userRoles).Error
	if err != nil {
		return false, err
	}

	for _, rolePolicy := range userRoles {
		var roleCount int64
		err = db.GetDBFromContext(c).Model(&rbac.Policy{}).
			Where("ptype = ? AND v0 = ? AND (v1 = ? OR v1 = ?) AND (v2 = ? OR v2 = ?) AND deleted = ?",
				"p", rolePolicy.V1, resource, "*", action, "*", false).
			Count(&roleCount).Error
		if err != nil {
			return false, err
		}
		if roleCount > 0 {
			return true, nil
		}
	}

	return false, nil
}
