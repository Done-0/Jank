// Package model 提供评论数据模型定义
// 创建者：Done-0
// 创建时间：2025-05-10
package model

import "jank.com/jank_blog/internal/model/base"

// Comment 评论模型
type Comment struct {
	base.Base
	Content          string     `gorm:"type:varchar(1024);not null" json:"content"`          // 评论内容
	UserId           int64      `gorm:"type:bigint;not null;index" json:"user_id"`           // 所属用户ID
	PostId           int64      `gorm:"type:bigint;not null;index" json:"post_id"`           // 所属文章ID
	ReplyToCommentId int64      `gorm:"type:bigint;default:null" json:"reply_to_comment_id"` // 目标评论ID
	Replies          []*Comment `gorm:"-" json:"replies"`                                    // 子评论列表，用于构建图结构
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (Comment) TableName() string {
	return "comments"
}
