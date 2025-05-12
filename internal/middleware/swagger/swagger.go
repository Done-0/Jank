package swagger_middleware

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/docs"
	"jank.com/jank_blog/internal/global"
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

// initSwagger 初始化 Swagger 配置信息
func initSwagger() {
	swaggerOnce.Do(func() {
		config, err := configs.LoadConfig()
		if err != nil {
			global.SysLog.Fatalf("加载 Swagger 配置失败: %v", err)
			return
		}

		docs.SwaggerInfo.Title = "Jank Blog API"
		docs.SwaggerInfo.Description = "这是 Jank Blog 的 API 文档，适用于账户管理、用户认证、角色权限管理，文章管理，类目管理、评论管理等功能。"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = config.SwaggerConfig.SwaggerHost
		if docs.SwaggerInfo.Host == "" {
			docs.SwaggerInfo.Host = "localhost:9010"
		}

		docs.SwaggerInfo.BasePath = "/"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}

		cmd := exec.Command("swag", "init", "-g", "pkg/router/router.go")
		output, err := cmd.CombinedOutput()
		if err != nil {
			global.SysLog.Errorf("初始化 Swagger 文档失败，错误: %v\n输出信息: %s", err, string(output))
			global.SysLog.Info("继续启动服务，但 Swagger 文档可能不可用")
		} else {
			global.SysLog.Info("成功生成 Swagger 文档")
		}

		fmt.Printf("Swagger service started on: http://%s/swagger/index.html\n", docs.SwaggerInfo.Host)
	})
}
