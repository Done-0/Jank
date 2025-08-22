// Package model 提供应用程序的数据模型定义和聚合
// 创建者：Done-0
// 创建时间：2025-08-05
package model

import (
	"github.com/Done-0/jank/internal/model/category"
	"github.com/Done-0/jank/internal/model/post"
	"github.com/Done-0/jank/internal/model/rbac"
	"github.com/Done-0/jank/internal/model/user"
)

// GetAllModels 获取并注册所有模型
// 返回值：
//
//	[]any: 所有模型列表
func GetAllModels() []any {
	return []any{
		&user.User{},         // 用户模型
		&rbac.Policy{},       // RBAC策略模型
		&post.Post{},         // 文章模型
		&category.Category{}, // 分类模型
	}
}
