package routes

import (
	"github.com/labstack/echo/v4"

	auth_middleware "jank.com/jank_blog/internal/middleware/auth"
	"jank.com/jank_blog/pkg/serve/controller/comment"
)

func RegisterCommentRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	commentGroupV1 := apiV1.Group("/comment")
	commentGroupV1.GET("/getOneComment", comment.GetOneComment)
	commentGroupV1.GET("/getCommentGraph", comment.GetCommentGraph)
	commentGroupV1.POST("/createOneComment", comment.CreateOneComment, auth_middleware.AuthMiddleware())
	commentGroupV1.POST("/deleteOneComment", comment.DeleteOneComment, auth_middleware.AuthMiddleware())
}
