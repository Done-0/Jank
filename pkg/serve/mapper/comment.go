package mapper

import (
	"fmt"

	"jank.com/jank_blog/internal/global"
	model "jank.com/jank_blog/internal/model/comment"
)

// CreateComment 保存评论到数据库
func CreateComment(comment *model.Comment) error {
	if err := global.DB.Create(comment).Error; err != nil {
		return fmt.Errorf("创建评论失败: %w", err)
	}
	return nil
}

// GetCommentByID 根据 ID 查询评论
func GetCommentByID(id int64) (*model.Comment, error) {
	var comment model.Comment
	if err := global.DB.Where("id = ? AND deleted = ?", id, false).First(&comment).Error; err != nil {
		return nil, fmt.Errorf("获取评论失败: %w", err)
	}
	return &comment, nil
}

// GetReplyByCommentID 获取评论的所有回复
func GetReplyByCommentID(id int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	if err := global.DB.Where("reply_to_comment_id = ? AND deleted = ?", id, false).Find(&comments).Error; err != nil {
		return nil, fmt.Errorf("获取评论回复失败: %w", err)
	}
	return comments, nil
}

// GetCommentsByPostID 根据文章 ID 查询所有评论
func GetCommentsByPostID(postID int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	if err := global.DB.Where("post_id = ? AND deleted = ?", postID, false).Find(&comments).Error; err != nil {
		return nil, fmt.Errorf("获取文章评论失败: %w", err)
	}
	return comments, nil
}

// UpdateComment 更新评论
func UpdateComment(comment *model.Comment) error {
	if err := global.DB.Save(comment).Error; err != nil {
		return fmt.Errorf("更新评论失败: %w", err)
	}
	return nil
}
