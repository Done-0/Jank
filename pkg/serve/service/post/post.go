package service

import (
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

// CreateOnePost 创建文章
func CreateOnePost(req *dto.CreateOnePostRequest, c echo.Context) (*post.PostsVO, error) {
	var ContentMarkdown string
	var CategoryID int64

	contentType := c.Request().Header.Get("Content-Type")
	switch {
	case contentType == "application/json":
		ContentMarkdown = req.ContentMarkdown
		CategoryID = req.CategoryID
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
		categoryIDStr := c.FormValue("category_id")
		if categoryIDStr != "" {
			id, err := strconv.ParseInt(categoryIDStr, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("类目ID格式错误: %v", err)
			}
			CategoryID = id
		}
	default:
		return nil, fmt.Errorf("不支持的 Content-Type: %v", contentType)
	}

	if CategoryID > 0 {
		_, err := mapper.GetCategoryByID(CategoryID)
		if err != nil {
			utils.BizLogger(c).Errorf("类目ID「%d」不存在: %v", CategoryID, err)
			return nil, fmt.Errorf("类目ID「%d」不存在: %w", CategoryID, err)
		}
	}

	ContentHTML, err := utils.RenderMarkdown([]byte(ContentMarkdown))
	if err != nil {
		utils.BizLogger(c).Errorf("渲染 Markdown 失败: %v", err)
		return nil, fmt.Errorf("渲染 Markdown 失败: %w", err)
	}

	newPost := &model.Post{
		Title:           req.Title,
		Image:           req.Image,
		Visibility:      req.Visibility,
		ContentMarkdown: ContentMarkdown,
		ContentHTML:     ContentHTML,
	}

	if err := mapper.CreatePost(newPost); err != nil {
		utils.BizLogger(c).Errorf("创建文章失败: %v", err)
		return nil, fmt.Errorf("创建文章失败: %w", err)
	}

	if CategoryID > 0 {
		if err := mapper.CreatePostCategory(newPost.ID, CategoryID); err != nil {
			utils.BizLogger(c).Errorf("创建文章-类目关联失败: %v", err)
			return nil, fmt.Errorf("创建文章-类目关联失败: %w", err)
		}
	}

	vo, err := utils.MapModelToVO(newPost, &post.PostsVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("创建文章时映射 VO 失败: %v", err)
		return nil, fmt.Errorf("创建文章时映射 VO 失败: %w", err)
	}

	postsVO := vo.(*post.PostsVO)
	postsVO.CategoryID = CategoryID

	return postsVO, nil
}

// GetOnePostByIDOrTitle 根据 ID 或 Title 获取文章
func GetOnePostByIDOrTitle(req *dto.GetOnePostRequest, c echo.Context) (interface{}, error) {
	switch {
	case req.ID > 0:
		pos, err := mapper.GetPostByID(req.ID)
		if err != nil {
			utils.BizLogger(c).Errorf("根据 ID 获取文章失败: %v", err)
			return nil, fmt.Errorf("根据 ID 获取文章失败: %w", err)
		}
		if pos == nil {
			utils.BizLogger(c).Errorf("文章不存在: %v", err)
			return nil, fmt.Errorf("文章不存在: %w", err)
		}

		vo, err := utils.MapModelToVO(pos, &post.PostsVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取文章时映射 VO 失败: %v", err)
			return nil, fmt.Errorf("获取文章时映射 VO 失败: %w", err)
		}

		postsVO := vo.(*post.PostsVO)

		postCategory, err := mapper.GetPostCategory(pos.ID)
		if err != nil {
			utils.BizLogger(c).Errorf("获取文章类目关联失败: %v", err)
		}

		if postCategory != nil {
			postsVO.CategoryID = postCategory.CategoryID
		}

		return postsVO, nil

	case req.Title != "":
		posts, err := mapper.GetPostsByTitle(req.Title)
		if err != nil {
			utils.BizLogger(c).Errorf("根据标题获取文章失败: %v", err)
			return nil, fmt.Errorf("根据标题获取文章失败: %w", err)
		}
		if len(posts) == 0 {
			utils.BizLogger(c).Errorf("没有找到与标题 \"%s\" 匹配的文章", req.Title)
			return nil, fmt.Errorf("没有找到与标题 \"%s\" 匹配的文章", req.Title)
		}

		postsVO := make([]*post.PostsVO, len(posts))
		for i, pos := range posts {
			vo, err := utils.MapModelToVO(pos, &post.PostsVO{})
			if err != nil {
				utils.BizLogger(c).Errorf("获取文章时映射 VO 失败: %v", err)
				return nil, fmt.Errorf("获取文章时映射 VO 失败: %w", err)
			}

			postVO := vo.(*post.PostsVO)

			postCategory, err := mapper.GetPostCategory(pos.ID)
			if err != nil {
				utils.BizLogger(c).Errorf("获取文章ID「%d」的类目关联失败: %v", pos.ID, err)
			}

			if postCategory != nil {
				postVO.CategoryID = postCategory.CategoryID
			}

			postsVO[i] = postVO
		}
		return postsVO, nil

	default:
		utils.BizLogger(c).Error("参数 id 和 title 至少需要传递一个")
		return nil, fmt.Errorf("参数 id 和 title 至少需要传递一个")
	}
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
		return nil, fmt.Errorf("获取文章列表失败: %w", err)
	}

	postResponse := make([]*post.PostsVO, len(posts))
	for i, pos := range posts {
		vo, err := utils.MapModelToVO(pos, &post.PostsVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取文章列表时映射 VO 失败: %v", err)
			return nil, fmt.Errorf("获取文章列表时映射 VO 失败: %w", err)
		}

		postVO := vo.(*post.PostsVO)

		postCategory, err := mapper.GetPostCategory(pos.ID)
		if err != nil {
			utils.BizLogger(c).Errorf("获取文章ID「%d」的类目关联失败: %v", pos.ID, err)
		}

		if postCategory != nil {
			postVO.CategoryID = postCategory.CategoryID
		}

		// 只保留 ContentHTML 的前 200 个字符
		if len(postVO.ContentHTML) > 200 {
			postVO.ContentHTML = postVO.ContentHTML[:200]
		}

		postResponse[i] = postVO
	}

	return map[string]interface{}{
		"posts":       &postResponse,
		"totalPages":  int(math.Ceil(float64(total) / float64(pageSize))),
		"currentPage": page,
	}, nil
}

