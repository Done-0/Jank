package db

import (
	"fmt"
	"log"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/utils"
)

const (
	DialectPostgres = "postgres"
	DialectSqlite   = "sqlite"
)

func New(config *configs.Config) {
	tempDB, err := connectDB(config, config.DBConfig.DBName)
	if err != nil {
		global.SysLog.Fatalf("数据库连接失败: %v", err)
	}

	// 仅 PostgreSQL 需要创建数据库
	if config.DBConfig.Dialect == DialectPostgres {
		if err := createDBIfNotExists(tempDB, config.DBConfig.DBName, config); err != nil {
			global.SysLog.Fatalf("数据库不存在且创建失败: %v", err)
		}
	}

	global.DB, err = connectDB(config, config.DBConfig.DBName)
	if err != nil {
		global.SysLog.Fatalf("连接数据库失败: %v", err)
	}
	log.Printf("「%s」数据库连接成功...", config.DBConfig.DBName)
	global.SysLog.Infof("「%s」数据库连接成功！", config.DBConfig.DBName)

	autoMigrate()
}

// connectDB 连接到指定数据库
func connectDB(config *configs.Config, dbName string) (*gorm.DB, error) {
	_, dialector, err := getDSNAndDialector(config, dbName)
	if err != nil {
		return nil, err
	}

	return gorm.Open(dialector, &gorm.Config{})
}

func getDSNAndDialector(config *configs.Config, dbName string) (string, gorm.Dialector, error) {
	var (
		dsn       string
		dialector gorm.Dialector
	)

	switch config.DBConfig.Dialect {
	case DialectPostgres:
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			config.DBConfig.DBHost,
			config.DBConfig.DBUser,
			config.DBConfig.DBPassword,
			dbName,
			config.DBConfig.DBPort,
		)
		dialector = postgres.Open(dsn)

	case DialectSqlite:
		if err := utils.MkDir(config.DBConfig.DBPath); err != nil {
			return "", nil, fmt.Errorf("创建sqlite数据库目录失败: %v", err)
		}

		dbPath := filepath.Join(config.DBConfig.DBPath, dbName+".db")
		dsn = dbPath
		dialector = sqlite.Open(dsn)

	default:
		return "", nil, fmt.Errorf("不支持的数据库类型: %s", config.DBConfig.Dialect)
	}

	return dsn, dialector, nil
}

// createDBIfNotExists 检查目标数据库是否存在，不存在则创建
func createDBIfNotExists(db *gorm.DB, dbName string, config *configs.Config) error {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = ?)"
	if err := db.Raw(query, dbName).Scan(&exists).Error; err != nil {
		log.Printf("查询「%s」数据库是否存在时失败: %v", dbName, err)
		return fmt.Errorf("查询「%s」数据库是否存在时失败: %v", dbName, err)
	}

	if !exists {
		log.Printf("「%s」数据库不存在，正在创建...", dbName)
		global.SysLog.Infof("「%s」数据库不存在，正在创建...", dbName)
		if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s ENCODING 'UTF8' OWNER %s", dbName, config.DBConfig.DBUser)).Error; err != nil {
			return fmt.Errorf("创建「%s」数据库失败: %v", dbName, err)
		}
	}
	return nil
}
