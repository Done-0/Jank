<p style="text-align: center;">
  <a><img src="https://s2.loli.net/2025/01/02/6F8fzMvrBDCATZk.png" alt="Jank"></a>
</p>
<p style="text-align: center;">
  <em>Jank，一个轻量级的博客系统，基于 Go 语言和 Echo 框架开发，强调极简、低耦合和高扩展</em>
</p>
<p style="text-align: center;">
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
## Jank-standalone

**本版本为单用户独立部署版，系统只能存在一个账号，请部署后立即注册账号，避免出现安全问题。**

> 注：本项目当前缺少前端部分，在此诚邀有志于前端开发的开发者加入，共同参与开发工作，期待您的宝贵意见和贡献！

## 预览

👉[【b站视频预览: 你见过 Go 语言开发的博客系统吗？】](https://www.bilibili.com/video/BV1W5wdeZEoY/?share_source=copy_web&vd_source=6fd45877cd498bfb9c2b449d1197363c)

![home-white.png](https://raw.githubusercontent.com/Done-0/Jank-website/main/public/images/home-white.png)
![home-black.png](https://raw.githubusercontent.com/Done-0/Jank-website/main/public/images/home-black.png)

## 技术栈

- **Go 语言**：热门后端开发语言，适合构建高并发应用。
- **Echo 框架**：高性能的 Web 框架，支持快速开发和灵活的路由管理。
- **PostgreSQL**：开源的关系型数据库，提供高性能、高可靠性的数据存储。
- **Redis**：热门缓存解决方案，提供快速数据存取和持久化选项。
- **JWT**：安全的用户身份验证机制，确保数据传输的完整性和安全性。
- **Docker**：容器化部署工具，简化应用的打包和分发流程。
- **前端**：Vue 3 + Nuxt + Shadcn-vue（原项目已不再维护）。
  > 注：目前缺少前端部分，欢迎有志于前端开发的开发者加入！

## 功能模块

- **账户模块**：实现 JWT 身份验证，支持用户登录、注册、注销、密码修改和个人信息更新。
- **文章模块**：提供文章的创建、查看、更新和删除功能。
- **分类模块**：支持类目树及子类目树递归查询，单一类目查询，以及类目的创建、更新和删除。
- **评论模块**：提供评论的创建、查看、删除和回复功能，支持评论树结构的展示。
- **插件系统**：正在开发中...
- **其他功能**：
    - 集成 Air 实现热重载
    - 提供 Logrus 实现日志记录
    - 支持 CORS 跨域请求
    - 提供 CSRF 和 XSS 防护
    - 支持 Markdown 的服务端渲染
    - **其他模块正在开发中**，欢迎提供宝贵意见和建议！

## 本地开发

1. **安装依赖**：

   ```bash
   # 安装依赖包
   go mod tidy
   ```

2. **配置数据库和邮箱**：  
   修改 `configs/config.yaml` 文件中的数据库配置和邮箱配置，示例如下：

   ```yaml
   # mysql 数据库配置
    DB_USER: "<DATABASE_USER>"
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

   > 此方法最为便捷，但提前配置环境变量 GOPATH。

   ```bash
   # 安装 air，需要 go 1.22 或更高版本
   go install github.com/air-verse/air@latest

   # 热重载启动
   air -c ./configs/.air.toml
   ```

4. **访问接口**：  
   本地启动应用后，浏览器访问 [http://localhost:9010/ping](http://localhost:9010/ping)

## Docker 容器部署

1. 修改 `configs/config.yaml` 文件中的数据库配置和邮箱配置，示例如下：

   ```yaml
   # 应用配置
   APP_HOST: "0.0.0.0"

   # mysql 数据库配置
    DB_USER: "<DATABASE_USER>"
    DB_PSW: "<DATABASE_PASSWORD>"

   # QQ 邮箱和 SMTP 授权码（可选）
   QQ_SMTP: "<QQ_SMTP>"
   FROM_EMAIL: "<FROM_QQ_EMAIL>"
   ```

2. 修改 `docker-compose.yaml` 文件中的环境变量，示例如下：

   ```yaml
   environment:
      - POSTGRES_USER=<DATABASE_USER>
      - POSTGRES_PASSWORD=<DATABASE_PASSWORD>
   
   healthcheck:
      test: ["CMD", "pg_isready", "-U", "<DATABASE_USER>", "-d", "jank_db"]
   ```

3. 启动容器：

    ```bash
    docker-compose up -d
    ```

## 官方社区

如果有任何疑问或建议，欢迎加入官方社区交流。

<img src="https://s2.loli.net/2025/01/25/L9BspuHnrIeim7S.jpg" alt="官方社区" width="300" />

## 特别鸣谢

感谢各位对本项目的支持！

<div style="display: flex; flex-wrap: wrap;">
  <img src="https://s2.loli.net/2025/02/21/B6Aq9HVOGvJzEyI.jpg" alt="c" style="border-radius: 50%; width: 120px; height: 120px; margin: 10px;" />
</div>

## 联系合作

- **QQ**: 927171598
- **邮箱**：<EMAIL>fenderisfine@outlook.com

## 许可证

本项目遵循 [MIT 协议](https://opensource.org/licenses/MIT)。

## GitHub 统计

<img src="https://github-readme-stats.vercel.app/api?username=Done-0&show_icons=true&hide_title=true&theme=radical" width="100%" height="65%" alt="GitHub Stats">

## 增长趋势

<img src="https://api.star-history.com/svg?repos=Done-0/Jank&type=timeline" width="100%" height="65%" alt="GitHub Stats">