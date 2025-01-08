package post

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	bizerr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/post/dto"
	"jank.com/jank_blog/pkg/serve/service"
	"jank.com/jank_blog/pkg/vo"
	_ "jank.com/jank_blog/pkg/vo/post"
)

// GetOnePost godoc
// @Summary      获取文章详情
// @Description  根据文章 ID 或标题获取文章的详细信息，至少需要提供其中一个参数
// @Tags         文章
// @Accept       json
// @Produce      json
// @Param        id       query     int     false  "文章 ID"
// @Param        title    query     string  false  "文章标题"
// @Success      200      {object}  vo.Result{data=post.GetAllPostsVo}  "获取成功"
// @Failure      400      {object}  vo.Result          "请求参数错误"
// @Failure      404      {object}  vo.Result          "文章不存在"
// @Failure      500      {object}  vo.Result          "服务器错误"
// @Router       /post/getOnePost [get]
func GetOnePost(c echo.Context) error {
	req := new(dto.GetOnePostRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(bizerr.New(bizerr.BadRequest, err.Error()), nil, c))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizerr.New(bizerr.BadRequest), c))
	}

	if req.ID == 0 && req.Title == "" {
		return c.JSON(http.StatusBadRequest, vo.Fail(bizerr.New(bizerr.BadRequest, "文章 ID 或标题不能为空"), nil, c))
	}

	post, err := service.GetPostByIDOrTitle(req.ID, req.Title, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(bizerr.New(bizerr.UnKnowErr, err.Error()), nil, c))
	}

	return c.JSON(http.StatusOK, vo.Success(post, c))
}

// GetAllPost godoc
// @Summary      获取文章列表
// @Description  获取所有的文章列表，按创建时间倒序排序
// @Tags         文章
// @Accept       json
// @Produce      json
// @Param        page     query    int     false  "页数"
// @Param        pageSize query    int     false  "每页显示数量"
// @Success      200  {object}  vo.Result{data=[]post.GetAllPostsVo}  "获取成功"
// @Failure      500  {object}  vo.Result                 "服务器错误"
// @Router       /post/getAllPost [get]
func GetAllPosts(c echo.Context) error {
    page, _ := strconv.Atoi(c.QueryParam("page"))
    pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

    response, err := service.GetAllPostsWithPagingAndFormat(page, pageSize, c)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, vo.Fail(bizerr.New(bizerr.UnKnowErr, err.Error()), nil, c))
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
// @Success      200     {object}   vo.Result{data=post.GetAllPostsVo}  "创建成功"
// @Failure      400     {object}   vo.Result          "请求参数错误"
// @Failure      500     {object}   vo.Result          "服务器错误"
// @Security     BearerAuth
// @Router       /post/createOnePost [post]
func CreateOnePost(c echo.Context) error {
	req := new(dto.CreateOnePostRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(bizerr.New(bizerr.BadRequest, err.Error()), nil, c))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizerr.New(bizerr.BadRequest), c))
	}

	ContentHTML, ok := c.Get("contentHtml").(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, vo.Fail(bizerr.New(bizerr.UnKnowErr, "渲染失败，缺少 contentHtml"), nil, c))
	}

	createdPost, err := service.CreatePost(req.Title, req.Image, req.Visibility, req.ContentMarkdown, ContentHTML, req.CategoryIDs, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(bizerr.New(bizerr.UnKnowErr, err.Error()), nil, c))
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
// @Success      200     {object}   vo.Result{data=post.GetAllPostsVo}  "更新成功"
// @Failure      400     {object}   vo.Result          "请求参数错误"
// @Failure      404     {object}   vo.Result          "文章不存在"
// @Failure      500     {object}   vo.Result          "服务器错误"
// @Security     BearerAuth
// @Router       /post/updateOnePost [post]
func UpdateOnePost(c echo.Context) error {
	req := new(dto.UpdateOnePostRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(bizerr.New(bizerr.BadRequest, err.Error()), nil, c))
	}

	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizerr.New(bizerr.BadRequest), c))
	}

	contentHTML := c.Get("contentHtml").(string)

	updatedPost, err := service.UpdatePost(req.ID, req.Title, req.Image, req.Visibility, req.ContentMarkdown, contentHTML, req.CategoryIDs, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(bizerr.New(bizerr.UnKnowErr, err.Error()), nil, c))
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
		return c.JSON(http.StatusBadRequest, vo.Fail(bizerr.New(bizerr.BadRequest, err.Error()), nil, c))
	}
	errors := utils.Validator(req)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, vo.Fail(errors, bizerr.New(bizerr.BadRequest), c))
	}

	err := service.DeletePost(req.ID, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, vo.Fail(bizerr.New(bizerr.UnKnowErr, err.Error()), nil, c))
	}

	return c.JSON(http.StatusOK, vo.Success("文章删除成功", c))
}
