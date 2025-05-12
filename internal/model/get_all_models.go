// Package model 提供应用程序的数据模型定义和聚合
// 创建者：Done-0
// 创建时间：2025-05-10
package model

import (
	account "jank.com/jank_blog/internal/model/account"
	association "jank.com/jank_blog/internal/model/association"
	category "jank.com/jank_blog/internal/model/category"
	comment "jank.com/jank_blog/internal/model/comment"
	post "jank.com/jank_blog/internal/model/post"
)

// GetAllModels 获取并注册所有模型
// 返回值：
//   - []interface{}: 所有需要注册到数据库的模型列表
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

		// association 跨模块中间表
		&association.PostCategory{},
	}
}
