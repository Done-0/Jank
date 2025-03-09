package service

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	model "jank.com/jank_blog/internal/model/post"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/post/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/post"
)

// CreatePost 创建文章
func CreatePost(req *dto.CreateOnePostRequest, c echo.Context) (*post.PostsVo, error) {
	var ContentMarkdown string
	var CategoryIDs []int64

	contentType := c.Request().Header.Get("Content-Type")
	switch {
	case contentType == "application/json":
		ContentMarkdown = req.ContentMarkdown
	case strings.HasPrefix(contentType, "multipart/form-data"):
		file, err := c.FormFile("content_markdown")
		if err != nil {
			return nil, fmt.Errorf("获取上传文件失败: %v", err)
		}
		src, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("打开上传文件失败: %v", err)
		}
		defer func(src multipart.File) {
			err := src.Close()
			if err != nil {
				utils.BizLogger(c).Errorf("关闭上传文件失败: %v", err)
			}
		}(src)
		content, err := io.ReadAll(src)
		if err != nil {
			return nil, fmt.Errorf("读取上传文件内容失败: %v", err)
		}
		ContentMarkdown = string(content)
	default:
		return nil, fmt.Errorf("不支持的 Content-Type: %s", contentType)
	}

	categoryIDsStr := req.CategoryIDs
	if categoryIDsStr == "" {
		categoryIDsStr = c.FormValue("category_ids")
	}
	if categoryIDsStr != "" {
		if err := json.Unmarshal([]byte(categoryIDsStr), &CategoryIDs); err != nil {
			categoryIDsStr = strings.Trim(categoryIDsStr, "[]")
			categoryIDStrs := strings.Split(categoryIDsStr, ",")
			for _, idStr := range categoryIDStrs {
				id, err := strconv.ParseInt(idStr, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("category_ids 格式错误: %w", err)
				}
				CategoryIDs = append(CategoryIDs, id)
			}
		}
	}

	ContentHTML, err := utils.RenderMarkdown([]byte(ContentMarkdown))
	if err != nil {
		return nil, fmt.Errorf("渲染 Markdown 失败: %v", err)
	}

	newPost := &model.Post{
		Title:           req.Title,
		Image:           req.Image,
		Visibility:      req.Visibility,
		ContentMarkdown: ContentMarkdown,
		ContentHTML:     ContentHTML,
		CategoryIDs:     CategoryIDs,
	}

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
	var ContentMarkdown string
	var CategoryIDs []int64

	contentType := c.Request().Header.Get("Content-Type")
	switch {
	case contentType == "application/json":
		ContentMarkdown = req.ContentMarkdown
	case strings.HasPrefix(contentType, "multipart/form-data"):
		if file, err := c.FormFile("content_markdown"); err == nil {
			src, err := file.Open()
			if err != nil {
				return nil, fmt.Errorf("打开上传文件失败: %v", err)
			}
			defer func(src multipart.File) {
				err := src.Close()
				if err != nil {
					utils.BizLogger(c).Errorf("关闭上传文件失败: %v", err)
				}
			}(src)
			content, err := io.ReadAll(src)
			if err != nil {
				return nil, fmt.Errorf("读取上传文件内容失败: %v", err)
			}
			ContentMarkdown = string(content)
		}
	}

	if req.CategoryIDs != "" {
		if err := json.Unmarshal([]byte(req.CategoryIDs), &CategoryIDs); err != nil {
			for _, idStr := range strings.Split(strings.Trim(req.CategoryIDs, "[]"), ",") {
				if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
					CategoryIDs = append(CategoryIDs, id)
				}
			}
		}
	}

	pos, err := mapper.GetPostByID(req.ID)
	if err != nil || pos == nil {
		return nil, fmt.Errorf("获取文章失败: %v", err)
	}

	if req.Title != "" {
		pos.Title = req.Title
	}
	if req.Image != "" {
		pos.Image = req.Image
	}
	if req.Visibility != false {
		pos.Visibility = req.Visibility
	}
	if ContentMarkdown != "" {
		pos.ContentMarkdown = ContentMarkdown
		pos.ContentHTML, err = utils.RenderMarkdown([]byte(ContentMarkdown))
		if err != nil {
			return nil, fmt.Errorf("渲染 Markdown 失败: %v", err)
		}
	}
	if len(CategoryIDs) > 0 {
		pos.CategoryIDs = CategoryIDs
	}

	if err := mapper.UpdateOnePostByID(req.ID, pos); err != nil {
		return nil, fmt.Errorf("更新文章失败: %v", err)
	}

	vo, err := utils.MapModelToVO(pos, &post.PostsVo{})
	if err != nil {
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
