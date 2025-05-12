// Package model 提供实体关联数据模型定义
// 创建者：Done-0
// 创建时间：2025-05-10
package model

import (
	"jank.com/jank_blog/internal/model/base"
)

// PostCategory 文章-类目关联模型
type PostCategory struct {
	base.Base
	PostID     int64 `gorm:"type:bigint;not null;index" json:"post_id"` // 文章ID
	CategoryID int64 `gorm:"type:bigint;index" json:"category_id"`      // 类目ID
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (PostCategory) TableName() string {
	return "post_categories"
}
