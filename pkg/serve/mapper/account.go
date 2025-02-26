package mapper

import (
	"fmt"

	"jank.com/jank_blog/internal/global"
	account "jank.com/jank_blog/internal/model/account"
)

// GetAccountByEmail 根据邮箱获取用户账户信息
func GetAccountByEmail(email string) (*account.Account, error) {
	var user account.Account
	if err := global.DB.Where("email = ? AND deleted = ?", email, false).First(&user).Error; err != nil {
		return nil, fmt.Errorf("获取用户失败: %v", err)
	}
	return &user, nil
}

// GetRoleByAccountID 根据用户 ID 获取角色
func GetRoleByAccountID(accountID int64) (*account.AccountRole, error) {
	var role account.AccountRole
	if err := global.DB.Where("account_id = ? AND deleted = ?", accountID, false).First(&role).Error; err != nil {
		return nil, fmt.Errorf("获取角色失败: %v", err)
	}
	return &role, nil
}

// GetAccountByAccountID 根据用户 ID 获取账户信息
func GetAccountByAccountID(accountID int64) (*account.Account, error) {
	var user account.Account
	if err := global.DB.Where("id = ? AND deleted = ?", accountID, false).First(&user).Error; err != nil {
		return nil, fmt.Errorf("获取用户失败: %v", err)
	}
	return &user, nil
}

// CreateAccount 创建新用户
func CreateAccount(acc *account.Account) error {
	if err := global.DB.Create(acc).Error; err != nil {
		return fmt.Errorf("创建用户失败: %v", err)
	}
	return nil
}

// UpdateAccount 更新账户信息
func UpdateAccount(acc *account.Account) error {
	if err := global.DB.Save(acc).Error; err != nil {
		return fmt.Errorf("更新账户失败: %v", err)
	}
	return nil
}

// GetRoleByCode 根据角色编码获取角色
func GetRoleByCode(code string) (*account.Role, error) {
	var role account.Role
	if err := global.DB.Where("code = ? AND deleted = ?", code, false).First(&role).Error; err != nil {
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
		Update("deleted", true).Error
}

// GetRoleByID 根据角色 ID 获取角色
func GetRoleByID(roleID int64) (*account.Role, error) {
	var role account.Role
	if err := global.DB.Where("id = ? AND deleted = ?", roleID, false).First(&role).Error; err != nil {
		return nil, fmt.Errorf("获取角色失败: %v", err)
	}
	return &role, nil
}

// GetAllRoles 获取所有角色
func GetAllRoles() ([]*account.Role, error) {
	var roles []*account.Role
	if err := global.DB.Where("deleted = ?", false).Find(&roles).Error; err != nil {
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
		Update("deleted", true).Error
}

// GetPermissionByID 根据权限 ID 获取权限
func GetPermissionByID(permissionID int64) (*account.Permission, error) {
	var permission account.Permission
	if err := global.DB.Where("id = ? AND deleted = ?", permissionID, false).First(&permission).Error; err != nil {
		return nil, fmt.Errorf("获取权限失败: %v", err)
	}
	return &permission, nil
}

// GetAllPermissions 获取所有权限
func GetAllPermissions() ([]*account.Permission, error) {
	var permissions []*account.Permission
	if err := global.DB.Where("deleted = ?", false).Find(&permissions).Error; err != nil {
		return nil, fmt.Errorf("获取权限列表失败: %v", err)
	}
	return permissions, nil
}

// AssignRoleToAcc 为用户分配角色
func AssignRoleToAcc(accountID, roleID int64) error {
	accountRole := &account.AccountRole{
		AccountID: accountID,
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

// DeleteRoleFromAccSoftly 移除用户角色
func DeleteRoleFromAccSoftly(accountID int64, roleID int64) error {
	return global.DB.Model(&account.AccountRole{}).
		Where("account_id = ? AND role_id = ?", accountID, roleID).
		Update("deleted", true).Error
}

// DeletePermissionFromRoleSoftly 移除角色权限
func DeletePermissionFromRoleSoftly(roleID, permissionID int64) error {
	return global.DB.Model(&account.RolePermission{}).
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).
		Update("deleted", true).Error
}

// UpdateRoleForAcc 更新用户角色
func UpdateRoleForAcc(AccountID, roleID int64) error {
	if err := global.DB.Model(&account.AccountRole{}).
		Where("account_id = ? AND role_id = ? AND deleted = ?", AccountID, roleID, false).
		Update("role_id", roleID).Error; err != nil {
		return fmt.Errorf("更新用户角色失败: %v", err)
	}
	return nil
}

// UpdatePermissionForRole 更新角色权限
func UpdatePermissionForRole(roleID, permissionID int64) error {
	if err := global.DB.Model(&account.RolePermission{}).
		Where("role_id = ? AND permission_id = ? AND deleted = ?", roleID, permissionID, false).
		Update("permission_id", permissionID).Error; err != nil {
		return fmt.Errorf("更新角色权限失败: %v", err)
	}
	return nil
}

// GetRolesByAccountID 根据用户 ID 获取所有角色
func GetRolesByAccountID(accountID string) ([]*account.AccountRole, error) {
	var roles []*account.AccountRole
	if err := global.DB.Where("account_id = ? AND deleted = ?", accountID, false).Find(&roles).Error; err != nil {
		return nil, fmt.Errorf("查询角色失败: %v", err)
	}
	return roles, nil
}

// GetPermissionsByRoleID 根据角色 ID 获取所有权限
func GetPermissionsByRoleID(roleID string) ([]*account.RolePermission, error) {
	var permissions []*account.RolePermission
	if err := global.DB.Where("role_id = ? AND deleted = ?", roleID, false).Find(&permissions).Error; err != nil {
		return nil, fmt.Errorf("查询权限失败: %v", err)
	}
	return permissions, nil
}
