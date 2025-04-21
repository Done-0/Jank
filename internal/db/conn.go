package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
)

const (
	DIALECT_POSTGRES = "postgres"
	DIALECT_SQLITE   = "sqlite"
	DIALECT_MYSQL    = "mysql"
)

// New 初始化数据库连接
func New(config *configs.Config) {
	var err error

	switch config.DBConfig.DBDialect {
	case DIALECT_SQLITE:
		global.DB, err = connectToDB(config, config.DBConfig.DBName)
		if err != nil {
			global.SysLog.Fatalf("连接 SQLite 数据库失败: %v", err)
		}
	case DIALECT_POSTGRES, DIALECT_MYSQL:
		systemDB, err := connectToSystemDB(config)
		if err != nil {
			global.SysLog.Fatalf("连接系统数据库失败: %v", err)
		}

		if err := ensureDBExists(systemDB, config); err != nil {
			global.SysLog.Fatalf("数据库不存在且创建失败: %v", err)
		}

		sqlDB, _ := systemDB.DB()
		sqlDB.Close()

		global.DB, err = connectToDB(config, config.DBConfig.DBName)
		if err != nil {
			global.SysLog.Fatalf("连接数据库失败: %v", err)
		}
	default:
		global.SysLog.Fatalf("不支持的数据库类型: %s", config.DBConfig.DBDialect)
	}

	log.Printf("「%s」数据库连接成功...", config.DBConfig.DBName)
	global.SysLog.Infof("「%s」数据库连接成功！", config.DBConfig.DBName)

	autoMigrate()
}

// connectToSystemDB 连接到系统数据库
func connectToSystemDB(config *configs.Config) (*gorm.DB, error) {
	switch config.DBConfig.DBDialect {
	case DIALECT_POSTGRES:
		return connectToDB(config, "postgres")
	case DIALECT_MYSQL:
		return connectToDB(config, "information_schema")
	default:
		return nil, fmt.Errorf("不支持的数据库类型: %s", config.DBConfig.DBDialect)
	}
}

// ensureDBExists 确保数据库存在，不存在则创建
func ensureDBExists(db *gorm.DB, config *configs.Config) error {
	switch config.DBConfig.DBDialect {
	case DIALECT_POSTGRES:
		return ensurePostgresDBExists(db, config.DBConfig.DBName, config.DBConfig.DBUser)
	case DIALECT_MYSQL:
		return ensureMySQLDBExists(db, config.DBConfig.DBName)
	default:
		return nil
	}
}

// connectToDB 连接到指定数据库
func connectToDB(config *configs.Config, dbName string) (*gorm.DB, error) {
	dialector, err := getDialector(config, dbName)
	if err != nil {
		return nil, fmt.Errorf("获取数据库驱动器失败: %v", err)
	}

	return gorm.Open(dialector, &gorm.Config{})
}

// getDialector 根据数据库类型获取对应的驱动器
func getDialector(config *configs.Config, dbName string) (gorm.Dialector, error) {
	switch config.DBConfig.DBDialect {
	case DIALECT_POSTGRES:
		return getPostgresDialector(config, dbName), nil
	case DIALECT_SQLITE:
		return getSqliteDialector(config, dbName)
	case DIALECT_MYSQL:
		return getMySQLDialector(config, dbName), nil
	default:
		return nil, fmt.Errorf("不支持的数据库类型: %s", config.DBConfig.DBDialect)
	}
}

// getPostgresDialector 获取 PostgreSQL 驱动器
func getPostgresDialector(config *configs.Config, dbName string) gorm.Dialector {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBConfig.DBHost,
		config.DBConfig.DBUser,
		config.DBConfig.DBPassword,
		dbName,
		config.DBConfig.DBPort,
	)
	return postgres.Open(dsn)
}

// getMySQLDialector 获取 MySQL 驱动器
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
func getSqliteDialector(config *configs.Config, dbName string) (gorm.Dialector, error) {
	if err := os.MkdirAll(config.DBConfig.DBPath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("创建 SQLite 数据库目录失败: %v", err)
	}

	dbPath := filepath.Join(config.DBConfig.DBPath, dbName+".db")
	return sqlite.Open(dbPath), nil
}

// ensurePostgresDBExists 确保 PostgreSQL 数据库存在，不存在则创建
func ensurePostgresDBExists(db *gorm.DB, dbName, dbUser string) error {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = ?)"
	if err := db.Raw(query, dbName).Scan(&exists).Error; err != nil {
		log.Printf("查询「%s」数据库是否存在时失败: %v", dbName, err)
		return fmt.Errorf("查询「%s」数据库是否存在时失败: %v", dbName, err)
	}

	if !exists {
		log.Printf("「%s」数据库不存在，正在创建...", dbName)
		global.SysLog.Infof("「%s」数据库不存在，正在创建...", dbName)

		createSQL := fmt.Sprintf("CREATE DATABASE %s ENCODING 'UTF8' OWNER %s", dbName, dbUser)
		if err := db.Exec(createSQL).Error; err != nil {
			return fmt.Errorf("创建「%s」数据库失败: %v", dbName, err)
		}
	}
	return nil
}

// ensureMySQLDBExists 确保 MySQL 数据库存在，不存在则创建
func ensureMySQLDBExists(db *gorm.DB, dbName string) error {
	var count int64
	query := "SELECT COUNT(*) FROM information_schema.schemata WHERE schema_name = ?"
	if err := db.Raw(query, dbName).Scan(&count).Error; err != nil {
		log.Printf("查询「%s」数据库是否存在时失败: %v", dbName, err)
		return fmt.Errorf("查询「%s」数据库是否存在时失败: %v", dbName, err)
	}

	if count == 0 {
		log.Printf("「%s」数据库不存在，正在创建...", dbName)
		global.SysLog.Infof("「%s」数据库不存在，正在创建...", dbName)
		createSQL := fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci", dbName)
		if err := db.Exec(createSQL).Error; err != nil {
			return fmt.Errorf("创建「%s」数据库失败: %v", dbName, err)
		}
	}
	return nil
}
