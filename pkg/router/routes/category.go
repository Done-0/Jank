package routes

import (
	"github.com/labstack/echo/v4"

	authMiddleware "jank.com/jank_blog/internal/middleware/auth"
	"jank.com/jank_blog/pkg/serve/controller/category"
)

func RegisterCategoryRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	categoryGroupV1 := apiV1.Group("/category")
	categoryGroupV1.GET("/getOneCategory", category.GetOneCategory)
	categoryGroupV1.GET("/getCategoryTree", category.GetCategoryTree)
	categoryGroupV1.GET("/getCategoryChildrenTree", category.GetCategoryChildrenTree)
	categoryGroupV1.POST("/createOneCategory", category.CreateOneCategory, authMiddleware.AuthMiddleware())
	categoryGroupV1.POST("/updateOneCategory", category.UpdateOneCategory, authMiddleware.AuthMiddleware())
	categoryGroupV1.POST("/deleteOneCategory", category.DeleteOneCategory, authMiddleware.AuthMiddleware())
}
