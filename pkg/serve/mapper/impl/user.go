// Package impl 提供用户相关的数据访问实现
// 创建者：Done-0
// 创建时间：2025-08-10
package impl

import (
	"strings"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/model/user"
	"github.com/Done-0/jank/internal/utils/db"
	"github.com/Done-0/jank/pkg/serve/mapper"
)

// UserMapperImpl 用户数据访问实现
type UserMapperImpl struct{}

// NewUserMapper 创建用户数据访问实例
func NewUserMapper() mapper.UserMapper {
	return &UserMapperImpl{}
}

// GetUserByEmail 根据邮箱获取用户
func (m *UserMapperImpl) GetUserByEmail(c *app.RequestContext, email string) (*user.User, error) {
	var u user.User
	err := db.GetDBFromContext(c).Where("email = ? AND deleted = ?", email, false).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetUserByID 根据ID获取用户
func (m *UserMapperImpl) GetUserByID(c *app.RequestContext, userID int64) (*user.User, error) {
	var u user.User
	err := db.GetDBFromContext(c).Where("id = ? AND deleted = ?", userID, false).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetUserByNickname 根据昵称获取用户
func (m *UserMapperImpl) GetUserByNickname(c *app.RequestContext, nickname string) (*user.User, error) {
	var u user.User
	err := db.GetDBFromContext(c).Where("nickname = ? AND deleted = ?", nickname, false).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// RegisterUser 注册用户（包含重复检查和事务处理）
func (m *UserMapperImpl) RegisterUser(c *app.RequestContext, u *user.User) error {
	_, err := db.RunDBTransaction(c, func() (any, error) {
		// 检查邮箱是否已被注册
		var existingUser user.User
		if err := db.GetDBFromContext(c).Where("email = ? AND deleted = ?", u.Email, false).First(&existingUser).Error; err == nil {
			return nil, err
		}

		// 检查昵称是否已被使用
		if err := db.GetDBFromContext(c).Where("nickname = ? AND deleted = ?", u.Nickname, false).First(&existingUser).Error; err == nil {
			return nil, err
		}

		// 创建用户
		if err := db.GetDBFromContext(c).Create(u).Error; err != nil {
			return nil, err
		}
		return nil, nil
	})
	return err
}

// UpdateUser 更新用户信息
func (m *UserMapperImpl) UpdateUser(c *app.RequestContext, u *user.User) error {
	if err := db.GetDBFromContext(c).Where("id = ? AND deleted = ?", u.ID, false).Updates(u).Error; err != nil {
		return err
	}
	return nil
}

// ListUsers 获取用户列表
func (m *UserMapperImpl) ListUsers(c *app.RequestContext, pageNo, pageSize int64, keyword, role string) ([]*user.User, int64, error) {
	var users []*user.User
	var total int64

	query := db.GetDBFromContext(c).Model(&user.User{}).Where("deleted = ?", false)

	// 关键词搜索
	if keyword != "" {
		keyword = "%" + strings.TrimSpace(keyword) + "%"
		query = query.Where("email LIKE ? OR nickname LIKE ?", keyword, keyword)
	}

	// 角色筛选
	if role != "" {
		query = query.Where("role = ?", role)
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询记录 - 分页查询
	offset := (pageNo - 1) * pageSize
	if err := query.
		Order("id DESC").
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// DeleteUser 删除用户
func (m *UserMapperImpl) DeleteUser(c *app.RequestContext, userID string) error {
	if err := db.GetDBFromContext(c).Where("id = ?", userID).Delete(&user.User{}).Error; err != nil {
		return err
	}
	return nil
}
