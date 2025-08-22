// Package cors 提供CORS跨域中间件配置
// 创建者：Done-0
// 创建时间：2025-08-05
package cors

import (
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/types/consts"
)

// New 创建 CORS 中间件
// 返回值：
// app.HandlerFunc: CORS中间件
func New() app.HandlerFunc {
	cfgs, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	return cors.New(cors.Config{
		AllowOrigins: cfgs.AppConfig.CORSConfig.AllowOrigins,
		AllowMethods: []string{
			consts.MethodGET,
			consts.MethodPOST,
			consts.MethodPUT,
			consts.MethodDELETE,
			consts.MethodOPTIONS,
			consts.MethodPATCH,
		},
		AllowHeaders: []string{
			consts.HeaderOrigin,
			consts.HeaderContentType,
			consts.HeaderAccept,
			consts.HeaderAuthorization,
			consts.HeaderXRequestedWith,
		},
		ExposeHeaders: []string{
			consts.HeaderContentLength,
			consts.HeaderAuthorization,
		},
		AllowCredentials: cfgs.AppConfig.CORSConfig.AllowCredentials,
		MaxAge:           time.Duration(cfgs.AppConfig.CORSConfig.MaxAge) * time.Hour,
	})
}
