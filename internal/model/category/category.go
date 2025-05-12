// Package model 提供类目数据模型定义
// 创建者：Done-0
// 创建时间：2025-05-10
package model

import "jank.com/jank_blog/internal/model/base"

// Category 类目模型
type Category struct {
	base.Base
	Name        string      `gorm:"type:varchar(255);not null;index" json:"name"`    // 类目名称
	Description string      `gorm:"type:varchar(255);default:''" json:"description"` // 类目描述
	ParentID    int64       `gorm:"index;default:null" json:"parent_id"`             // 父类目ID
	Path        string      `gorm:"type:varchar(225);not null;index" json:"path"`    // 类目路径
	Children    []*Category `gorm:"-" json:"children"`                               // 子类目，不存储在数据库，用于递归构建树结构
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (Category) TableName() string {
	return "categories"
}
