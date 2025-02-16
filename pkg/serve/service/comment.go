package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
	model "jank.com/jank_blog/internal/model/comment"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/comment"
)

// CreateComment 创建评论
func CreateComment(content string, userId int64, postId int64, replyToCommentId int64, c echo.Context) (*comment.CommentsVo, error) {
	com := &model.Comment{
		Content:          content,
		UserId:           userId,
		PostId:           postId,
		ReplyToCommentId: replyToCommentId,
	}

	if err := mapper.CreateComment(com); err != nil {
		utils.BizLogger(c).Errorf("创建评论失败：%v", err)
		return nil, fmt.Errorf("创建评论失败：%v", err)
	}

	commentVo, err := utils.MapModelToVO(com, &comment.CommentsVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("创建评论时映射 vo 失败：%v", err)
		return nil, fmt.Errorf("创建评论时映射 vo 失败：%v", err)
	}

	return commentVo.(*comment.CommentsVo), nil
}

// GetCommentWithReplies 根据 ID 获取评论及其所有回复
func GetCommentWithReplies(id int64, c echo.Context) (*comment.CommentsVo, error) {
	com, err := mapper.GetCommentByID(id)
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论失败：%v", err)
		return nil, fmt.Errorf("获取评论失败：%v", err)
	}

	// 获取评论的所有回复
	replies, err := mapper.GetReplyByCommentID(id)
	if err != nil {
		utils.BizLogger(c).Errorf("获取子评论失败：%v", err)
		return nil, fmt.Errorf("获取子评论失败：%v", err)
	}

	com.Reply = replies

	commentVo, err := utils.MapModelToVO(com, &comment.CommentsVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论时映射 vo 失败：%v", err)
		return nil, fmt.Errorf("获取评论时映射 vo 失败：%v", err)
	}

	return commentVo.(*comment.CommentsVo), nil
}

// GetCommentGraphByPostID 根据文章 ID 获取评论图结构
func GetCommentGraphByPostID(postID int64, c echo.Context) ([]*comment.CommentsVo, error) {
	comments, err := mapper.GetCommentsByPostID(postID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论图失败：%v", err)
		return nil, fmt.Errorf("获取评论图失败：%v", err)
	}

	commentMap := make(map[int64]*model.Comment)
	var rootComments []*model.Comment

	// 将评论添加到映射
	for i := range comments {
		commentMap[comments[i].ID] = comments[i]
	}

	// 构建图结构
	for i := range comments {
		com := comments[i]
		if com.ReplyToCommentId == 0 {
			// 根评论
			rootComments = append(rootComments, com)
		} else {
			// 回复评论
			if parentComment, exists := commentMap[com.ReplyToCommentId]; exists {
				// 直接将回复加入父评论的回复列表
				if parentComment.Reply == nil {
					parentComment.Reply = make([]*model.Comment, 0)
				}
				// 添加回复
				parentComment.Reply = append(parentComment.Reply, com)
			}
		}
	}

	var rootCommentsVo []*comment.CommentsVo
	for _, rootComment := range rootComments {
		commentVo, err := utils.MapModelToVO(rootComment, &comment.CommentsVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取评论图时映射 vo 失败：%v", err)
			return nil, fmt.Errorf("获取评论图时映射 vo 失败：%v", err)
		}
		rootCommentsVo = append(rootCommentsVo, commentVo.(*comment.CommentsVo))
	}

	return rootCommentsVo, nil
}

// DeleteComment 软删除评论
func DeleteComment(id int64, c echo.Context) (*comment.CommentsVo, error) {
	com, err := mapper.GetCommentByID(id)
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论失败：%v", err)
		return nil, fmt.Errorf("评论不存在：%v", err)
	}

	com.Deleted = true
	if err := mapper.UpdateComment(com); err != nil {
		utils.BizLogger(c).Errorf("软删除评论失败：%v", err)
		return nil, fmt.Errorf("软删除评论失败：%v", err)
	}

	commentVo, err := utils.MapModelToVO(com, &comment.CommentsVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("软删除评论时映射 vo 失败：%v", err)
		return nil, fmt.Errorf("软删除评论时映射 vo 失败：%v", err)
	}

	return commentVo.(*comment.CommentsVo), nil
}
