package test

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/vo"
)

// TestPing          @Summary       Ping API
// @Description  测试接口
// @Tags         test
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "Pong successfully!\n"
// @Router       /test/testPing [get]
func TestPing(c echo.Context) error {
	utils.BizLogger(c).Info("Ping...")
	return c.String(http.StatusOK, "Pong successfully!\n")
}

// TestHello         @Summary       Hello API
// @Description  测试接口
// @Tags         test
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "Hello, Jank 🎉!\n"
// @Router       /test/testHello [get]
func TestHello(c echo.Context) error {
	utils.BizLogger(c).Info("Hello, Jank!")
	return c.String(http.StatusOK, "Hello, Jank 🎉!\n")
}

// TestLogger    @Summary       测试日志接口
// @Description  用于测试日志功能
// @Tags         test
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "测试日志成功!"
// @Router       /test/testLogger [get]
func TestLogger(c echo.Context) error {
	utils.BizLogger(c).Infof("测试日志...")
	return c.String(http.StatusOK, "测试日志成功!")
}

// TestRedis     @Summary      测试 Redis 接口
// @Description  用于测试 Redis 功能
// @Tags         test
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "测试缓存功能完成!"
// @Router       /test/testRedis [get]
func TestRedis(c echo.Context) error {
	utils.BizLogger(c).Infof("开始写入缓存...")
	err := global.RedisClient.Set(c.Request().Context(), "TEST:", "测试 value", 0).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("测试写入缓存失败: %v", err)
		return err
	}
	utils.BizLogger(c).Infof("写入缓存成功...")

	utils.BizLogger(c).Infof("开始读取缓存...")
	articlesCache, err := global.RedisClient.Get(c.Request().Context(), "TEST:").Result()
	if err != nil {
		utils.BizLogger(c).Errorf("测试读取缓存失败: %v", err)
		return err
	}
	utils.BizLogger(c).Infof("读取缓存成功, key: %s , value: %s", "TEST:", articlesCache)
	return c.String(http.StatusOK, "测试缓存功能完成!")
}

// TestSuccRes   @Summary       测试成功响应接口
// @Description  用于测试成功响应
// @Tags         test
// @Accept       json
// @Produce      json
// @Success      200  {object}  vo.Result "测试成功响应成功!"
// @Router       /test/testSuccessRes [get]
func TestSuccRes(c echo.Context) error {
	utils.BizLogger(c).Info("测试成功响应...")
	return c.JSON(http.StatusOK, vo.Success(c, "测试成功响应成功!"))
}

// TestErrRes    @Summary      测试错误响应接口
// @Description  用于测试错误响应
// @Tags         test
// @Accept       json
// @Produce      json
// @Success      500  {object}  vo.Result
// @Router       /test/testErrRes [get]
func TestErrRes(c echo.Context) error {
	utils.BizLogger(c).Info("测试失败响应...")
	return c.JSON(http.StatusInternalServerError, vo.Fail(c, nil, bizErr.New(bizErr.SERVER_ERR)))
}

// TestErrorMiddleware         @Summary    测试错误处理中间件接口
// @Description  用于测试错误中间件
// @Tags         test
// @Accept       json
// @Produce      json
// @Success      500  {string}  nil
// @Router       /test/testErrorMiddleware [get]
func TestErrorMiddleware(c echo.Context) error {
	utils.BizLogger(c).Info("测试错误处理中间件...")
	panic("测试错误处理中间件...")
}

// TestLongReq       @Summary       长时间请求接口
// @Description  模拟一个耗时请求
// @Tags         test
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "模拟耗时请求处理完成!\n"
// @Router       /test/testLongReq [get]
func TestLongReq(c echo.Context) error {
	utils.BizLogger(c).Info("开始测试耗时请求...")
	time.Sleep(20 * time.Second)
	return c.String(http.StatusOK, "模拟耗时请求处理完成!\n")
}
