// Package logger 提供应用程序日志功能的初始化和配置
// 创建者：Done-0
// 创建时间：2025-05-10
package logger

import (
	"io"
	"log"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
)

// New 初始化日志组件
func New() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("初始化日志组件时加载配置失败: %v", err)
		return
	}

	// 设置路径
	logFilePath := cfg.LogConfig.LogFilePath
	logFileName := cfg.LogConfig.LogFileName
	fileName := path.Join(logFilePath, logFileName)
	_ = os.MkdirAll(logFilePath, 0755)

	// 初始化 logger
	formatter := &logrus.JSONFormatter{TimestampFormat: cfg.LogConfig.LogTimestampFmt}
	logger := logrus.New()
	logger.SetFormatter(formatter)
	logger.SetOutput(io.Discard)

	// 设置日志级别
	logLevel, err := logrus.ParseLevel(cfg.LogConfig.LogLevel)
	switch err {
	case nil:
		logger.SetLevel(logLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	// 配置日志轮转
	writer, err := rotatelogs.New(
		path.Join(logFilePath, "%Y%m%d.log"),
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Duration(cfg.LogConfig.LogMaxAge)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(cfg.LogConfig.LogRotationTime)*time.Hour),
	)

	switch {
	case err != nil:
		log.Printf("配置日志轮转失败: %v，使用标准文件", err)
		fileHandle, fileErr := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)

		switch {
		case fileErr != nil:
			log.Printf("创建日志文件失败: %v，使用标准输出", fileErr)
			logger.SetOutput(os.Stdout)
			global.LogFile = nil
		default:
			logger.SetOutput(fileHandle)
			global.LogFile = fileHandle
		}
	default:
		allLevels := []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
			logrus.InfoLevel,
			logrus.DebugLevel,
			logrus.TraceLevel,
		}

		writeMap := make(lfshook.WriterMap, len(allLevels))
		for _, level := range allLevels {
			writeMap[level] = writer
		}

		logger.AddHook(lfshook.NewHook(writeMap, formatter))

		global.LogFile = writer
	}

	global.SysLog = logger
}
