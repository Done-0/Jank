<p align="center">
  <a><img src="https://s2.loli.net/2025/01/02/6F8fzMvrBDCATZk.png" alt="Jank"></a>
</p>
<p align="center">
    <em>Jank，一个轻量级的博客系统，基于 Go 语言和 Echo 框架开发，强调极简、低耦合和高扩展</em>
</p>
<p align="center">
  <a href="https://img.shields.io/github/stars/Done-0/Jank?style=social" target="_blank">
    <img src="https://img.shields.io/github/stars/Done-0/Jank?style=social" alt="Stars">
  </a> &nbsp;
  <a href="https://img.shields.io/github/forks/Done-0/Jank?style=social" target="_blank">
    <img src="https://img.shields.io/github/forks/Done-0/Jank?style=social" alt="Forks">
  </a> &nbsp;
  <a href="https://img.shields.io/github/contributors/Done-0/Jank" target="_blank">
    <img src="https://img.shields.io/github/contributors/Done-0/Jank" alt="Contributors">
  </a> &nbsp;
  <a href="https://img.shields.io/github/issues/Done-0/Jank" target="_blank">
    <img src="https://img.shields.io/github/issues/Done-0/Jank" alt="Issues">
  </a> &nbsp;
  <a href="https://img.shields.io/github/issues-pr/Done-0/Jank" target="_blank">
    <img src="https://img.shields.io/github/issues-pr/Done-0/Jank" alt="Pull Requests">
  </a> &nbsp;
  <a href="https://img.shields.io/github/license/Done-0/Jank" target="_blank">
    <img src="https://img.shields.io/github/license/Done-0/Jank" alt="License">
  </a>
</p>

---

Jank 是一个轻量级的博客系统，基于 Go 语言和 Echo 框架开发，设计理念强调极简、低耦合和高扩展，旨在为用户提供功能丰富、界面简洁、操作简单且安全可靠的博客体验。

## 技术栈

- **Go 语言**：热门后端开发语言，适合构建高并发应用。
- **Echo 框架**：高性能的 Web 框架，支持快速开发和灵活的路由管理。
- **MySQL**：成熟的关系型数据库管理系统，支持复杂查询和事务处理。
- **Redis**：热门缓存解决方案，提供快速数据存取和持久化选项。
- **JWT**：安全的用户身份验证机制，确保数据传输的完整性和安全性。
- **Docker**：容器化部署工具，简化应用的打包和分发流程。
- **前端**：Vue 3 + Nuxt + Shadcn-vue（暂时搁置）。

## 功能模块

- **账户模块**：实现 JWT 身份验证，支持用户登录、注册、注销、密码修改和个人信息更新。
- **文章模块**：提供文章的创建、查看、更新和删除功能。
- **分类模块**：支持类目树及子类目树递归查询，单一类目查询，以及类目的创建、更新和删除。
- **其他功能**：
  - 提供 OpenAPI 接口文档
  - 集成 Air 实现热重载
  - 使用 Logrus 实现日志记录
  - 支持 CORS 跨域请求
  - 提供 CSRF 和 XSS 防护
  - 支持 Markdown 的服务端渲染

> **其他模块正在开发中**，欢迎提供宝贵意见和建议！

## 本地开发

1. **安装依赖**：

   ```bash
   go mod tidy
   ```

2. **配置数据库和邮箱**：  
   修改 `configs/config.yaml` 文件中的数据库配置和邮箱配置，示例如下：

   ```yaml
   # 数据库密码
   DB_PSW: "<DATABASE_PASSWORD>"

   # QQ 邮箱和 SMTP 授权码（可选）
   QQ_SMTP: "<QQ_SMTP>"
   FROM_EMAIL: "<FROM_QQ_EMAIL>"
   ```

3. **启动服务**：  
   使用以下命令启动应用：

   ```bash
   go run main.go
   ```

   或使用 Air 进行热重载：

   ```bash
   air -c ./configs/.air.toml
   ```

4. **访问接口**：  
   打开浏览器，访问 [http://localhost:9010/ping](http://localhost:9010/ping)

## Docker 容器部署

```bash
docker-compose up -d
```

## 官方社区

如果有任何疑问或建议，欢迎加入官方社区交流。

<img src="https://s2.loli.net/2025/01/04/cVqDO7a4djAPmEJ.jpg" alt="官方社区" width="300" />

## 代码统计

```bash
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              67            530            464           4002
JSON                             1              0              0           1532
YAML                             3              8              6           1030
Markdown                        26             33              0            135
TOML                             1              6              0             36
-------------------------------------------------------------------------------
TOTAL                           98            577            470           6735
-------------------------------------------------------------------------------
```

## 许可证

本项目遵循 [MIT 协议](https://opensource.org/licenses/MIT)。

## GitHub 统计

<img src="https://github-readme-stats.vercel.app/api?username=Done-0&show_icons=true&hide_title=true&theme=radical" width="100%" height="65%">

## 增长趋势

<img src="https://api.star-history.com/svg?repos=Done-0/Jank&type=timeline" width="100%" height="65%">
