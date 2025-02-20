package mapper

import (
	"fmt"
	"jank.com/jank_blog/internal/global"
	account "jank.com/jank_blog/internal/model/account"
)

// GetAccountByEmail 根据邮箱获取用户账户信息
func GetAccountByEmail(email string) (*account.Account, error) {
	var user account.Account
	if err := global.DB.Where("email = ? AND deleted = ?", email, 0).First(&user).Error; err != nil {
		return nil, fmt.Errorf("获取用户失败: %v", err)
	}
	return &user, nil
}

// GetAccountByUserID 根据用户ID获取账户信息
func GetAccountByUserID(userID int64) (*account.Account, error) {
	var user account.Account
	if err := global.DB.Where("id = ? AND deleted = ?", userID, 0).First(&user).Error; err != nil {
		return nil, fmt.Errorf("获取用户失败: %v", err)
	}
	return &user, nil
}

// CreateAccount 创建新用户
func CreateAccount(user *account.Account) error {
	if err := global.DB.Create(user).Error; err != nil {
		return fmt.Errorf("创建用户失败: %v", err)
	}
	return nil
}

// UpdateAccount 更新账户信息
func UpdateAccount(account *account.Account) error {
	if err := global.DB.Save(account).Error; err != nil {
		return fmt.Errorf("更新账户失败: %v", err)
	}
	return nil
}

// GetRoleByCode 根据角色编码获取角色
func GetRoleByCode(code string) (*account.Role, error) {
	var role account.Role
	if err := global.DB.Where("code = ? AND deleted = ?", code, 0).First(&role).Error; err != nil {
		return nil, fmt.Errorf("获取角色失败: %v", err)
	}
	return &role, nil
}

// CreateRole 创建角色
func CreateRole(role *account.Role) error {
	if err := global.DB.Create(role).Error; err != nil {
		return fmt.Errorf("创建角色失败: %v", err)
	}
	return nil
}

// UpdateRole 更新角色
func UpdateRole(role *account.Role) error {
	if err := global.DB.Save(role).Error; err != nil {
		return fmt.Errorf("更新角色失败: %v", err)
	}
	return nil
}

// DeleteRoleSoftly 删除角色
func DeleteRoleSoftly(roleID int64) error {
	return global.DB.Model(&account.Role{}).
		Where("id = ?", roleID).
		Update("deleted", 1).Error
}

// GetRoleByID 根据角色ID获取角色
func GetRoleByID(roleID int64) (*account.Role, error) {
	var role account.Role
	if err := global.DB.Where("id = ? AND deleted = ?", roleID, 0).First(&role).Error; err != nil {
		return nil, fmt.Errorf("获取角色失败: %v", err)
	}
	return &role, nil
}

// GetAllRoles 获取所有角色
func GetAllRoles() ([]*account.Role, error) {
	var roles []*account.Role
	if err := global.DB.Where("deleted = ?", 0).Find(&roles).Error; err != nil {
		return nil, fmt.Errorf("获取角色列表失败: %v", err)
	}
	return roles, nil
}

// CreatePermission 创建权限
func CreatePermission(permission *account.Permission) error {
	if err := global.DB.Create(permission).Error; err != nil {
		return fmt.Errorf("创建权限失败: %v", err)
	}
	return nil
}

// UpdatePermission 更新权限
func UpdatePermission(permission *account.Permission) error {
	if err := global.DB.Save(permission).Error; err != nil {
		return fmt.Errorf("更新权限失败: %v", err)
	}
	return nil
}

// DeletePermissionSoftly 删除权限
func DeletePermissionSoftly(permissionID int64) error {
	return global.DB.Model(&account.Permission{}).
		Where("id = ?", permissionID).
		Update("deleted", 1).Error
}

// GetPermissionByID 根据权限ID获取权限
func GetPermissionByID(permissionID int64) (*account.Permission, error) {
	var permission account.Permission
	if err := global.DB.Where("id = ? AND deleted = ?", permissionID, 0).First(&permission).Error; err != nil {
		return nil, fmt.Errorf("获取权限失败: %v", err)
	}
	return &permission, nil
}

// GetAllPermissions 获取所有权限
func GetAllPermissions() ([]*account.Permission, error) {
	var permissions []*account.Permission
	if err := global.DB.Where("deleted = ?", 0).Find(&permissions).Error; err != nil {
		return nil, fmt.Errorf("获取权限列表失败: %v", err)
	}
	return permissions, nil
}

// AssignRoleToUser 为用户分配角色
func AssignRoleToUser(userID, roleID int64) error {
	accountRole := &account.AccountRole{
		AccountID: userID,
		RoleID:    roleID,
	}
	if err := global.DB.Create(accountRole).Error; err != nil {
		return fmt.Errorf("为用户分配角色失败: %v", err)
	}
	return nil
}

// AssignPermissionToRole 为角色分配权限
func AssignPermissionToRole(roleID, permissionID int64) error {
	rolePermission := &account.RolePermission{
		RoleID:       roleID,
		PermissionID: permissionID,
	}
	if err := global.DB.Create(rolePermission).Error; err != nil {
		return fmt.Errorf("为角色分配权限失败: %v", err)
	}
	return nil
}

// DeleteRoleFromUserSoftly 移除用户角色
func DeleteRoleFromUserSoftly(roleID, userID int64) error {
	return global.DB.Model(&account.AccountRole{}).
		Where("role_id = ? AND account_id = ?", roleID, userID).
		Update("deleted", 1).Error
}

// DeletePermissionFromRoleSoftly 移除角色权限
func DeletePermissionFromRoleSoftly(roleID, permissionID int64) error {
	return global.DB.Model(&account.RolePermission{}).
		Where("role_id = ? AND permission = ?", roleID, permissionID).
		Update("deleted", 1).Error
}

// UpdateRoleForUser 更新用户角色
func UpdateRoleForUser(roleID, userID int64) error {
	if err := global.DB.Model(&account.AccountRole{}).
		Where("account_id = ? AND role_id = ? AND deleted = ?", userID, roleID, 0).
		Update("role_id", roleID).Error; err != nil {
		return fmt.Errorf("更新用户角色失败: %v", err)
	}
	return nil
}

// UpdatePermissionForRole 更新角色权限
func UpdatePermissionForRole(roleID, permissionID int64) error {
	// 假设更新角色权限的逻辑
	if err := global.DB.Model(&account.RolePermission{}).
		Where("role_id = ? AND permission_id = ? AND deleted = ?", roleID, permissionID, 0).
		Update("permission_id", permissionID).Error; err != nil {
		return fmt.Errorf("更新角色权限失败: %v", err)
	}
	return nil
}

// GetRolesByUser 根据用户ID获取所有角色
func GetRolesByUser(userID string) ([]*account.Role, error) {
	var roles []*account.Role
	if err := global.DB.Where("user_id = ? AND deleted = ?", userID, 0).Find(&roles).Error; err != nil {
		return nil, fmt.Errorf("查询角色失败: %v", err)
	}
	return roles, nil
}

// GetPermissionsByRole 根据角色ID获取所有权限
func GetPermissionsByRole(roleID string) ([]*account.Permission, error) {
	var permissions []*account.Permission
	if err := global.DB.Where("role_id = ? AND deleted = ?", roleID, 0).Find(&permissions).Error; err != nil {
		return nil, fmt.Errorf("查询权限失败: %v", err)
	}
	return permissions, nil
}
