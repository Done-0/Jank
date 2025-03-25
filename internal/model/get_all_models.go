package model

import (
	account "jank.com/jank_blog/internal/model/account"
	category "jank.com/jank_blog/internal/model/category"
	comment "jank.com/jank_blog/internal/model/comment"
	post "jank.com/jank_blog/internal/model/post"
)

// GetAllModels 获取并注册所有模型
func GetAllModels() []interface{} {
	return []interface{}{
		// account 模块
		&account.Account{},

		// post 模块
		&post.Post{},

		// category 模块
		&category.Category{},

		// comment 模块
		&comment.Comment{},
	}
}
