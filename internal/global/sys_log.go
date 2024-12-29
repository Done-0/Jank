package global

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// 用于记录系统级别日志。
var SysLog *log.Logger

// 用于记录特定业务逻辑日志。
// Entry 可以在日志中包含附加的上下文信息。
var BizLog *log.Entry

// 指向日志文件。
var LogFile *os.File
