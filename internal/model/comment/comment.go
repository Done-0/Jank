package model

import "jank.com/jank_blog/internal/model/base"

type Comment struct {
	base.BaseModel
	Content          string `gorm:"type:varchar(1024);not null" json:"content"`       // 评论内容
	UserId           int64  `gorm:"type:int;not null;index" json:"user_id"`           // 所属用户ID
	PostId           int64  `gorm:"type:bigint;not null;index" json:"post_id"`        // 所属文章ID
	RootCommentId    int64  `gorm:"type:bigint;default:0" json:"root_comment_id"`     // 顶级评论ID
	ReplyToCommentId int64  `gorm:"type:bigint;default:0" json:"reply_to_comment_id"` // 目标评论ID
}

func (Comment) TableName() string {
	return "comments"
}
