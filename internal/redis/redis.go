package redis

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
)

var ctx = context.Background()

func New(config *configs.Config) {
	client := newRedisClient(config)
	if err := client.Ping(ctx).Err(); err != nil {
		global.SysLog.Errorf("Redis 连接失败: %v", err)
		return
	}
	global.Redis = client
	global.SysLog.Infof("Redis 连接成功!")
}

func newRedisClient(config *configs.Config) *redis.Client {
	db, _ := strconv.Atoi(config.RedisDB)
	return redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:6379", config.RedisHost),
		Password:     config.RedisPassword,   // 数据库密码，默认为空字符串
		DB:           db,                     // 数据库索引
		DialTimeout:  10 * time.Second,       // 连接超时时间
		ReadTimeout:  1 * time.Second,        // 读超时时间
		WriteTimeout: 2 * time.Second,        // 写超时时间
		PoolSize:     runtime.GOMAXPROCS(10), // 最大连接池大小
		MinIdleConns: 50,                     // 最小空闲连接数
	})
}