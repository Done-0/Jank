// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-05-10
package routes

import (
	"github.com/labstack/echo/v4"

	auth_middleware "jank.com/jank_blog/internal/middleware/auth"
	"jank.com/jank_blog/pkg/serve/controller/category"
)

// RegisterCategoryRoutes 注册类目相关路由
// 参数：
//   - r: Echo 路由组数组，r[0] 为 API v1 版本组
func RegisterCategoryRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	categoryGroupV1 := apiV1.Group("/category")
	categoryGroupV1.GET("/getOneCategory", category.GetOneCategory)
	categoryGroupV1.GET("/getCategoryTree", category.GetCategoryTree)
	categoryGroupV1.GET("/getCategoryChildrenTree", category.GetCategoryChildrenTree)
	categoryGroupV1.POST("/createOneCategory", category.CreateOneCategory, auth_middleware.AuthMiddleware())
	categoryGroupV1.POST("/updateOneCategory", category.UpdateOneCategory, auth_middleware.AuthMiddleware())
	categoryGroupV1.POST("/deleteOneCategory", category.DeleteOneCategory, auth_middleware.AuthMiddleware())
}
