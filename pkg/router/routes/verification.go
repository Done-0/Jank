// Package routes 提供验证码路由注册功能
// 创建者：Done-0
// 创建时间：2025-08-10
package routes

import (
	"log"

	"github.com/cloudwego/hertz/pkg/route"

	"github.com/Done-0/jank/pkg/wire"
)

// RegisterVerificationRoutes 注册验证码相关路由
func RegisterVerificationRoutes(r *route.RouterGroup) {
	verificationController, err := wire.NewVerificationController()
	if err != nil {
		log.Fatalf("Failed to initialize verification controller: %v", err)
	}

	// 验证码路由组
	verificationGroup := r.Group("/verification")
	{
		verificationGroup.GET("/email", verificationController.SendEmailVerificationCode) // 发送邮箱验证码
	}
}
