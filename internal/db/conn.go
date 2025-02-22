package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
)

func New(config *configs.Config) {
	tempDB, err := connectDB(config, config.DBConfig.DBName)
	if err != nil {
		global.SysLog.Fatalf("「%s」数据库连接失败: %v", config.DBConfig.DBName, err)
	}

	if err := createDBIfNotExists(tempDB, config.DBConfig.DBName); err != nil {
		global.SysLog.Fatalf("「%s」数据库不存在: %v", config.DBConfig.DBName, err)
	}

	global.DB, err = connectDB(config, config.DBConfig.DBName)
	if err != nil {
		global.SysLog.Fatalf("「%s」数据库连接失败: %v", config.DBConfig.DBName, err)
	}

	log.Printf("「%s」数据库连接成功 \n", config.DBConfig.DBName)
	global.SysLog.Infof("「%s」数据库连接成功", config.DBConfig.DBName)
}

// connectDB 连接数据库
func connectDB(config *configs.Config, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBConfig.DBUser,
		config.DBConfig.DBPassword,
		config.DBConfig.DBHost,
		config.DBConfig.DBPort,
		dbName,
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// createDBIfNotExists 检查目标数据库是否存在，不存在则创建
func createDBIfNotExists(db *gorm.DB, dbName string) error {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM information_schema.schemata WHERE schema_name = ?)"
	if err := db.Raw(query, dbName).Scan(&exists).Error; err != nil {
		log.Printf("查询「%s」数据库是否存在时失败: %v", dbName, err)
		return fmt.Errorf("查询「%s」数据库是否存在时失败: %v", dbName, err)
	}

	if !exists {
		log.Printf("「%s」数据库不存在，正在创建...", dbName)
		global.SysLog.Infof("「%s」数据库不存在，正在创建...", dbName)
		return db.Exec(fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", dbName)).Error
	}
	return nil
}
