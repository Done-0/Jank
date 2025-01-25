package model

import "jank.com/jank_blog/internal/model/base"

type Category struct {
	base.Base
	Name        string      `gorm:"type:varchar(255);not null;index" json:"name"`    // 类目名称
	Description string      `gorm:"type:varchar(255);default:''" json:"description"` // 类目描述
	ParentID    int64       `gorm:"index;default:null" json:"parent_id"`             // 父类目ID
	Path        string      `gorm:"type:varchar(225);not null;index" json:"path"`    // 类目路径
	Children    []*Category `gorm:"-" json:"children"`                               // 子类目，不存储在数据库，用于递归构建树结构
}

func (Category) TableName() string {
	return "categories"
}
