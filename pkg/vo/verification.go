// Package vo 提供验证码相关的视图对象定义
// 创建者：Done-0
// 创建时间：2025-08-10
package vo

// SendEmailCodeResponse 发送邮箱验证码响应
type SendEmailCodeResponse struct {
	Message string `json:"message"` // 响应消息
}
