// Package db 提供数据库连接和管理功能
// 创建者：Done-0
// 创建时间：2025-08-05
package db

import (
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Done-0/jank/configs"
)

// getDialector 根据数据库类型获取对应的驱动器
// 参数：
//
//	config: 应用配置
//	dbName: 数据库名称
//
// 返回值：
//
//	gorm.Dialector: 数据库方言
//	error: 获取方言过程中的错误
func getDialector(config *configs.Config, dbName string) (gorm.Dialector, error) {
	switch config.DBConfig.DBDialect {
	case DIALECT_POSTGRES:
		return getPostgresDialector(config, dbName), nil
	case DIALECT_SQLITE:
		return getSqliteDialector(config, dbName)
	case DIALECT_MYSQL:
		return getMySQLDialector(config, dbName), nil
	default:
		return nil, fmt.Errorf("unsupported database dialect: %s", config.DBConfig.DBDialect)
	}
}

// getPostgresDialector 获取 PostgreSQL 驱动器
// 参数：
//
//	config: 应用配置
//	dbName: 数据库名称
//
// 返回值：
//
//	gorm.Dialector: PostgreSQL 方言
func getPostgresDialector(config *configs.Config, dbName string) gorm.Dialector {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai client_encoding=UTF8",
		config.DBConfig.DBHost,
		config.DBConfig.DBUser,
		config.DBConfig.DBPassword,
		dbName,
		config.DBConfig.DBPort,
	)
	return postgres.Open(dsn)
}

// getMySQLDialector 获取 MySQL 驱动器
// 参数：
//
//	config: 应用配置
//	dbName: 数据库名称
//
// 返回值：
//
//	gorm.Dialector: MySQL 方言
func getMySQLDialector(config *configs.Config, dbName string) gorm.Dialector {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBConfig.DBUser,
		config.DBConfig.DBPassword,
		config.DBConfig.DBHost,
		config.DBConfig.DBPort,
		dbName,
	)
	return mysql.Open(dsn)
}

// getSqliteDialector 获取 SQLite 驱动器并确保目录存在
// 参数：
//
//	config: 应用配置
//	dbName: 数据库名称
//
// 返回值：
//
//	gorm.Dialector: SQLite 方言
//	error: 创建目录过程中的错误
func getSqliteDialector(config *configs.Config, dbName string) (gorm.Dialector, error) {
	if err := os.MkdirAll(config.DBConfig.DBPath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create SQLite database directory: %v", err)
	}

	dbPath := filepath.Join(config.DBConfig.DBPath, dbName+".db")
	return sqlite.Open(dbPath), nil
}
