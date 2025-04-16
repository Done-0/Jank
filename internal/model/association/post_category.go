package model

import (
	"jank.com/jank_blog/internal/model/base"
)

// PostCategory 文章-类目关联模型
type PostCategory struct {
	base.Base
	PostID     int64 `gorm:"type:bigint;not null;index" json:"post_id"`     // 文章ID
	CategoryID int64 `gorm:"type:bigint;not null;index" json:"category_id"` // 类目ID
}

func (PostCategory) TableName() string {
	return "post_categories"
}
