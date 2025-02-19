package swaggerMiddleware

import (
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/docs"
)

var swaggerOnce sync.Once

func InitSwagger() echo.MiddlewareFunc {
	initSwagger()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasPrefix(c.Request().URL.Path, "/swagger/") {
				return echoSwagger.WrapHandler(c)
			}
			return next(c)
		}
	}
}

// 初始化 Swagger 配置信息
func initSwagger() {
	swaggerOnce.Do(func() {
		time.Sleep(2 * time.Second)
		config, err := configs.LoadConfig()
		if err != nil {
			log.Fatalf("配置加载失败: %v", err)
		}

		docs.SwaggerInfo.Title = "Jank Blog API"
		docs.SwaggerInfo.Description = "这是 Jank Blog 的 API 文档，适用于账户管理、用户认证，文章管理，类目管理等功能。"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = config.SwaggerConfig.SwaggerHost
		if docs.SwaggerInfo.Host == "" {
			docs.SwaggerInfo.Host = "localhost:9010"
		}

		docs.SwaggerInfo.BasePath = "/"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}

		cmd := exec.Command("swag", "init")
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("初始化 Swagger 文档失败，错误: %v\n输出信息: %s", err, string(output))
		}

		log.Printf("Swagger service started on: http://%s/swagger/index.html", docs.SwaggerInfo.Host)
	})
}
