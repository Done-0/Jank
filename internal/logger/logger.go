package logger

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	logrus "github.com/sirupsen/logrus"
	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
)

func init() {
	initLogger()
}

func initLogger() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("初始化日志组件时加载配置失败: %v", err)
	}

	logFilePath := cfg.LogConfig.LogFilePath
	logFileName := cfg.LogConfig.LogFileName

	// 打开日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 0755: Unix/Linux 系统中常用的文件权限表示法。使用八进制（octal）数字系统来表示文件或目录的权限。每个数字表示一组权限，分别对应用户、用户组和其他人
	// 第一个数字（0）：表示文件类型。对于常规文件，通常为 0
	// 第二个数字（7）：表示文件所有者（用户）的权限 (这里 7 表示文件所有者拥有读（4）、写（2）和执行（1）的权限，合计 4 + 2 + 1 = 7)
	// 第三个数字（5）：表示与文件所有者同组的用户组的权限 (这里 5 表示用户组和其他用户拥有读（4）和执行（1）的权限，合计 4 + 1 = 5)
	// 第四个数字（5）：表示其他用户的权限
	// 因此 0755 表示：
	// 文件所有者可以读、写、执行。
	// 用户组成员可以读、执行。
	// 其他用户可以读、执行。
	// 创建日志文件目录
	_ = os.Mkdir(logFilePath, 0755)
	global.LogFile, _ = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)

	// 初始化 logrus
	logger := logrus.New()
	log.Printf(cfg.LogConfig.LogTimestampFmt)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: cfg.LogConfig.LogTimestampFmt,
	})
	logger.Out = global.LogFile
	logLevel, err := logrus.ParseLevel(cfg.LogConfig.LogLevel)
	if err != nil {
		return
	}
	logger.SetLevel(logLevel)

	// 设置日志轮转
	maxAge := time.Duration(cfg.LogConfig.LogMaxAge) * time.Hour
	rotationTime := time.Duration(cfg.LogConfig.LogRotationTime) * time.Hour

	writer, err := rotatelogs.New(
		logFilePath+"%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		return
	}

	// 配置日志级别与轮转日志的映射
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.FatalLevel: writer,
		logrus.DebugLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
	}

	// 添加钩子到 logrus
	hook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: cfg.LogConfig.LogTimestampFmt,
	})
	logger.AddHook(hook)
	global.SysLog = logger
}

func New() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqId := c.Response().Header().Get(echo.HeaderXRequestID)
			if reqId == "" {
				reqId = middleware.DefaultRequestIDConfig.Generator()
				c.Response().Header().Set(echo.HeaderXRequestID, reqId)
			}
			bizLog := global.SysLog.WithFields(logrus.Fields{
				"requestId": reqId,
				"requestIp": c.RealIP(),
			})
			// 将 BizLog 存储到当前请求上下文中
			c.Set("BizLog", bizLog)
			return next(c)
		}
	}
}
