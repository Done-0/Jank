package model

import "jank.com/jank_blog/internal/model/base"

type Category struct {
	base.BaseModel
	Name        string      `gorm:"type:varchar(255);not null;index" json:"name"`        // 类目名称，非空, 索引
	Description string      `gorm:"type:text" json:"description,omitempty"`              // 类目描述, 可空
	ParentID    int64       `gorm:"index;default:null" json:"parent_id,omitempty"`       // 父类目ID, 可空, 索引
	IsActive    bool        `gorm:"type:boolean;default:true;not null" json:"is_active"` // 类目是否活跃（软删除标记）, 默认为 true, 非空
	Path        string      `gorm:"type:varchar(1024);not null" json:"path"`             // 类目路径，用于优化类目路径查询, 非空
	Children    []*Category `gorm:"-" json:"children,omitempty"`                         // 子类目，不存储在数据库，用于递归构建树结构, 可空
}

func (Category) TableName() string {
	return "category"
}
