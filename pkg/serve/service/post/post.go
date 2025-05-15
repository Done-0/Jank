// Package service 提供业务逻辑处理，处理文章相关业务
// 创建者：Done-0
// 创建时间：2025-05-10
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
// 参数：
//   - c: Echo 上下文
//   - req: 创建文章请求
//
// 返回值：
//   - *post.PostsVO: 创建后的文章视图对象
//   - error: 操作过程中的错误
func CreateOnePost(c echo.Context, req *dto.CreateOnePostRequest) (*post.PostsVO, error) {
	var contentMarkdown string
	var categoryID int64

	contentType := c.Request().Header.Get("Content-Type")
	switch {
	case contentType == "application/json":
		contentMarkdown = req.ContentMarkdown
		categoryID = req.CategoryID
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
		contentMarkdown = string(content)
		categoryIDStr := c.FormValue("category_id")
		if categoryIDStr != "" {
			id, err := strconv.ParseInt(categoryIDStr, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("类目ID格式错误: %v", err)
			}
			categoryID = id
		}
	default:
		return nil, fmt.Errorf("不支持的 Content-Type: %v", contentType)
	}

	if categoryID > 0 {
		_, err := mapper.GetCategoryByID(c, categoryID)
		if err != nil {
			utils.BizLogger(c).Errorf("类目ID「%d」不存在: %v", categoryID, err)
			return nil, fmt.Errorf("类目ID「%d」不存在: %w", categoryID, err)
		}
	}

	contentHTML, err := utils.RenderMarkdown([]byte(contentMarkdown))
	if err != nil {
		utils.BizLogger(c).Errorf("渲染 Markdown 失败: %v", err)
		return nil, fmt.Errorf("渲染 Markdown 失败: %w", err)
	}

	var postsVO *post.PostsVO

	err = utils.RunDBTransaction(c, func(tx error) error {
		newPost := &model.Post{
			Title:           req.Title,
			Image:           req.Image,
			Visibility:      req.Visibility,
			ContentMarkdown: contentMarkdown,
			ContentHTML:     contentHTML,
		}

		if err := mapper.CreatePost(c, newPost); err != nil {
			utils.BizLogger(c).Errorf("创建文章失败: %v", err)
			return fmt.Errorf("创建文章失败: %w", err)
		}

		if err := mapper.CreatePostCategory(c, newPost.ID, categoryID); err != nil {
			utils.BizLogger(c).Errorf("创建文章-类目关联失败: %v", err)
			return fmt.Errorf("创建文章-类目关联失败: %w", err)
		}

		vo, err := utils.MapModelToVO(newPost, &post.PostsVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("创建文章时映射 VO 失败: %v", err)
			return fmt.Errorf("创建文章时映射 VO 失败: %w", err)
		}

		postsVO = vo.(*post.PostsVO)
		postsVO.CategoryID = categoryID

		return nil
	})

	if err != nil {
		return nil, err
	}

	return postsVO, nil
}

// GetOnePostByID 根据 ID 获取文章
// 参数：
//   - c: Echo 上下文
//   - req: 获取文章请求
//
// 返回值：
//   - interface{}: 获取到的文章视图对象
//   - error: 操作过程中的错误
func GetOnePostByID(c echo.Context, req *dto.GetOnePostRequest) (*post.PostsVO, error) {
	pos, err := mapper.GetPostByID(c, req.ID)
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

	postCategory, err := mapper.GetPostCategory(c, pos.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取文章类目关联失败: %v", err)
	}

	if postCategory != nil {
		postsVO.CategoryID = postCategory.CategoryID
	}

	return postsVO, nil
}

// GetAllPostsWithPagingAndFormat 获取格式化后的分页文章列表、总页数和当前页数
// 参数：
//   - c: Echo 上下文
//   - page: 页码
//   - pageSize: 每页文章数量
//
// 返回值：
//   - map[string]interface{}: 包含文章列表、总页数和当前页数的映射
//   - error: 操作过程中的错误
func GetAllPostsWithPagingAndFormat(c echo.Context, page, pageSize int) (map[string]interface{}, error) {
	posts, total, err := mapper.GetAllPostsWithPaging(c, page, pageSize)
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

		postCategory, err := mapper.GetPostCategory(c, pos.ID)
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
// 参数：
//   - c: Echo 上下文
//   - req: 更新文章请求
//
// 返回值：
//   - *post.PostsVO: 更新后的文章视图对象
//   - error: 操作过程中的错误
func UpdateOnePost(c echo.Context, req *dto.UpdateOnePostRequest) (*post.PostsVO, error) {
	var contentMarkdown string
	var categoryID int64

	pos, err := mapper.GetPostByID(c, req.ID)
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
			contentMarkdown = req.ContentMarkdown
			pos.ContentMarkdown = contentMarkdown
			pos.ContentHTML, err = utils.RenderMarkdown([]byte(contentMarkdown))
			if err != nil {
				utils.BizLogger(c).Errorf("渲染 Markdown 失败: %v", err)
				return nil, fmt.Errorf("渲染 Markdown 失败: %w", err)
			}
		}
		categoryID = req.CategoryID

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
			contentMarkdown = string(content)
			pos.ContentMarkdown = contentMarkdown
			pos.ContentHTML, err = utils.RenderMarkdown([]byte(contentMarkdown))
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
			categoryID = id
		}
	default:
		return nil, fmt.Errorf("不支持的 Content-Type: %v", contentType)
	}

	if categoryID > 0 {
		_, err := mapper.GetCategoryByID(c, categoryID)
		if err != nil {
			utils.BizLogger(c).Errorf("类目ID「%d」不存在: %v", categoryID, err)
			return nil, fmt.Errorf("类目ID「%d」不存在: %w", categoryID, err)
		}
	}

	var postsVO *post.PostsVO

	err = utils.RunDBTransaction(c, func(tx error) error {
		if err := mapper.UpdateOnePostByID(c, req.ID, pos); err != nil {
			utils.BizLogger(c).Errorf("更新文章失败: %v", err)
			return fmt.Errorf("更新文章失败: %w", err)
		}

		if err := mapper.UpdatePostCategory(c, req.ID, categoryID); err != nil {
			utils.BizLogger(c).Errorf("更新文章-类目关联失败: %v", err)
			return fmt.Errorf("更新文章-类目关联失败: %w", err)
		}

		vo, err := utils.MapModelToVO(pos, &post.PostsVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("更新文章时映射 VO 失败: %v", err)
			return fmt.Errorf("更新文章时映射 VO 失败: %w", err)
		}

		postsVO = vo.(*post.PostsVO)
		postsVO.CategoryID = categoryID

		return nil
	})

	if err != nil {
		return nil, err
	}

	return postsVO, nil
}

// DeleteOnePost 删除文章
// 参数：
//   - c: Echo 上下文
//   - req: 删除文章请求
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteOnePost(c echo.Context, req *dto.DeleteOnePostRequest) error {
	return utils.RunDBTransaction(c, func(tx error) error {
		if err := mapper.DeleteOnePostByID(c, req.ID); err != nil {
			utils.BizLogger(c).Errorf("删除文章失败: %v", err)
			return fmt.Errorf("删除文章失败: %w", err)
		}

		if err := mapper.DeletePostCategory(c, req.ID); err != nil {
			utils.BizLogger(c).Errorf("删除文章-类目关联失败: %v", err)
			return fmt.Errorf("删除文章-类目关联失败: %w", err)
		}

		return nil
	})
}
