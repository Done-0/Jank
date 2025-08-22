// Package consts 提供应用程序常量定义
// 创建者：Done-0
// 创建时间：2025-08-05
package consts

// HTTP 方法常量
const (
	MethodGET     = "GET"     // GET 方法
	MethodPOST    = "POST"    // POST 方法
	MethodPUT     = "PUT"     // PUT 方法
	MethodDELETE  = "DELETE"  // DELETE 方法
	MethodOPTIONS = "OPTIONS" // OPTIONS 方法
	MethodPATCH   = "PATCH"   // PATCH 方法
)

// HTTP 头部常量
const (
	// CORS 相关头部
	HeaderOrigin         = "Origin"           // 请求来源
	HeaderContentType    = "Content-Type"     // 内容类型
	HeaderAccept         = "Accept"           // 接受类型
	HeaderAuthorization  = "Authorization"    // 授权头部
	HeaderXRequestedWith = "X-Requested-With" // AJAX请求标识
	HeaderContentLength  = "Content-Length"   // 内容长度

	// 网络相关头部
	HeaderRequestID     = "X-Request-ID"    // 请求 ID 头部
	HeaderXForwardedFor = "X-Forwarded-For" // 代理转发的原始客户端 IP
	HeaderXRealIP       = "X-Real-IP"       // 真实客户端 IP
	HeaderXClientIP     = "X-Client-IP"     // 客户端 IP（某些代理使用）
	HeaderUserAgent     = "User-Agent"      // 用户代理字符串
)
