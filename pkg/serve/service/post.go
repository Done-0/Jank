package service

import (
	"fmt"
	"math"

	"github.com/labstack/echo/v4"

	model "jank.com/jank_blog/internal/model/post"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/post/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/post"
)

// CreatePost 创建文章页
func CreatePost(req *dto.CreateOnePostRequest, c echo.Context) (*post.PostsVo, error) {
	ContentMarkdown, ok := c.Get("_temp_content_markdown").(string)
	if !ok {
		return nil, fmt.Errorf("获取渲染前的 Markdown 内容失败")
	}

	ContentHTML, ok := c.Get("_temp_content_html").(string)
	if !ok {
		return nil, fmt.Errorf("获取渲染后的 HTML 内容失败")
	}

	// 创建文章对象
	newPost := &model.Post{
		Title:           req.Title,
		Image:           req.Image,
		Visibility:      req.Visibility,
		ContentMarkdown: ContentMarkdown,
		ContentHTML:     ContentHTML,
		CategoryIDs:     req.CategoryIDs,
	}

	defer func() {
		c.Set("_temp_content_markdown", nil)
		c.Set("_temp_content_html", nil)
	}()

	if err := mapper.CreatePost(newPost); err != nil {
		utils.BizLogger(c).Errorf("创建文章失败: %v", err)
		return nil, fmt.Errorf("创建文章失败: %v", err)
	}

	vo, err := utils.MapModelToVO(newPost, &post.PostsVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("创建文章时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("创建文章时映射 vo 失败: %v", err)
	}

	return vo.(*post.PostsVo), nil
}

// GetPostByIDOrTitle 根据 ID 或 Title 获取文章
func GetPostByIDOrTitle(req *dto.GetOnePostRequest, c echo.Context) (interface{}, error) {
	if req.ID == 0 && req.Title == "" {
		utils.BizLogger(c).Error("参数 id 和 title 不能同时为空")
		return nil, fmt.Errorf("参数 id 和 title 不能同时为空")
	}

	// 如果传递了 ID，优先使用 ID 查询
	if req.ID > 0 {
		pos, err := mapper.GetPostByID(req.ID)
		if err != nil {
			utils.BizLogger(c).Errorf("根据 ID 获取文章失败: %v", err)
			return nil, fmt.Errorf("根据 ID 获取文章失败: %v", err)
		}
		if pos == nil {
			utils.BizLogger(c).Errorf("文章不存在: %v", err)
			return nil, fmt.Errorf("文章不存在: %v", err)
		}

		vo, err := utils.MapModelToVO(pos, &post.PostsVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取文章时映射 vo 失败: %v", err)
			return nil, fmt.Errorf("获取文章时映射 vo 失败: %v", err)
		}

		return vo.(*post.PostsVo), nil
	}

	// 如果没有传 ID，使用 Title 查询
	posts, err := mapper.GetPostsByTitle(req.Title)
	if err != nil {
		utils.BizLogger(c).Errorf("根据标题获取文章失败: %v", err)
		return nil, fmt.Errorf("根据标题获取文章失败: %v", err)
	}
	if len(posts) == 0 {
		utils.BizLogger(c).Errorf("没有找到与标题 \"%s\" 匹配的文章", req.Title)
		return nil, fmt.Errorf("没有找到与标题 \"%s\" 匹配的文章", req.Title)
	}

	postResponse := make([]*post.PostsVo, len(posts))
	for i, pos := range posts {
		vo, err := utils.MapModelToVO(pos, &post.PostsVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取文章时映射 vo 失败: %v", err)
			return nil, fmt.Errorf("获取文章时映射 vo 失败: %v", err)
		}

		postResponse[i] = vo.(*post.PostsVo)
	}
	return postResponse, nil
}

// GetAllPostsWithPagingAndFormat 获取格式化后的分页文章列表、总页数和当前页数
func GetAllPostsWithPagingAndFormat(page, pageSize int, c echo.Context) (map[string]interface{}, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 5
	}

	// 获取分页数据和总页数
	posts, total, err := mapper.GetAllPostsWithPaging(page, pageSize)
	if err != nil {
		utils.BizLogger(c).Errorf("获取文章列表失败: %v", err)
		return nil, fmt.Errorf("获取文章列表失败: %v", err)
	}

	postResponse := make([]*post.PostsVo, len(posts))
	for i, pos := range posts {
		vo, err := utils.MapModelToVO(pos, &post.PostsVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取文章列表时映射 vo 失败: %v", err)
			return nil, fmt.Errorf("获取文章列表时映射 vo 失败: %v", err)
		}

		postVo := vo.(*post.PostsVo)

		// 只保留 ContentHTML 的前 150 个字符
		if len(postVo.ContentHTML) > 150 {
			postVo.ContentHTML = postVo.ContentHTML[:150]
		}

		postResponse[i] = postVo
	}

	return map[string]interface{}{
		"posts":       &postResponse,
		"totalPages":  int(math.Ceil(float64(total) / float64(pageSize))),
		"currentPage": page,
	}, nil
}

// UpdatePost 更新文章
func UpdatePost(req *dto.UpdateOnePostRequest, c echo.Context) (*post.PostsVo, error) {
	ContentHTML, ok := c.Get("content_html").(string)
	if !ok {
		return nil, fmt.Errorf("获取文章内容失败: %v", ContentHTML)
	}

	pos, err := mapper.GetPostByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取文章失败: %v", err)
		return nil, fmt.Errorf("获取文章失败: %v", err)
	}
	if pos == nil {
		utils.BizLogger(c).Errorf("文章不存在")
		return nil, fmt.Errorf("文章不存在")
	}

	pos.Title = req.Title
	pos.Image = req.Image
	pos.Visibility = req.Visibility
	pos.ContentMarkdown = req.ContentMarkdown
	pos.ContentHTML = ContentHTML
	pos.CategoryIDs = req.CategoryIDs

	if err := mapper.UpdateOnePostByID(req.ID, pos); err != nil {
		utils.BizLogger(c).Errorf("更新文章失败: %v", err)
		return nil, fmt.Errorf("更新文章失败: %v", err)
	}

	vo, err := utils.MapModelToVO(pos, &post.PostsVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("更新文章时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("更新文章时映射 vo 失败: %v", err)
	}

	return vo.(*post.PostsVo), nil
}

// DeletePost 删除文章
func DeletePost(req *dto.DeleteOnePostRequest, c echo.Context) error {
	pos, err := mapper.GetPostByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取文章失败: %v", err)
		return fmt.Errorf("获取文章失败: %v", err)
	}
	if pos == nil {
		utils.BizLogger(c).Errorf("文章不存在")
		return fmt.Errorf("文章不存在")
	}

	if err := mapper.DeleteOnePostByID(req.ID); err != nil {
		utils.BizLogger(c).Errorf("删除文章失败: %v", err)
		return fmt.Errorf("删除文章失败: %v", err)
	}

	return nil
}
