package mapper

import (
	"fmt"

	"jank.com/jank_blog/internal/global"
	account "jank.com/jank_blog/internal/model/account"
)

// GetTotalAccounts 获取系统中的总账户数
func GetTotalAccounts() (int64, error) {
	var count int64
	if err := global.DB.Model(&account.Account{}).Where("deleted = ?", false).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("获取用户总数失败: %w", err)
	}
	return count, nil
}

// GetAccountByEmail 根据邮箱获取用户账户信息
func GetAccountByEmail(email string) (*account.Account, error) {
	var user account.Account
	if err := global.DB.Where("email = ? AND deleted = ?", email, false).First(&user).Error; err != nil {
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}
	return &user, nil
}

// GetAccountByAccountID 根据用户 ID 获取账户信息
func GetAccountByAccountID(accountID int64) (*account.Account, error) {
	var user account.Account
	if err := global.DB.Where("id = ? AND deleted = ?", accountID, false).First(&user).Error; err != nil {
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}
	return &user, nil
}

// CreateAccount 创建新用户
func CreateAccount(acc *account.Account) error {
	if err := global.DB.Create(acc).Error; err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}
	return nil
}

// UpdateAccount 更新账户信息
func UpdateAccount(acc *account.Account) error {
	if err := global.DB.Save(acc).Error; err != nil {
		return fmt.Errorf("更新账户失败: %w", err)
	}
	return nil
}
