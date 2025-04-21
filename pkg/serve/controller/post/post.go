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
// @Description  根据文章 ID 或标题获取文章的详细信息，至少需要提供其中一个参数
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
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BAD_REQUEST, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BAD_REQUEST), c))
	}

	pos, err := service.GetOnePostByIDOrTitle(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.SERVER_ERR, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(pos, c))
}

// GetAllPosts   godoc
// @Summary      获取文章列表
// @Description  获取所有的文章列表，按创建时间倒序排序
// @Tags         文章
// @Accept       json
// @Produce      json
// @Param        page     query    int     false  "页码"
// @Param        pageSize query    int     false  "每页显示数量"
// @Success      200  {object}  vo.Result{data=[]post.PostsVO}  "获取成功"
// @Failure      500  {object}  vo.Result                 "服务器错误"
// @Router       /post/getAllPosts [get]
func GetAllPosts(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	response, err := service.GetAllPostsWithPagingAndFormat(page, pageSize, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.SERVER_ERR, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(response, c))
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
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BAD_REQUEST, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BAD_REQUEST), c))
	}

	createdPost, err := service.CreateOnePost(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.SERVER_ERR, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(createdPost, c))
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
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BAD_REQUEST, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BAD_REQUEST), c))
	}

	updatedPost, err := service.UpdateOnePost(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.SERVER_ERR, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success(updatedPost, c))
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
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(err, bizErr.New(bizErr.BAD_REQUEST, err.Error()), c))
	}

	errors := utils.Validator(*req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizErr.New(bizErr.BAD_REQUEST), c))
	}

	err := service.DeleteOnePost(req, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(err, bizErr.New(bizErr.SERVER_ERR, err.Error()), c))
	}

	return c.JSON(http.StatusOK, vo.Success("文章删除成功", c))
}
