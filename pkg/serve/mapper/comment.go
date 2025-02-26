package mapper

import (
	"jank.com/jank_blog/internal/global"
	model "jank.com/jank_blog/internal/model/comment"
)

// CreateComment 保存评论到数据库
func CreateComment(comment *model.Comment) error {
	return global.DB.Create(comment).Error
}

// GetCommentByID 根据 ID 查询评论
func GetCommentByID(id int64) (*model.Comment, error) {
	var comment model.Comment
	err := global.DB.Where("id = ? AND deleted = ?", id, false).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetReplyByCommentID 获取评论的所有回复
func GetReplyByCommentID(id int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := global.DB.Where("reply_to_comment_id = ? AND deleted = ?", id, false).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// GetCommentsByPostID 根据文章 ID 查询所有评论
func GetCommentsByPostID(postID int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := global.DB.Where("post_id = ? AND deleted = ?", postID, false).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// UpdateComment 更新评论
func UpdateComment(comment *model.Comment) error {
	return global.DB.Save(comment).Error
}
