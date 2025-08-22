// Package logger 提供应用程序日志功能的初始化和配置
// 创建者：Done-0
// 创建时间：2025-08-05
package logger

import (
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"

	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
)

// New 初始化日志组件
// 参数：
//
//	config: 配置信息
func New(config *configs.Config) {
	logFilePath := config.LogConfig.LogFilePath
	logFileName := config.LogConfig.LogFileName
	fileName := path.Join(logFilePath, logFileName)
	_ = os.MkdirAll(logFilePath, 0755)

	// 初始化 logger
	formatter := &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"}
	logger := logrus.New()
	logger.SetFormatter(formatter)

	// 设置日志级别
	logLevel, err := logrus.ParseLevel(config.LogConfig.LogLevel)
	switch err {
	case nil:
		logger.SetLevel(logLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	// 配置日志轮转
	writer, err := rotateLogs.New(
		path.Join(logFilePath, "%Y%m%d.log"),
		rotateLogs.WithLinkName(fileName),
		rotateLogs.WithMaxAge(time.Duration(config.LogConfig.LogMaxAge)*24*time.Hour),
		rotateLogs.WithRotationTime(24*time.Hour),
	)

	switch {
	case err != nil:
		log.Printf("Failed to initialize log file rotation: %v, using standard output", err)
		fileHandle, fileErr := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)

		switch {
		case fileErr != nil:
			log.Printf("Failed to create log file: %v, using standard output", fileErr)
			logger.SetOutput(os.Stdout)
			global.LogFile = nil
		default:
			logger.SetOutput(io.MultiWriter(os.Stdout, fileHandle))
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

		logger.SetOutput(os.Stdout)

		global.LogFile = writer
	}

	global.SysLog = logger
	global.BizLog = logger.WithField("module", "business")
}
