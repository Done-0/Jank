// Package utils 提供各种工具函数，包括数据库事务管理
// 创建者：Done-0
// 创建时间：2025-05-10
package utils

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"jank.com/jank_blog/internal/global"
)

// DB_TRANSACTION_CONTEXT_KEY 事务相关常量
const DB_TRANSACTION_CONTEXT_KEY = "tx" // 存储在Echo 上下文中的数据库事务键名

// GetDBFromContext 从上下文中获取数据库连接
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - *gorm.DB: 数据库连接（事务优先，无事务则返回全局连接）
func GetDBFromContext(c echo.Context) *gorm.DB {
	if tx, ok := c.Get(DB_TRANSACTION_CONTEXT_KEY).(*gorm.DB); ok && tx != nil {
		return tx
	}
	return global.DB
}

// RunDBTransaction 在事务中执行函数
// 参数：
//   - c: Echo 上下文
//   - fn: 事务内执行的函数
//
// 返回值：
//   - error: 执行过程中的错误
func RunDBTransaction(c echo.Context, fn func(error) error) error {
	tx := global.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("开始事务失败: %w", tx.Error)
	}

	c.Set(DB_TRANSACTION_CONTEXT_KEY, tx)
	defer c.Set(DB_TRANSACTION_CONTEXT_KEY, nil)

	// panic 处理
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// 执行业务逻辑
	if err := fn(nil); err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("提交事务失败: %w", err)
	}

	return nil
}
