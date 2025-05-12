// Package mapper 提供数据模型与数据库交互的映射层，处理账户相关数据操作
// 创建者：Done-0
// 创建时间：2025-05-10
package mapper

import (
	"fmt"

	"github.com/labstack/echo/v4"

	account "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
)

// GetTotalAccounts 获取系统中的总账户数
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - int64: 账户总数
//   - error: 操作过程中的错误
func GetTotalAccounts(c echo.Context) (int64, error) {
	var count int64
	db := utils.GetDBFromContext(c)
	if err := db.Model(&account.Account{}).Where("deleted = ?", false).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("获取用户总数失败: %w", err)
	}
	return count, nil
}

// GetAccountByEmail 根据邮箱获取用户账户信息
// 参数：
//   - c: Echo 上下文
//   - email: 用户邮箱
//
// 返回值：
//   - *account.Account: 账户信息
//   - error: 操作过程中的错误
func GetAccountByEmail(c echo.Context, email string) (*account.Account, error) {
	var user account.Account
	db := utils.GetDBFromContext(c)
	if err := db.Where("email = ? AND deleted = ?", email, false).First(&user).Error; err != nil {
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}
	return &user, nil
}

// GetAccountByAccountID 根据用户 ID 获取账户信息
// 参数：
//   - c: Echo 上下文
//   - accountID: 账户 ID
//
// 返回值：
//   - *account.Account: 账户信息
//   - error: 操作过程中的错误
func GetAccountByAccountID(c echo.Context, accountID int64) (*account.Account, error) {
	var user account.Account
	db := utils.GetDBFromContext(c)
	if err := db.Where("id = ? AND deleted = ?", accountID, false).First(&user).Error; err != nil {
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}
	return &user, nil
}

// CreateAccount 创建新用户
// 参数：
//   - c: Echo 上下文
//   - acc: 账户信息
//
// 返回值：
//   - error: 操作过程中的错误
func CreateAccount(c echo.Context, acc *account.Account) error {
	db := utils.GetDBFromContext(c)
	if err := db.Create(acc).Error; err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}
	return nil
}

// UpdateAccount 更新账户信息
// 参数：
//   - c: Echo 上下文
//   - acc: 账户信息
//
// 返回值：
//   - error: 操作过程中的错误
func UpdateAccount(c echo.Context, acc *account.Account) error {
	db := utils.GetDBFromContext(c)
	if err := db.Save(acc).Error; err != nil {
		return fmt.Errorf("更新账户失败: %w", err)
	}
	return nil
}
