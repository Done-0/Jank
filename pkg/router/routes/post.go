package routes

import (
	"github.com/labstack/echo/v4"
	"jank.com/jank_blog/pkg/serve/controller/post"

	// "jank.com/jank_blog/middleware/auth"
	render_middleware "jank.com/jank_blog/internal/middleware/render"
)

// func RegisterPostRoutes(app *echo.Echo) {
//     postGroup := app.Group("/post")
//     postGroup.POST("/getOnePost", post.GetOnePost)
//     postGroup.POST("/createOnePost", post.CreateOnePost, auth_middleware.JWTMiddleware(), render_middleware.MarkdownRender())
//     postGroup.POST("/updateOnePost", post.UpdateOnePost, auth_middleware.JWTMiddleware(), render_middleware.MarkdownRender())
//     postGroup.POST("/deleteOnePost", post.DeleteOnePost, auth_middleware.JWTMiddleware())
// }

func RegisterPostRoutes(app *echo.Echo) {
	postGroup := app.Group("/post")
	postGroup.POST("/getOnePost", post.GetOnePost)
	postGroup.GET("/getAllPosts", post.GetAllPosts)
	postGroup.POST("/createOnePost", post.CreateOnePost, render_middleware.MarkdownRender())
	postGroup.POST("/updateOnePost", post.UpdateOnePost, render_middleware.MarkdownRender())
	postGroup.POST("/deleteOnePost", post.DeleteOnePost)
}
