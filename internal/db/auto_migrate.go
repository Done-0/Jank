// Package db 提供数据库连接和管理功能
// 创建者：Done-0
// 创建时间：2025-08-05
package db

import (
	"fmt"
	"log"

	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/model"
)

// autoMigrate 执行数据库表结构自动迁移
// 返回值：
//
//	error: 错误信息
func autoMigrate() error {
	err := global.DB.AutoMigrate(
		model.GetAllModels()...,
	)
	if err != nil {
		return fmt.Errorf("failed to auto migrate database: %w", err)
	}

	log.Println("Database auto migration succeeded...")
	global.SysLog.Info("Database auto migration succeeded...")

	return nil
}
