// Package global 提供全局变量和对象定义
// 创建者：Done-0
// 创建时间：2025-08-05
package global

import (
	"io"

	"github.com/casbin/casbin/v2"
	"github.com/hashicorp/go-plugin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 数据库相关全局变量
var (
	DB *gorm.DB // 全局 db 对象，用于数据库操作
)

// 缓存相关全局变量
var (
	RedisClient *redis.Client // 全局 redis 客户端对象，用于缓存操作
)

// 日志相关全局变量
var (
	SysLog  *logrus.Logger // 全局系统级日志对象，用于记录系统级日志
	BizLog  *logrus.Entry  // 全局业务级日志对象，用于记录业务级日志
	LogFile io.Closer      // 全局日志文件对象，用于日志文件资源管理
)

// 权限相关全局变量
var (
	Enforcer *casbin.Enforcer // 全局 Casbin 执行器对象，用于权限控制
)

// 插件系统相关全局变量
var (
	PluginClients map[string]*plugin.Client // 全局插件客户端映射，用于插件系统管理
)
