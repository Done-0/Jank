// Package redis 提供Redis连接和管理功能
// 创建者：Done-0
// 创建时间：2025-08-05
package redis

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"
)

// New 初始化 Redis 连接
// 参数：
//
//	config: 应用配置
func New(config *configs.Config) {
	client := newRedisClient(config)
	if err := client.Ping(context.Background()).Err(); err != nil {
		global.SysLog.Errorf("failed to ping Redis: %v", err)
		return
	}
	global.RedisClient = client
	global.SysLog.Infof("Redis connected successfully...")
}

// Close 关闭 Redis 连接
// 返回值：
//
//	error: 错误信息
func Close() error {
	if global.RedisClient == nil {
		return nil
	}

	if err := global.RedisClient.Close(); err != nil {
		return fmt.Errorf("failed to close Redis connection: %w", err)
	}

	global.SysLog.Info("Redis connection closed")
	return nil
}

// newRedisClient 创建新的 Redis 客户端
// 参数：
//
//	config: 应用配置
//
// 返回值：
//
//	*redis.Client: Redis 客户端实例
func newRedisClient(config *configs.Config) *redis.Client {
	db, _ := strconv.Atoi(config.RedisConfig.RedisDB)
	return redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.RedisConfig.RedisHost, config.RedisConfig.RedisPort),
		Password:     config.RedisConfig.RedisPassword, // 数据库密码，默认为空字符串
		DB:           db,                               // 数据库索引
		DialTimeout:  10 * time.Second,                 // 连接超时时间
		ReadTimeout:  1 * time.Second,                  // 读超时时间
		WriteTimeout: 2 * time.Second,                  // 写超时时间
		PoolSize:     runtime.GOMAXPROCS(10),           // 最大连接池大小
		MinIdleConns: 50,                               // 最小空闲连接数
	})
}
