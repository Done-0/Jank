package service

import (
	"fmt"

	"github.com/labstack/echo/v4"

	model "jank.com/jank_blog/internal/model/comment"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/comment/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/comment"
)

// CreateComment 创建评论
func CreateComment(req *dto.CreateCommentRequest, c echo.Context) (*comment.CommentsVO, error) {
	com := &model.Comment{
		Content:          req.Content,
		UserId:           req.UserId,
		PostId:           req.PostId,
		ReplyToCommentId: req.ReplyToCommentId,
	}

	if err := mapper.CreateComment(com); err != nil {
		utils.BizLogger(c).Errorf("创建评论失败：%v", err)
		return nil, fmt.Errorf("创建评论失败：%w", err)
	}

	commentVO, err := utils.MapModelToVO(com, &comment.CommentsVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("创建评论时映射 VO 失败：%v", err)
		return nil, fmt.Errorf("创建评论时映射 VO 失败：%w", err)
	}

	return commentVO.(*comment.CommentsVO), nil
}

// GetCommentWithReplies 根据 ID 获取评论及其所有回复
func GetCommentWithReplies(req *dto.GetOneCommentRequest, c echo.Context) (*comment.CommentsVO, error) {
	com, err := mapper.GetCommentByID(req.CommentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论失败：%v", err)
		return nil, fmt.Errorf("获取评论失败：%w", err)
	}

	replies, err := mapper.GetReplyByCommentID(req.CommentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取子评论失败：%v", err)
		return nil, fmt.Errorf("获取子评论失败：%w", err)
	}

	com.Replies = replies

	commentVO, err := utils.MapModelToVO(com, &comment.CommentsVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论时映射 VO 失败：%v", err)
		return nil, fmt.Errorf("获取评论时映射 VO 失败：%w", err)
	}

	return commentVO.(*comment.CommentsVO), nil
}

// GetCommentGraphByPostID 根据文章 ID 获取评论图结构
func GetCommentGraphByPostID(req *dto.GetCommentGraphRequest, c echo.Context) ([]*comment.CommentsVO, error) {
	comments, err := mapper.GetCommentsByPostID(req.PostID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论图失败：%v", err)
		return nil, fmt.Errorf("获取评论图失败：%w", err)
	}

	commentMap := make(map[int64]*comment.CommentsVO)
	var rootCommentsVO []*comment.CommentsVO

	for _, com := range comments {
		commentVO, err := utils.MapModelToVO(com, &comment.CommentsVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取评论图时映射 VO 失败：%v", err)
			return nil, fmt.Errorf("获取评论图时映射 VO 失败：%w", err)
		}
		vo := commentVO.(*comment.CommentsVO)
		vo.Replies = make([]*comment.CommentsVO, 0)
		commentMap[com.ID] = vo

		if com.ReplyToCommentId == 0 {
			rootCommentsVO = append(rootCommentsVO, vo)
		}
	}

	for _, com := range comments {
		if com.ReplyToCommentId != 0 {
			if parentVO, exists := commentMap[com.ReplyToCommentId]; exists {
				parentVO.Replies = append(parentVO.Replies, commentMap[com.ID])
			}
		}
	}

	processed := make(map[int64]bool)
	var processComment func(*comment.CommentsVO) *comment.CommentsVO
	processComment = func(vo *comment.CommentsVO) *comment.CommentsVO {
		if processed[vo.ID] {
			newVO := *vo
			newVO.Replies = make([]*comment.CommentsVO, 0)
			return &newVO
		}
		processed[vo.ID] = true

		for i, reply := range vo.Replies {
			vo.Replies[i] = processComment(reply)
		}
		return vo
	}

	for i, rootVO := range rootCommentsVO {
		rootCommentsVO[i] = processComment(rootVO)
	}

	return rootCommentsVO, nil
}

// DeleteComment 软删除评论
func DeleteComment(req *dto.DeleteCommentRequest, c echo.Context) (*comment.CommentsVO, error) {
	com, err := mapper.GetCommentByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论失败：%v", err)
		return nil, fmt.Errorf("评论不存在：%w", err)
	}

	com.Deleted = true
	if err := mapper.UpdateComment(com); err != nil {
		utils.BizLogger(c).Errorf("软删除评论失败：%v", err)
		return nil, fmt.Errorf("软删除评论失败：%w", err)
	}

	commentVO, err := utils.MapModelToVO(com, &comment.CommentsVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("软删除评论时映射 VO 失败：%v", err)
		return nil, fmt.Errorf("软删除评论时映射 VO 失败：%w", err)
	}

	return commentVO.(*comment.CommentsVO), nil
}
