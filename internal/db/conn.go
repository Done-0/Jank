package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
)

func New(config *configs.Config) {
	var dsn string = buildDSN(config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.SysLog.Errorf("MySQL 数据库连接失败: %v", err)
	}

	global.DB = db
	global.SysLog.Infof("MySQL 数据库连接成功！")

	fmt.Println("数据库连接成功")
}

func buildDSN(config *configs.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
}
