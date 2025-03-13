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
		return nil, fmt.Errorf("创建评论失败：%v", err)
	}

	commentVo, err := utils.MapModelToVO(com, &comment.CommentsVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("创建评论时映射 vo 失败：%v", err)
		return nil, fmt.Errorf("创建评论时映射 vo 失败：%v", err)
	}

	return commentVo.(*comment.CommentsVO), nil
}

// GetCommentWithReplies 根据 ID 获取评论及其所有回复
func GetCommentWithReplies(req *dto.GetOneCommentRequest, c echo.Context) (*comment.CommentsVO, error) {
	com, err := mapper.GetCommentByID(req.CommentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论失败：%v", err)
		return nil, fmt.Errorf("获取评论失败：%v", err)
	}

	replies, err := mapper.GetReplyByCommentID(req.CommentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取子评论失败：%v", err)
		return nil, fmt.Errorf("获取子评论失败：%v", err)
	}

	com.Replies = replies

	commentVo, err := utils.MapModelToVO(com, &comment.CommentsVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论时映射 vo 失败：%v", err)
		return nil, fmt.Errorf("获取评论时映射 vo 失败：%v", err)
	}

	return commentVo.(*comment.CommentsVO), nil
}

// GetCommentGraphByPostID 根据文章 ID 获取评论图结构
func GetCommentGraphByPostID(req *dto.GetCommentGraphRequest, c echo.Context) ([]*comment.CommentsVO, error) {
	comments, err := mapper.GetCommentsByPostID(req.PostID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论图失败：%v", err)
		return nil, fmt.Errorf("获取评论图失败：%v", err)
	}

	commentMap := make(map[int64]*comment.CommentsVO)
	var rootCommentsVo []*comment.CommentsVO

	for _, com := range comments {
		commentVo, err := utils.MapModelToVO(com, &comment.CommentsVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取评论图时映射 vo 失败：%v", err)
			return nil, fmt.Errorf("获取评论图时映射 vo 失败：%v", err)
		}
		vo := commentVo.(*comment.CommentsVO)
		vo.Replies = make([]*comment.CommentsVO, 0)
		commentMap[com.ID] = vo

		if com.ReplyToCommentId == 0 {
			rootCommentsVo = append(rootCommentsVo, vo)
		}
	}

	for _, com := range comments {
		if com.ReplyToCommentId != 0 {
			if parentVo, exists := commentMap[com.ReplyToCommentId]; exists {
				parentVo.Replies = append(parentVo.Replies, commentMap[com.ID])
			}
		}
	}

	processed := make(map[int64]bool)
	var processComment func(*comment.CommentsVO) *comment.CommentsVO
	processComment = func(vo *comment.CommentsVO) *comment.CommentsVO {
		if processed[vo.ID] {
			newVo := *vo
			newVo.Replies = make([]*comment.CommentsVO, 0)
			return &newVo
		}
		processed[vo.ID] = true

		for i, reply := range vo.Replies {
			vo.Replies[i] = processComment(reply)
		}
		return vo
	}

	for i, rootVo := range rootCommentsVo {
		rootCommentsVo[i] = processComment(rootVo)
	}

	return rootCommentsVo, nil
}

// DeleteComment 软删除评论
func DeleteComment(req *dto.DeleteCommentRequest, c echo.Context) (*comment.CommentsVO, error) {
	com, err := mapper.GetCommentByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论失败：%v", err)
		return nil, fmt.Errorf("评论不存在：%v", err)
	}

	com.Deleted = true
	if err := mapper.UpdateComment(com); err != nil {
		utils.BizLogger(c).Errorf("软删除评论失败：%v", err)
		return nil, fmt.Errorf("软删除评论失败：%v", err)
	}

	commentVo, err := utils.MapModelToVO(com, &comment.CommentsVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("软删除评论时映射 vo 失败：%v", err)
		return nil, fmt.Errorf("软删除评论时映射 vo 失败：%v", err)
	}

	return commentVo.(*comment.CommentsVO), nil
}
