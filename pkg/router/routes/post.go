package routes

import (
	"github.com/labstack/echo/v4"

	auth_middleware "jank.com/jank_blog/internal/middleware/auth"
	"jank.com/jank_blog/pkg/serve/controller/post"
)

func RegisterPostRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	postGroupV1 := apiV1.Group("/post")
	postGroupV1.POST("/getOnePost", post.GetOnePost)
	postGroupV1.GET("/getAllPosts", post.GetAllPosts)
	postGroupV1.POST("/createOnePost", post.CreateOnePost, auth_middleware.AuthMiddleware())
	postGroupV1.POST("/updateOnePost", post.UpdateOnePost, auth_middleware.AuthMiddleware())
	postGroupV1.POST("/deleteOnePost", post.DeleteOnePost, auth_middleware.AuthMiddleware())
}
