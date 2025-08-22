// Package db 提供各种数据库工具函数，包括事务管理
// 创建者：Done-0
// 创建时间：2025-08-10
package db

import (
	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"

	"github.com/Done-0/jank/internal/global"
)

// DB_TRANSACTION_CONTEXT_KEY 事务相关常量
const DB_TRANSACTION_CONTEXT_KEY = "tx" // 存储在 Hertz 上下文中的数据库事务键名

// GetDBFromContext 从上下文中获取数据库连接
// 参数：
//
//	c: Hertz 请求上下文
//
// 返回值：
//
//	*gorm.DB: 数据库连接（事务优先，无事务则返回全局连接）
func GetDBFromContext(c *app.RequestContext) *gorm.DB {
	if tx, exists := c.Get(DB_TRANSACTION_CONTEXT_KEY); exists {
		if db, ok := tx.(*gorm.DB); ok && db != nil {
			return db
		}
	}
	return global.DB
}

// RunDBTransaction 在事务中执行函数
// 参数：
//
//	c: Hertz 请求上下文
//	fn: 事务内执行的函数
//
// 返回值：
//
//	T: 函数返回的结果
//	error: 执行过程中的错误
func RunDBTransaction[T any](c *app.RequestContext, fn func() (T, error)) (T, error) {
	var zero T
	tx := global.DB.Begin()
	if tx.Error != nil {
		return zero, tx.Error
	}

	c.Set(DB_TRANSACTION_CONTEXT_KEY, tx)
	defer c.Set(DB_TRANSACTION_CONTEXT_KEY, nil)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	result, err := fn()
	if err != nil {
		tx.Rollback()
		return zero, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return zero, err
	}

	return result, nil
}
