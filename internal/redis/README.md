# Redis 组件

Redis 组件用于处理应用程序与 Redis 缓存服务器的连接和交互。该组件提供了一个全局 Redis 客户端实例，可以在整个应用程序中使用。

## 功能

- **连接管理**: 创建并维护与 Redis 服务器的连接
- **连接池**: 配置最优的连接池设置，包括最大连接数和最小空闲连接数
- **超时控制**: 设置合理的连接、读取和写入超时时间
- **健康检查**: 确保 Redis 连接正常工作

## 配置项

Redis 连接从应用配置中读取以下参数：

- 主机地址 (RedisHost)
- 端口 (RedisPort)
- 密码 (RedisPassword)
- 数据库索引 (RedisDB)

## 使用方式

通过全局变量 `global.RedisClient` 在应用的任何位置访问 Redis 客户端，例如：

```go
// 设置缓存
global.RedisClient.Set(context.Background(), key, value, expiration)

// 获取缓存
global.RedisClient.Get(context.Background(), key)
```
