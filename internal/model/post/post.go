package model

import "jank.com/jank_blog/internal/model/base"

// Post 博客文章模型
type Post struct {
	base.Base
	Title           string `gorm:"type:varchar(255);not null;index" json:"title"`                                    // 标题
	Image           string `gorm:"type:varchar(255)" json:"image"`                                                   // 图片
	Visibility      string `gorm:"type:enum('public','private');not null;index" json:"visibility" default:"private"` // 可见性
	ContentMarkdown string `gorm:"type:text" json:"contentMarkdown"`                                                 // Markdown 内容
	ContentHTML     string `gorm:"type:text" json:"contentHtml"`                                                     // 渲染后的 HTML 内容
	CategoryIDs     string `gorm:"type:text" json:"categoryIds"`                                                     // 类目 ID 列表，以逗号分隔
}

func (Post) TableName() string {
	return "posts"
}
