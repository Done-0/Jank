// Package post 提供文章数据模型定义
// 创建者：Done-0
// 创建时间：2025-08-13
package post

import (
	"github.com/Done-0/jank/internal/model/base"
)

// Post 文章模型
type Post struct {
	base.Base
	Title       string `gorm:"type:varchar(255);not null;index" json:"title"`                 // 标题
	Description string `gorm:"type:varchar(500)" json:"description"`                          // 文章描述/摘要（可选）
	Image       string `gorm:"type:varchar(255)" json:"image"`                                // 图片
	Status      string `gorm:"type:varchar(20);not null;default:'draft';index" json:"status"` // 文章状态
	CategoryID  *int64 `gorm:"type:bigint;index" json:"category_id"`                          // 分类 ID，NULL表示未分类
	Markdown    string `gorm:"type:text" json:"Markdown"`                                     // Markdown 内容
	HTML        string `gorm:"type:text" json:"Html"`                                         // 渲染后的 HTML 内容
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (Post) TableName() string {
	return "posts"
}
