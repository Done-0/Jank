package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
	model "jank.com/jank_blog/internal/model/post"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/mapper"
)

// CreatePost 处理文章的创建
func CreatePost(title string, image string, visibility string, contentMarkdown string, contentHTML string, categoryIDs []int64) (*model.Post, error) {
    if visibility == "" {
        visibility = "private"
    }

    categoryIDsStr := utils.ConvertInt64SliceToString(categoryIDs)

    newPost := &model.Post{
        Title:           title,
        Image:           image,
        Visibility:      visibility,
        ContentMarkdown: contentMarkdown,
        ContentHTML:     contentHTML,
        CategoryIDs:     categoryIDsStr,
    }

    if err := mapper.CreatePost(newPost); err != nil {
        return nil, fmt.Errorf("创建文章失败: %w", err)
    }

    return newPost, nil
}

// GetOnePostByIDOrTitle 根据 ID 或 Title 获取文章
func GetPostByIDOrTitle(id int64, title string, c echo.Context) (interface{}, error) {
	if id <= 0 && title == "" {
		utils.BizLogger(c).Error("参数 id 和 title 不能同时为空")
		return nil, fmt.Errorf("参数 id 和 title 不能同时为空")
	}

	if id > 0 {
		post, err := mapper.GetPostByID(id)
		if err != nil {
			return nil, fmt.Errorf("根据 ID 获取文章失败: %w", err)
		}
		if post == nil {
			return nil, fmt.Errorf("文章不存在")
		}
		return post, nil
	}

	posts, err := mapper.GetPostsByTitle(title)
	if err != nil {
		return nil, fmt.Errorf("根据标题获取文章失败: %w", err)
	}
	if len(posts) == 0 {
		return nil, fmt.Errorf("没有找到与标题 \"%s\" 匹配的文章", title)
	}
	return posts, nil
}

// GetAllPostsWithPaging 获取分页后的文章列表和总页数
func GetAllPostsWithPaging(page, pageSize int) ([]*model.Post, int, error) {
	posts, total, err := mapper.GetAllPostsWithPaging(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalPages := total / int64(pageSize)
	if total%int64(pageSize) > 0 {
		totalPages++
	}

	return posts, int(totalPages), nil
}

// UpdateOnePost 更新文章
func UpdatePost(id int64, title string, image string, visibility string, contentMarkdown string, contentHTML string, categoryIDs []int64, c echo.Context) (*model.Post, error) {
    post, err := mapper.GetPostByID(id)
    if err != nil {
        return nil, fmt.Errorf("获取文章失败: %w", err)
    }
    if post == nil {
        return nil, fmt.Errorf("文章不存在")
    }

    categoryIDsStr := utils.ConvertInt64SliceToString(categoryIDs)

    post.Title = title
    post.Image = image
    post.Visibility = visibility
    post.ContentMarkdown = contentMarkdown
    post.ContentHTML = contentHTML
    post.CategoryIDs = categoryIDsStr

    if err := mapper.UpdateOnePostByID(id, post); err != nil {
        return nil, fmt.Errorf("更新文章失败: %w", err)
    }

    return post, nil
}

// DeleteOnePostByID 删除文章
func DeletePost(id int64, c echo.Context) error {
	post, err := mapper.GetPostByID(id)
	if err != nil {
		return fmt.Errorf("获取文章失败: %w", err)
	}
	if post == nil {
		utils.BizLogger(c).Errorf("文章不存在")
		return fmt.Errorf("文章不存在")
	}

	if err := mapper.DeleteOnePostByID(id); err != nil {
		return fmt.Errorf("删除文章失败: %w", err)
	}

	return nil
}
