// Package comment 提供评论相关的HTTP接口处理
// 创建者：Done-0
// 创建时间：2025-05-10
package comment

import (
	"net/http"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/comment/dto"
	service "jank.com/jank_blog/pkg/serve/service/comment"
	"jank.com/jank_blog/pkg/vo"
)

// GetOneComment godoc
// @Summary      获取评论详情
// @Description  根据评论 ID 获取单个评论以及子评论
// @Tags         评论
// @Accept       json
// @Produce      json
// @Param        id    query     int  true  "评论ID"
// @Success      200   {object}  vo.Result{data=comment.CommentsVO}  "获取成功"
// @Failure      400   {object}  vo.Result  "请求参数错误"
// @Failure      404   {object}  vo.Result  "评论不存在"
// @Router       /comment/getOneComment [get]
func GetOneComment(c echo.Context) error {
	req := new(dto.GetOneCommentRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	comment, err := service.GetCommentWithReplies(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, comment))
}

// GetCommentGraph godoc
// @Summary      获取评论图
// @Description  根据文章 ID 获取评论图结构
// @Tags         评论
// @Accept       json
// @Produce      json
// @Param        post_id    query     int  true  "文章ID"
// @Success      200        {object} vo.Result{data=[]comment.CommentsVO}  "获取成功"
// @Failure      500        {object} vo.Result  "服务器错误"
// @Router       /comment/getOneComment [get]
func GetCommentGraph(c echo.Context) error {
	req := new(dto.GetCommentGraphRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	comments, err := service.GetCommentGraphByPostID(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, comments))
}

// CreateOneComment godoc
// @Summary      创建评论
// @Description  创建一条新的评论
// @Tags         评论
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateCommentRequest  true  "创建评论请求参数"
// @Success      200     {object}   vo.Result{data=comment.CommentsVO}  "创建成功"
// @Failure      400     {object}   vo.Result          "请求参数错误"
// @Router       /comment/createOneComment [post]
func CreateOneComment(c echo.Context) error {
	req := new(dto.CreateCommentRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	comment, err := service.CreateComment(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, comment))
}

// DeleteOneComment godoc
// @Summary      软删除评论
// @Description  通过评论 ID 进行软删除
// @Tags         评论
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "评论ID"
// @Success      200   {object} vo.Result{data=comment.CommentsVO}  "软删除成功"
// @Failure      400   {object} vo.Result  "请求参数错误"
// @Failure      404   {object} vo.Result  "评论不存在"
// @Router       /comment/deleteOneComment [post]
func DeleteOneComment(c echo.Context) error {
	req := new(dto.DeleteCommentRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	comment, err := service.DeleteComment(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, comment))
}
