// Package dto 提供验证码相关的数据传输对象定义
// 创建者：Done-0
// 创建时间：2025-08-10
package dto

// SendEmailCodeRequest 发送邮箱验证码请求
type SendEmailCodeRequest struct {
	Email string `query:"email" validate:"required,email"` // 邮箱地址
}