// UpdateOnePost 更新文章
func UpdateOnePost(req *dto.UpdateOnePostRequest, c echo.Context) (*post.PostsVO, error) {
	var ContentMarkdown string
	var CategoryID int64

	pos, err := mapper.GetPostByID(req.ID)
	if err != nil || pos == nil {
		utils.BizLogger(c).Errorf("获取文章失败: %v", err)
		return nil, fmt.Errorf("获取文章失败: %w", err)
	}

	contentType := c.Request().Header.Get("Content-Type")
	switch {
	case contentType == "application/json":
		if req.Title != "" {
			pos.Title = req.Title
		}
		if req.Image != "" {
			pos.Image = req.Image
		}
		pos.Visibility = req.Visibility
		if req.ContentMarkdown != "" {
			ContentMarkdown = req.ContentMarkdown
			pos.ContentMarkdown = ContentMarkdown
			pos.ContentHTML, err = utils.RenderMarkdown([]byte(ContentMarkdown))
			if err != nil {
				utils.BizLogger(c).Errorf("渲染 Markdown 失败: %v", err)
				return nil, fmt.Errorf("渲染 Markdown 失败: %w", err)
			}
		}
		CategoryID = req.CategoryID

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
			pos.ContentMarkdown = ContentMarkdown
			pos.ContentHTML, err = utils.RenderMarkdown([]byte(ContentMarkdown))
			if err != nil {
				utils.BizLogger(c).Errorf("渲染 Markdown 失败: %v", err)
				return nil, fmt.Errorf("渲染 Markdown 失败: %w", err)
			}
		}

		if title := c.FormValue("title"); title != "" {
			pos.Title = title
		}
		if image := c.FormValue("image"); image != "" {
			pos.Image = image
		}
		if visibility := c.FormValue("visibility"); visibility != "" {
			pos.Visibility = visibility == "true"
		}
		if categoryIDStr := c.FormValue("category_id"); categoryIDStr != "" {
			id, err := strconv.ParseInt(categoryIDStr, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("category_id 格式错误: %w", err)
			}
			CategoryID = id
		}
	}

	if CategoryID > 0 {
		if _, err := mapper.GetCategoryByID(CategoryID); err != nil {
			utils.BizLogger(c).Errorf("类目ID「%d」不存在: %v", CategoryID, err)
			return nil, fmt.Errorf("类目ID「%d」不存在: %w", CategoryID, err)
		}
	}

	if err := mapper.UpdateOnePostByID(req.ID, pos); err != nil {
		utils.BizLogger(c).Errorf("更新文章失败: %v", err)
		return nil, fmt.Errorf("更新文章失败: %w", err)
	}

	if CategoryID > 0 {
		if err := mapper.UpdatePostCategory(req.ID, CategoryID); err != nil {
			utils.BizLogger(c).Errorf("更新文章-类目关联失败: %v", err)
			return nil, fmt.Errorf("更新文章-类目关联失败: %w", err)
		}
	}

	vo, err := utils.MapModelToVO(pos, &post.PostsVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("更新文章时映射 VO 失败: %v", err)
		return nil, fmt.Errorf("更新文章时映射 VO 失败: %w", err)
	}

	postsVO := vo.(*post.PostsVO)
	postsVO.CategoryID = CategoryID

	return postsVO, nil
}

// DeleteOnePost 删除文章
func DeleteOnePost(req *dto.DeleteOnePostRequest, c echo.Context) error {
	pos, err := mapper.GetPostByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取文章失败: %v", err)
		return fmt.Errorf("获取文章失败: %w", err)
	}
	if pos == nil {
		utils.BizLogger(c).Errorf("文章不存在")
		return fmt.Errorf("文章不存在")
	}

	if err := mapper.DeleteOnePostByID(req.ID); err != nil {
		utils.BizLogger(c).Errorf("删除文章失败: %v", err)
		return fmt.Errorf("删除文章失败: %w", err)
	}

	if err := mapper.DeletePostCategory(req.ID); err != nil {
		utils.BizLogger(c).Errorf("删除文章-类目关联失败: %v", err)
		return fmt.Errorf("删除文章-类目关联失败: %w", err)
	}

	return nil
}
