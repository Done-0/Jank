package routes

import (
	"github.com/labstack/echo/v4"
	"jank.com/jank_blog/pkg/serve/controller/category"
)

func RegisterCategoryRoutes(app *echo.Echo) {
	categoryGroup := app.Group("/category")
	categoryGroup.GET("/getOneCategory", category.GetOneCategory)
	categoryGroup.GET("/getCategoryTree", category.GetCategoryTree)
	categoryGroup.POST("/getCategoryChildrenTree", category.GetCategoryChildrenTree)
	categoryGroup.POST("/createOneCategory", category.CreateOneCategory)
	categoryGroup.POST("/updateOneCategory", category.UpdateOneCategory)
	categoryGroup.POST("/deleteOneCategory", category.DeleteOneCategory)
}
