package global

import (
	"io"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB      // 全局 db 对象
	RedisClient *redis.Client // 全局 redis 客户端对象
)

var (
	SysLog  *logrus.Logger // 全局系统级日志对象
	BizLog  *logrus.Entry  // 全局业务级日志对象
	LogFile io.Closer      // 全局日志文件对象
)
