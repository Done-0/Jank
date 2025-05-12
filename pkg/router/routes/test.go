// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-05-10
package routes

import (
	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/pkg/serve/controller/test"
)

// RegisterTestRoutes 注册测试相关路由
// 参数：
//   - r: Echo 路由组数组，r[0] 为 API v1 版本组，r[1] 为 API v2 版本组
func RegisterTestRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	testGroupV1 := apiV1.Group("/test")
	testGroupV1.GET("/testPing", test.TestPing)
	testGroupV1.GET("/testHello", test.TestHello)
	testGroupV1.GET("/testLogger", test.TestLogger)
	testGroupV1.GET("/testRedis", test.TestRedis)
	testGroupV1.GET("/testSuccessRes", test.TestSuccRes)
	testGroupV1.GET("/testErrRes", test.TestErrRes)
	testGroupV1.GET("/testErrorMiddleware", test.TestErrorMiddleware)

	// api v2 group
	apiV2 := r[1]
	testGroupV2 := apiV2.Group("/test")
	testGroupV2.GET("/testLongReq", test.TestLongReq)
}
