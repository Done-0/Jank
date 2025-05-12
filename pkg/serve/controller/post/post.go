// Package post 提供文章相关的HTTP接口处理
// 创建者：Done-0
// 创建时间：2025-05-10
package post

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/post/dto"
	service "jank.com/jank_blog/pkg/serve/service/post"
	"jank.com/jank_blog/pkg/vo"
)

// GetOnePost    godoc
// @Summary      获取文章详情
// @Description  根据文章 ID 获取文章的详细信息
// @Tags         文章
// @Accept       json
// @Produce      json
// @Param        request  body      dto.GetOnePostRequest  true  "获取文章请求参数"
// @Success      200      {object}  vo.Result{data=post.PostsVO}  "获取成功"
// @Failure      400      {object}  vo.Result          "请求参数错误"
// @Failure      404      {object}  vo.Result          "文章不存在"
// @Failure      500      {object}  vo.Result          "服务器错误"
// @Router       /post/getOnePost [get]
func GetOnePost(c echo.Context) error {
	req := new(dto.GetOnePostRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	pos, err := service.GetOnePostByID(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, pos))
}

// GetAllPosts   godoc
// @Summary      获取文章列表
// @Description  获取所有的文章列表，按创建时间倒序排序
// @Tags         文章
// @Accept       json
// @Produce      json
// @Param        page        query     int     true   "页码"
// @Param        page_size   query     int     true   "每页条数"
// @Success      200  {object}  vo.Result{data=[]post.PostsVO}  "获取成功"
// @Failure      500  {object}  vo.Result                 "服务器错误"
// @Router       /post/getAllPosts [get]
func GetAllPosts(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 5
	}

	posts, err := service.GetAllPostsWithPagingAndFormat(c, page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, posts))
}

// CreateOnePost godoc
// @Summary      创建文章
// @Description  创建新的文章，支持 Markdown 格式内容，系统会自动转换为 HTML
// @Tags         文章
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateOnePostRequest  true  "创建文章请求参数"
// @Success      200     {object}   vo.Result{data=post.PostsVO}  "创建成功"
// @Failure      400     {object}   vo.Result          "请求参数错误"
// @Failure      500     {object}   vo.Result          "服务器错误"
// @Security     BearerAuth
// @Router       /post/createOnePost [post]
func CreateOnePost(c echo.Context) error {
	req := new(dto.CreateOnePostRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	createdPost, err := service.CreateOnePost(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, createdPost))
}

// UpdateOnePost godoc
// @Summary      更新文章
// @Description  更新已存在的文章内容
// @Tags         文章
// @Accept       json
// @Produce      json
// @Param        request  body      dto.UpdateOnePostRequest  true  "更新文章请求参数"
// @Success      200     {object}   vo.Result{data=post.PostsVO}  "更新成功"
// @Failure      400     {object}   vo.Result          "请求参数错误"
// @Failure      404     {object}   vo.Result          "文章不存在"
// @Failure      500     {object}   vo.Result          "服务器错误"
// @Security     BearerAuth
// @Router       /post/updateOnePost [post]
func UpdateOnePost(c echo.Context) error {
	req := new(dto.UpdateOnePostRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	updatedPost, err := service.UpdateOnePost(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, updatedPost))
}

// DeleteOnePost godoc
// @Summary      删除文章
// @Description  根据文章 ID 删除指定文章
// @Tags         文章
// @Accept       json
// @Produce      json
// @Param        request  body      dto.DeleteOnePostRequest  true  "删除文章请求参数"
// @Success      200     {object}   vo.Result          "删除成功"
// @Failure      400     {object}   vo.Result          "请求参数错误"
// @Failure      404     {object}   vo.Result          "文章不存在"
// @Failure      500     {object}   vo.Result          "服务器错误"
// @Security     BearerAuth
// @Router       /post/deleteOnePost [post]
func DeleteOnePost(c echo.Context) error {
	req := new(dto.DeleteOnePostRequest)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, err, bizErr.New(bizErr.BAD_REQUEST, err.Error())))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, errors, bizErr.New(bizErr.BAD_REQUEST)))
	}

	err := service.DeleteOnePost(c, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "文章删除成功"))
}
