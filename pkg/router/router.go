package router

import (
    "github.com/labstack/echo/v4"
    "jank.com/jank_blog/pkg/router/routes"
)

//	@title			Jank Blog API
//	@version		1.0
//	@description	This is the API documentation for Jank Blog.
//	@host			localhost:9010
//	@BasePath		/
func RegisterRoutes(app *echo.Echo) {    
    routes.RegisterTestRoutes(app)
    routes.RegisterAccountRoutes(app)
    routes.RegisterPostRoutes(app)
    routes.RegisterCategoryRoutes(app)
}