// Package db 提供数据库初始化管理员用户功能
// 创建者：Done-0
// 创建时间：2025-08-22
package db

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/model/rbac"
	"github.com/Done-0/jank/internal/model/user"
)

// InitAdminUser 初始化管理员用户
// 如果数据库中没有用户，则创建默认的super_admin管理员账户
func InitAdminUser(config *configs.Config) {
	var userCount int64
	if err := global.DB.Model(&user.User{}).Where("deleted = ?", false).Count(&userCount).Error; err != nil {
		global.SysLog.Errorf("Failed to count users: %v", err)
		return
	}

	if userCount > 0 {
		global.SysLog.Info("Users already exist, skipping admin user initialization")
		return
	}

	tx := global.DB.Begin()
	if tx.Error != nil {
		global.SysLog.Errorf("Failed to begin transaction: %v", tx.Error)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			global.SysLog.Errorf("Transaction rolled back due to panic: %v", r)
		}
	}()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.AppConfig.User.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		global.SysLog.Errorf("Failed to hash password: %v", err)
		return
	}

	adminUser := &user.User{
		Email:    config.AppConfig.User.AdminEmail,
		Password: string(hashedPassword),
		Nickname: config.AppConfig.User.AdminNickname,
		Role:     config.AppConfig.User.AdminRole,
	}

	if err := tx.Create(adminUser).Error; err != nil {
		tx.Rollback()
		global.SysLog.Errorf("Failed to create admin user: %v", err)
		return
	}

	userIDStr := strconv.FormatInt(adminUser.ID, 10)
	rolePolicy := &rbac.Policy{
		Ptype: "g",
		V0:    userIDStr,
		V1:    config.AppConfig.User.AdminRole,
	}

	if err := tx.Create(rolePolicy).Error; err != nil {
		tx.Rollback()
		global.SysLog.Errorf("Failed to assign admin role: %v", err)
		return
	}

	if err := tx.Commit().Error; err != nil {
		global.SysLog.Errorf("Failed to commit transaction: %v", err)
		return
	}

	global.SysLog.Infof("Admin user initialized successfully: %s (ID: %d) with role '%s'",
		adminUser.Email, adminUser.ID, config.AppConfig.User.AdminRole)
}
