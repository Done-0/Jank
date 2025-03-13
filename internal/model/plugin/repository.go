package model

import (
	"jank.com/jank_blog/internal/model/base"
)

// Repository 插件仓库模型
type Repository struct {
	base.Base
	Name        string `gorm:"type:varchar(128);not null" json:"name"` // Name 仓库名称
	URL         string `gorm:"type:varchar(255);not null" json:"url"`  // URL 仓库地址
	Description string `gorm:"type:varchar(2048)" json:"description"`  // Description 仓库描述
}

func (Repository) TableName() string {
	return "repositories"
}
