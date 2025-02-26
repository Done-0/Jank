package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
)

func New(config *configs.Config) {
	tempDB, err := connectDB(config, "postgres")
	if err != nil {
		global.SysLog.Fatalf("数据库连接失败: %v", err)
	}

	if err := createDBIfNotExists(tempDB, config.DBConfig.DBName, config); err != nil {
		global.SysLog.Fatalf("数据库不存在且创建失败: %v", err)
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
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBConfig.DBHost,
		config.DBConfig.DBUser,
		config.DBConfig.DBPassword,
		dbName,
		config.DBConfig.DBPort,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
