# 数据库交互组件 (Database Interaction Component)

## 简介

数据库交互组件负责应用程序与各种数据库系统的连接和交互，支持多种数据库类型，包括 PostgreSQL、MySQL 和 SQLite。该组件提供了数据库连接、自动迁移和管理功能，为上层业务逻辑提供稳定的数据存储和访问支持。

## 支持的数据库类型

- **PostgreSQL**: 企业级关系型数据库，适合大规模应用
- **MySQL**: 流行的开源关系型数据库，性能优良
- **SQLite**: 轻量级嵌入式数据库，适合开发和小型应用

## 核心功能

- **数据库连接管理**: 建立和管理与不同类型数据库的连接
- **自动创建数据库**: 在系统级数据库中自动创建应用数据库（针对 PostgreSQL 和 MySQL）
- **自动迁移**: 根据模型定义自动创建和更新数据库表结构
- **多数据库支持**: 根据配置灵活切换不同类型的数据库
- **连接参数配置**: 支持连接超时、字符集等参数配置

## 实现细节

- 基于 GORM 框架实现 ORM（对象关系映射）功能
- 使用特定数据库驱动：`gorm.io/driver/postgres`、`gorm.io/driver/mysql` 和 `gorm.io/driver/sqlite`
- 支持数据库方言设置，可以通过配置切换数据库类型
- 提供数据库不存在时的自动创建逻辑
- 支持数据库路径和文件权限管理（特别是 SQLite）

## 数据库配置参数

数据库组件从应用配置中读取以下参数：

- **DBDialect**: 数据库类型（`POSTGRES`、`MYSQL` 或 `SQLITE`）
- **DBHost**: 数据库服务器主机地址
- **DBPort**: 数据库服务器端口
- **DBUser**: 数据库用户名
- **DBPassword**: 数据库密码
- **DBName**: 应用数据库名称
- **DBPath**: SQLite 数据库文件路径

## 使用方式

通过全局变量 `global.DB` 在应用的任何位置访问数据库：

```go
// 查询数据
user := new(model.Account)
if err := global.DB.Where("email = ?", email).First(user).Error; err != nil {
    // 处理错误
}

// 创建数据
newPost := &model.Post{Title: "新文章", ContentMarkdown: "# 标题"}
if err := global.DB.Create(newPost).Error; err != nil {
    // 处理错误
}
```

## 自动迁移

应用启动时会自动执行数据库迁移，根据模型定义创建或更新表结构：

```go
// 获取所有模型
models := model.GetAllModels()

// 执行迁移
global.DB.AutoMigrate(models...)
```
