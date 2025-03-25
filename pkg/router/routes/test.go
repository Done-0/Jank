package routes

import (
	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/pkg/serve/controller/test"
)

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
