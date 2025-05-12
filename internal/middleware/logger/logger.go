// Package logger_middleware 提供HTTP请求日志记录中间件
// 创建者：Done-0
// 创建时间：2025-05-10
package logger_middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	biz_err "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/global"
)

const CTX_KEY = "logger_processed"

// InitLogger 返回 HTTP 请求日志中间件，使用默认配置
func InitLogger() echo.MiddlewareFunc {
	return loggerWithConfig(defaultConfig)
}

// 日志字段键定义
const (
	LOG_KEY_REQ_ID  = "requestId" // 请求ID
	LOG_KEY_METHOD  = "method"    // HTTP方法
	LOG_KEY_URI     = "uri"       // 请求路径
	LOG_KEY_IP      = "ip"        // 客户端IP
	LOG_KEY_HOST    = "host"      // 主机名
	LOG_KEY_UA      = "ua"        // User-Agent
	LOG_KEY_STATUS  = "status"    // 状态码
	LOG_KEY_BYTES   = "bytes"     // 响应大小
	LOG_KEY_LATENCY = "latency"   // 响应时间(毫秒)
	LOG_KEY_BODY    = "body"      // 请求体
	LOG_KEY_ERROR   = "error"     // 错误信息
)

// loggerConfig 日志中间件配置
type loggerConfig struct {
	Skipper         func(echo.Context) bool // 跳过中间件的条件
	LogRequestBody  bool                    // 是否记录请求体
	LogResponseSize bool                    // 是否记录响应大小
	MaskSensitive   bool                    // 是否屏蔽敏感字段
	MaxBodySize     int                     // 最大记录请求体大小
	MaskValue       string                  // 敏感信息掩码值
	SensitiveFields map[string]struct{}     // 敏感字段列表
}

// 默认日志配置
var defaultConfig = loggerConfig{
	Skipper:         func(c echo.Context) bool { return false }, // 默认不跳过
	LogRequestBody:  true,                                       // 默认记录请求体
	LogResponseSize: true,                                       // 默认记录响应大小
	MaskSensitive:   true,                                       // 默认屏蔽敏感信息
	MaxBodySize:     20 * 1024,                                  // 默认最大10KB
	MaskValue:       "********",                                 // 默认掩码
	SensitiveFields: map[string]struct{}{ // 默认敏感字段
		"password": {}, "token": {}, "secret": {},
		"auth": {}, "key": {}, "credential": {},
	},
}

// sizeWriter 用于跟踪 HTTP 响应大小
type sizeWriter struct {
	http.ResponseWriter
	size int
}

// Write 实现 ResponseWriter 接口，记录写入的字节数
func (w *sizeWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.size += n
	return n, err
}

// writerPool 响应写入器对象池
var writerPool = sync.Pool{New: func() interface{} { return &sizeWriter{} }}

// loggerWithConfig 返回带自定义配置的日志中间件
func loggerWithConfig(config loggerConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = defaultConfig.Skipper
	}
	if config.MaxBodySize <= 0 {
		config.MaxBodySize = defaultConfig.MaxBodySize
	}
	if config.MaskValue == "" {
		config.MaskValue = defaultConfig.MaskValue
	}
	if len(config.SensitiveFields) == 0 && config.MaskSensitive {
		config.SensitiveFields = defaultConfig.SensitiveFields
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			// 避免重复处理
			if _, ok := c.Get(CTX_KEY).(bool); ok {
				return next(c)
			}
			c.Set(CTX_KEY, true)

			// 获取请求信息
			req := c.Request()
			reqID := c.Response().Header().Get(echo.HeaderXRequestID)
			if reqID == "" {
				reqID = middleware.DefaultRequestIDConfig.Generator()
				c.Response().Header().Set(echo.HeaderXRequestID, reqID)
			}

			// 初始化日志字段
			fields := logrus.Fields{
				LOG_KEY_REQ_ID: reqID,
				LOG_KEY_METHOD: req.Method,
				LOG_KEY_URI:    req.RequestURI,
				LOG_KEY_IP:     c.RealIP(),
				LOG_KEY_HOST:   req.Host,
				LOG_KEY_UA:     req.UserAgent(),
			}

			// 处理请求体
			if config.LogRequestBody && req.Body != nil && req.ContentLength > 0 && req.ContentLength < int64(config.MaxBodySize) {
				if body, _ := io.ReadAll(io.LimitReader(req.Body, int64(config.MaxBodySize))); len(body) > 0 {
					req.Body.Close()
					req.Body = io.NopCloser(bytes.NewReader(body))

					// 处理 JSON 请求体
					if len(body) > 2 && body[0] == '{' {
						var data map[string]interface{}
						if json.Unmarshal(body, &data) == nil {
							// 屏蔽敏感数据
							if config.MaskSensitive {
								var maskData func(map[string]interface{})
								maskData = func(data map[string]interface{}) {
									for k, v := range data {
										kl := strings.ToLower(k)
										for s := range config.SensitiveFields {
											if strings.Contains(kl, s) {
												data[k] = config.MaskValue
												break
											}
										}
										if m, ok := v.(map[string]interface{}); ok {
											maskData(m)
										}
									}
								}
								maskData(data)
							}

							if j, err := json.Marshal(data); err == nil {
								fields[LOG_KEY_BODY] = string(j)
							}
						}
					}
				}
			}

			// 设置响应跟踪
			var sw *sizeWriter
			if config.LogResponseSize {
				sw = writerPool.Get().(*sizeWriter)
				sw.ResponseWriter = c.Response().Writer
				sw.size = 0
				c.Response().Writer = sw
				defer writerPool.Put(sw)
			}

			// 执行请求处理
			start := time.Now()
			err := next(c)
			latency := time.Since(start)

			// 记录响应信息
			status := c.Response().Status
			fields[LOG_KEY_STATUS] = status
			fields[LOG_KEY_LATENCY] = float64(latency.Nanoseconds()) / 1e6

			if config.LogResponseSize && sw != nil {
				fields[LOG_KEY_BYTES] = sw.size
			}
			if err != nil {
				fields[LOG_KEY_ERROR] = err.Error()
			}

			log := global.SysLog.WithFields(fields)
			switch {
			case status >= 500:
				log.Error(biz_err.GetMessage(biz_err.SERVER_ERR))
			case status >= 400:
				log.Warn(biz_err.GetMessage(biz_err.BAD_REQUEST))
			default:
				log.Info(biz_err.GetMessage(biz_err.SUCCESS))
			}

			return err
		}
	}
}
