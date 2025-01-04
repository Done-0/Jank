package model

import "jank.com/jank_blog/internal/model/base"

type Category struct {
	base.BaseModel
	Name        string      `gorm:"type:varchar(255);not null;index" json:"name"`              // 类目名称
	Description string      `gorm:"type:varchar(255);default:''" json:"description"`           // 类目描述
	ParentID    int64       `gorm:"index;default:null" json:"parent_id"`                       // 父类目ID
	IsActive    bool        `gorm:"type:boolean;default:true;not null;index" json:"is_active"` // 类目是否活跃（软删除标记）, 默认为 true
	Path        string      `gorm:"type:varchar(225);not null;index" json:"path"`              // 类目路径
	Children    []*Category `gorm:"-" json:"children"`                                         // 子类目，不存储在数据库，用于递归构建树结构
}

func (Category) TableName() string {
	return "categories"
}
