// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-08-05
package routes

import (
	"log"

	"github.com/cloudwego/hertz/pkg/route"

	"github.com/Done-0/jank/internal/middleware/jwt"
	"github.com/Done-0/jank/pkg/wire"
)

// RegisterPostRoutes 注册文章相关路由
func RegisterPostRoutes(r *route.RouterGroup) {
	postController, err := wire.NewPostController()
	if err != nil {
		log.Fatalf("Failed to initialize post controller: %v", err)
	}

	// 文章路由组
	postGroup := r.Group("/post")
	{
		postGroup.GET("/get", postController.GetPost)                                 // 获取单篇文章
		postGroup.GET("/list-published", postController.ListPublishedPosts)           // 获取已发布文章列表
		postGroup.GET("/list-by-status", jwt.New(), postController.ListPostsByStatus) // 根据状态获取文章列表（支持管理员查询所有文章）
		postGroup.POST("/create", jwt.New(), postController.Create)                   // 创建文章
		postGroup.POST("/update", jwt.New(), postController.Update)                   // 更新文章
		postGroup.POST("/delete", jwt.New(), postController.Delete)                   // 删除文章
	}
}
