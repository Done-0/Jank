package db

import (
	"log"

	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/model"
)

// autoMigrate 执行数据库表结构自动迁移
func autoMigrate() {
	if global.DB == nil {
		log.Fatal("数据库初始化失败，无法执行自动迁移...")
	}

	err := global.DB.AutoMigrate(
		model.GetAllModels()...,
	)
	if err != nil {
		log.Fatalf("数据库自动迁移失败: %v", err)
	}

	log.Println("数据库自动迁移成功...")
	global.SysLog.Infof("数据库自动迁移成功...")
}
