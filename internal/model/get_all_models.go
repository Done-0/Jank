package model

import (
	account "jank.com/jank_blog/internal/model/account"
	category "jank.com/jank_blog/internal/model/category"
	post "jank.com/jank_blog/internal/model/post"
)

// GetAllModels 获取并注册所有模型
func GetAllModels() []interface{} {
	return []interface{}{
		&account.Account{},
		&post.Post{},
		&category.Category{},
	}
}
