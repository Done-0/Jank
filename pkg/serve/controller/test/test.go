package testapi

import (
	"github.com/labstack/echo/v4"
	"jank.com/jank_blog/pkg/vo"
)

// @Summary		Ping API
// @Description	Test if the API is alive
// @Tags			test
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]interface{}
// @Router			/test/ping [get]
func Ping(c echo.Context) error {
	return c.JSON(200, vo.Success("Pong", c))
}
