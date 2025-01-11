package routes

import (
	"github.com/labstack/echo/v4"
	"jank.com/jank_blog/pkg/serve/controller/post"

	// "jank.com/jank_blog/middleware/auth"
	render_middleware "jank.com/jank_blog/internal/middleware/render"
)

// func RegisterPostRoutes(app *echo.Echo) {
//     // api v1 group
//     apiV1 := app.Group("/api/v1")
//     postGroup := apiV1.Group("/post")
//     postGroup.POST("/getOnePost", post.GetOnePost)
//     postGroup.POST("/createOnePost", post.CreateOnePost, auth_middleware.JWTMiddleware(), render_middleware.MarkdownRender())
//     postGroup.POST("/updateOnePost", post.UpdateOnePost, auth_middleware.JWTMiddleware(), render_middleware.MarkdownRender())
//     postGroup.POST("/deleteOnePost", post.DeleteOnePost, auth_middleware.JWTMiddleware())
//     postGroup.POST("/deleteOnePost", post.DeleteOnePost)
// }

func RegisterPostRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	postGroupV1 := apiV1.Group("/post")
	postGroupV1.POST("/getOnePost", post.GetOnePost)
	postGroupV1.GET("/getAllPosts", post.GetAllPosts)
	postGroupV1.POST("/createOnePost", post.CreateOnePost, render_middleware.MarkdownRender())
	postGroupV1.POST("/updateOnePost", post.UpdateOnePost, render_middleware.MarkdownRender())
	postGroupV1.POST("/deleteOnePost", post.DeleteOnePost)
}
