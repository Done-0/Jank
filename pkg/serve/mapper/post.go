package mapper

import (
	"fmt"

	"jank.com/jank_blog/internal/global"
	model "jank.com/jank_blog/internal/model/post"
)

const (
	activePostCondition = "`deleted` = 0"
)

// CreatePost 将文章保存到数据库
func CreatePost(newPost *model.Post) error {
	if newPost == nil {
		return fmt.Errorf("文章不能为空")
	}

	return global.DB.Create(newPost).Error
}

// GetPostByID 根据 ID 获取文章
func GetPostByID(id int64) (*model.Post, error) {
	if id <= 0 {
		return nil, fmt.Errorf("无效文章ID: %d", id)
	}

	var post model.Post
	err := global.DB.Where("id = ? AND "+activePostCondition, id).
		First(&post).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

// GetPostsByTitle 通过 Title 获取所有匹配的文章
func GetPostsByTitle(title string) ([]model.Post, error) {
	if title == "" {
		return nil, fmt.Errorf("文章标题不能为空")
	}

	var posts []model.Post
	err := global.DB.Where("title LIKE ? AND "+activePostCondition, "%"+title+"%").
		Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

// GetAllPostsWithPaging 获取分页后的文章列表和文章总数
func GetAllPostsWithPaging(page, pageSize int) ([]*model.Post, int64, error) {
	var posts []*model.Post
	var total int64

	offset := (page - 1) * pageSize

	// 查询文章总数
	err := global.DB.Model(&model.Post{}).
		Where("deleted = ?", false).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询分页数据
	err = global.DB.Model(&model.Post{}).
		Where("deleted = ?", false).
		Order("gmt_create DESC").
		Offset(offset).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// UpdatePost 更新文章
func UpdateOnePostByID(id int64, newPost *model.Post) error {
	if id <= 0 || newPost == nil {
		return fmt.Errorf("无效文章ID: %d", id)
	}

	result := global.DB.Model(&model.Post{}).
		Where("id = ? AND "+activePostCondition, id).
		Updates(newPost)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("文章不存在或已经删除")
	}

	return nil
}

// DeletePostByID 根据 ID 进行软删除操作
func DeleteOnePostByID(id int64) error {
	if id <= 0 {
		return fmt.Errorf("无效文章ID: %d", id)
	}

	result := global.DB.Model(&model.Post{}).
		Where("id = ? AND "+activePostCondition, id).
		Update("deleted", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("文章不存在或已经删除")
	}

	return nil
}
