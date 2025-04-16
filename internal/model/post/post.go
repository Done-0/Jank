package model

import (
	"jank.com/jank_blog/internal/model/base"
)

// Post 博客文章模型
type Post struct {
	base.Base
	Title           string `gorm:"type:varchar(255);not null;index" json:"title"`               // 标题
	Image           string `gorm:"type:varchar(255)" json:"image"`                              // 图片
	Visibility      bool   `gorm:"type:boolean;not null;default:false;index" json:"visibility"` // 可见性，默认不可见
	ContentMarkdown string `gorm:"type:text" json:"contentMarkdown"`                            // Markdown 内容
	ContentHTML     string `gorm:"type:text" json:"contentHtml"`                                // 渲染后的 HTML 内容
}

func (Post) TableName() string {
	return "posts"
}
