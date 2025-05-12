// Package service 提供业务逻辑处理，处理评论相关业务
// 创建者：Done-0
// 创建时间：2025-05-10
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
// 参数：
//   - c: Echo 上下文
//   - req: 创建评论请求
//
// 返回值：
//   - *comment.CommentsVO: 创建后的评论视图对象
//   - error: 操作过程中的错误
func CreateComment(c echo.Context, req *dto.CreateCommentRequest) (*comment.CommentsVO, error) {
	var commentVO *comment.CommentsVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		com := &model.Comment{
			Content:          req.Content,
			UserId:           req.UserId,
			PostId:           req.PostId,
			ReplyToCommentId: req.ReplyToCommentId,
		}

		if err := mapper.CreateComment(c, com); err != nil {
			utils.BizLogger(c).Errorf("创建评论失败：%v", err)
			return fmt.Errorf("创建评论失败：%w", err)
		}

		vo, err := utils.MapModelToVO(com, &comment.CommentsVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("创建评论时映射 VO 失败：%v", err)
			return fmt.Errorf("创建评论时映射 VO 失败：%w", err)
		}

		commentVO = vo.(*comment.CommentsVO)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return commentVO, nil
}

// GetCommentWithReplies 根据 ID 获取评论及其所有回复
// 参数：
//   - c: Echo 上下文
//   - req: 获取评论请求
//
// 返回值：
//   - *comment.CommentsVO: 评论及其回复的视图对象
//   - error: 操作过程中的错误
func GetCommentWithReplies(c echo.Context, req *dto.GetOneCommentRequest) (*comment.CommentsVO, error) {
	com, err := mapper.GetCommentByID(c, req.CommentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取评论失败：%v", err)
		return nil, fmt.Errorf("获取评论失败：%w", err)
	}

	replies, err := mapper.GetReplyByCommentID(c, req.CommentID)
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
// 参数：
//   - c: Echo 上下文
//   - req: 获取评论图请求
//
// 返回值：
//   - []*comment.CommentsVO: 评论图结构列表
//   - error: 操作过程中的错误
func GetCommentGraphByPostID(c echo.Context, req *dto.GetCommentGraphRequest) ([]*comment.CommentsVO, error) {
	comments, err := mapper.GetCommentsByPostID(c, req.PostID)
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
// 参数：
//   - c: Echo 上下文
//   - req: 删除评论请求
//
// 返回值：
//   - *comment.CommentsVO: 被删除的评论视图对象
//   - error: 操作过程中的错误
func DeleteComment(c echo.Context, req *dto.DeleteCommentRequest) (*comment.CommentsVO, error) {
	var commentVO *comment.CommentsVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		com, err := mapper.GetCommentByID(c, req.ID)
		if err != nil {
			utils.BizLogger(c).Errorf("获取评论失败：%v", err)
			return fmt.Errorf("评论不存在：%w", err)
		}

		com.Deleted = true
		if err := mapper.UpdateComment(c, com); err != nil {
			utils.BizLogger(c).Errorf("软删除评论失败：%v", err)
			return fmt.Errorf("软删除评论失败：%w", err)
		}

		vo, err := utils.MapModelToVO(com, &comment.CommentsVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("软删除评论时映射 VO 失败：%v", err)
			return fmt.Errorf("软删除评论时映射 VO 失败：%w", err)
		}

		commentVO = vo.(*comment.CommentsVO)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return commentVO, nil
}
