### Jank 简客博客系统

Jank 是一个极简、轻量且高度可扩展的博客系统，基于 Go 语言和 Echo 框架开发。

#### 技术栈

- **Go 语言**：用于后端开发
- **Echo 框架**：用于构建高性能 Web 应用
- **MySQL**：关系型数据库
- **Redis**：缓存系统
- **JWT**：用于用户身份验证
- **Docker**：容器化部署
- **前端**：Vue 3 + Nuxt + Shadcn-vue（因前端开发经验不足，暂时废弃）

#### 功能模块

- **账户模块**：提供 JWT 身份验证、登录、注册、注销、修改密码和编辑个人信息
- **文章模块**：支持文章的创建、查看、编辑和删除
- **分类模块**：支持类目树/子类目树递序查询、单一类目查询、类目创建、类目编辑和类目删除
- **其他功能**：
  - OpenAPI 接口文档
  - Air 热重载
  - Logrus 日志记录
  - CORS 跨域支持
  - CSRF 防护
  - XSS 防护
  - Markdown 服务端渲染

> **其他模块正在开发中**，欢迎提供宝贵意见和建议！

#### 本地开发

1. 安装依赖：

   ```bash
   go mod tidy
   ```

2. 配置数据库和邮箱：  
   修改 `configs/config.yaml` 文件中的数据库配置和邮箱配置，示例如下：

   ```yaml
   DB_PSW: "<DATABASE_PASSWORD>"

   # QQ 邮箱和 SMTP 授权码（可选）
   QQ_SMTP: "<QQ_SMTP>"
   FROM_EMAIL: "<FROM_QQ_EMAIL>"
   ```

3. 启动服务：  
   使用以下命令启动应用：

   ```bash
   go run main.go
   ```

   或使用 Air 进行热重载：

   ```bash
   air -c ./configs/.air.toml
   ```

4. 访问接口：  
   打开浏览器，访问 [http://localhost:9010/ping](http://localhost:9010/ping)

#### Docker 容器部署

```bash
docker-compose up -d
```

#### 联系方式

- **QQ 群**：828270460
- **作者 QQ**：927171598

#### 许可证

本项目采用 [MIT 开源许可证](https://opensource.org/licenses/MIT)。

#### GitHub 统计

![Jank GitHub Stats](https://github-readme-stats.vercel.app/api?username=Done-0&show_icons=true&count_private=true&hide=prs&theme=radical)

[![Jank Stars](https://img.shields.io/github/stars/Done-0/Jank?style=social)](https://github.com/Done-0/Jank)

#### 特别鸣谢

[![Contributors](https://img.shields.io/github/contributors/Done-0/Jank)](https://github.com/Done-0/Jank/graphs/contributors)
