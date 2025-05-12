<p style="text-align: center;">
  <a><img src="https://s2.loli.net/2025/03/14/BnchjpPLeIaoO75.png" alt="Jank"></a>
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
<p align="center">
  <span style="text-decoration: underline; color: grey;">简体中文</span> | <a href="README_en.md" style="text-decoration: none;">English</a>
</p>

---

Jank 是一个轻量级的博客系统，基于 Go 语言和 Echo 框架开发，设计理念强调极简、低耦合和高扩展，旨在为用户提供功能丰富、界面简洁、操作简单且安全可靠的博客体验。

## 产品动态

- 4 月 20 日，Jank 首届线上交流会。
- 4 月 27 日，[Jank-Community](https://github.com/Jank-Community) 组织。

> 官方社区见本文下方[官方社区](https://github.com/Done-0/Jank/edit/main/README.md#%E5%AE%98%E6%96%B9%E7%A4%BE%E5%8C%BA)部分，欢迎各位优秀开发者参与贡献，共同推进 Jank 的开发！

## 速览

👉 演示站｜ Demo：[https://www.jank.org.cn](https://www.jank.org.cn)

👉[【Jank 博客系统】全新技术栈与 UI】](https://www.bilibili.com/video/BV1bjQ8YNEEo/?share_source=copy_web&vd_source=6fd45877cd498bfb9c2b449d1197363c)

👉 前端仓库：[https://github.com/Done-0/Jank-website](https://github.com/Done-0/Jank-website)

![home-page.png](https://s2.loli.net/2025/04/07/l1tGYV4WkmoiIHv.png)
![posts-page.png](https://s2.loli.net/2025/04/07/xR62vhWKsmgw3Ht.png)
![post1-page.png](https://s2.loli.net/2025/04/07/DbcJzryKmBNR7vQ.png)
![post2-page.png](https://s2.loli.net/2025/04/07/iNpXyMdkjaDbn92.png)

> 注：因为还在推出阶段，部分配置文件可能需要根据实际情况更改，具体请使用下面的联系方式联系作者，或进入开发者社区交流。

## 技术栈

- **Go 语言**：热门后端开发语言，适合构建高并发应用。
- **Echo 框架**：高性能的 Web 框架，支持快速开发和灵活的路由管理。
- **数据库**：开源的关系型数据库，支持 Postgres、MySQL 和 SQLite。
- **Redis**：热门缓存解决方案，提供快速数据存取和持久化选项。
- **JWT**：安全的用户身份验证机制，确保数据传输的完整性和安全性。
- **Docker**：容器化部署工具，简化应用的打包和分发流程。
- **前端**：react + ts + shadcn/ui + tailwindcss4。

## 功能模块

- **账户模块**：实现 JWT 身份验证，支持用户登录、注册、注销、密码修改和个人信息更新。
- **权限模块**：实现 RBAC（Role-Based Access Control）角色权限管理，支持用户-角色-权限的增删改查。
  - 基本功能已实现，考虑到用户使用的不友好性和复杂性，因此暂不推出此功能。
- **文章模块**：提供文章的创建、查看、更新和删除功能。
- **分类模块**：支持类目树及子类目树递归查询，单一类目查询，以及类目的创建、更新和删除。
- **评论模块**：提供评论的创建、查看、删除和回复功能，支持评论树结构的展示。
- **插件系统**：正在火热开发中，即将推出...
- **其他功能**：
  - 提供 snowflake 算法生成唯一 ID
  - 提供 OpenAPI 接口文档
  - 集成 Air 实现热重载
  - 提供 Logrus 实现日志记录
  - 支持 CORS 跨域请求
  - 提供 CSRF 和 XSS 防护
  - 支持 Markdown 的服务端渲染
  - 集成图形验证码功能
  - 支持 QQ/Gmail/Outlook 等主流邮箱服务端发送能力
  - **其他模块正在开发中**，欢迎提供宝贵意见和建议！

## 本地开发

1. **安装依赖**：

   ```bash
   # 安装 swagger 工具
   go install github.com/swaggo/swag/cmd/swag@latest

   # 安装依赖包
   go mod tidy
   ```

2. **配置数据库和邮箱**：  
    修改 `configs/config.yaml` 文件中的数据库配置和邮箱配置，示例如下：

   ```yaml
   database:
     DB_DIALECT: "postgres" # 数据库类型, 可选值: postgres, mysql, sqlite
     DB_NAME: "jank_db"
     DB_HOST: "127.0.0.1" # 如果使用docker，则改为"postgres_db"
     DB_PORT: "5432"
     DB_USER: "<DATABASE_USER>"
     DB_PSW: "<DATABASE_PASSWORD>"
     DB_PATH: "./database" # SQLite 数据库文件路径

   # 邮箱类型和 SMTP 授权码（可选）
   EMAIL_TYPE: "qq" # 邮箱类型，可选值: qq, gmail, outlook
   FROM_EMAIL: "<FROM_EMAIL>" # 发件人邮箱
   EMAIL_SMTP: "<EMAIL_SMTP>" # SMTP 授权码

   # Swagger 相关
   swagger:
    SWAGGER_HOST: "localhost:9010" # 启用后访问 http://localhost:9010/swagger/index.html 页面获取 Swagger 可视化文档
    SWAGGER_ENABLED: "true" # 是否启用Swagger，可选值: true, false
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
   本地启动应用后，浏览器访问 [http://localhost:9010/testPing](http://localhost:9010/testPing)

## Docker 容器部署（postgres）

1. 修改 `configs/config.yaml` 文件中的数据库配置和邮箱配置，示例如下：

   ```yaml
   APP_HOST: "0.0.0.0" # 如果使用docker，则改为"0.0.0.0"

   database:
     DB_DIALECT: "postgres" # 数据库类型, 可选值: postgres, mysql, sqlite
     DB_NAME: "jank_db"
     DB_HOST: "postgres_db" # 如果使用docker，则改为"postgres_db"
     DB_PORT: "5432"
     DB_USER: "<DATABASE_USER>"
     DB_PSW: "<DATABASE_PASSWORD>"
     DB_PATH: "./database" # SQLite 数据库文件路径

   # 邮箱类型和 SMTP 授权码（可选）
   EMAIL_TYPE: "qq" # 邮箱类型，可选值: qq, gmail, outlook
   FROM_EMAIL: "<FROM_EMAIL>" # 发件人邮箱
   EMAIL_SMTP: "<EMAIL_SMTP>" # SMTP 授权码
   ```

2. 修改 `docker-compose.yaml` 文件中的环境变量，示例如下：

   ```yaml
   environment:
     - POSTGRES_USER=<DATABASE_USER>
     - POSTGRES_PASSWORD=<DATABASE_PASSWORD>
   ```

3. 启动容器：

   ```bash
   docker-compose up -d
   ```

## 接口文档

1. **本地启动查看 swagger 文档**：本地启动应用后，通过浏览器访问 [http://localhost:9010/swagger/index.html](http://localhost:9010/swagger/index.html)

2. **README.md 文档**：在 `docs` 目录下，打开 `README.md` 文件查看。

3. **postman 文档**：在 `docs` 目录下，导入 `docs/Jank_blog.postman_collection.json` 至 Postman 查看。

## roadmap（船新推出）

![image.png](https://s2.loli.net/2025/03/09/qJrtOeFvD95PV4Y.png)

> 注：黑色为已完成部分，白色色为待完成部分。

## 官方社区

最新开发进度在 [Jank-Community](https://github.com/Jank-Community)，如果有任何疑问或建议，欢迎加入微信官方社区讨论。

<img src="https://s2.loli.net/2025/05/10/kLbM4tWTSe2fjXg.jpg" alt="官方社区" width="300" />

> 注：因社群成员较多，请自觉遵守规范。  
> 严禁讨论涉黄、赌、毒及政治敏感内容，禁止发布任何形式的不良广告。

## 贡献者名单

<a href="https://github.com/Done-0/Jank/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Done-0/Jank" alt="贡献者名单" />
</a>

## 特别鸣谢

感谢各位对开源社区的支持，在此诚挚地对每一位赞助者表示感谢！

<p>
  <a href="https://github.com/vxincode">
    <img src="https://github.com/vxincode.png" width="70" height="70" style="border-radius: 50%;" />
  </a>
  <a href="https://github.com/WowDoers">
    <img src="https://github.com/WowDoers.png" width="70" height="70" style="border-radius: 50%;" />
  </a>
</p>

## 联系合作

- **QQ**: 927171598
- **微信**: l927171598
- **邮箱**：<EMAIL>fenderisfine@outlook.com

> 广告、合作或者赞助可根据以上联系方式联系作者。

## 代码统计

<p align="left">
  <img src="https://img.shields.io/github/languages/top/Done-0/Jank?label=主要语言&color=00ADD8" alt="主要语言" />
  <img src="https://img.shields.io/github/languages/code-size/Done-0/Jank?label=代码体积&color=success" alt="代码体积" />
  <img src="https://img.shields.io/github/last-commit/Done-0/Jank?label=最后提交&color=blue" alt="最后提交" />
  <img src="https://img.shields.io/github/commit-activity/m/Done-0/Jank?label=月提交&color=orange" alt="提交频率" />
</p>

### 详细统计
| 语言 | 文件数 | 代码行数 | 注释行数 | 空白行数 | 占比 |
|:----:|:------:|:--------:|:--------:|:--------:|:----:|
| Go | 82 | 3582 | 1619 | 749 | 92.2% |
| Docker | 1 | 18 | 13 | 12 | 0.5% |
| YAML | 4 | 246 | 23 | 33 | 6.3% |
| Markdown | 1 | 1 | 0 | 0 | 0.0% |
| 其他 | 1 | 36 | 0 | 6 | 0.9% |
| **总计** | **89** | **3883** | **1655** | **800** | **100%** |

*注：统计数据由 GitHub Actions 自动更新，最后更新于 2025-05-12*
*排除了 docs、tmp 目录和 go.mod、go.sum、LICENSE、.gitignore、.dockerignore、README.md、README_en.md 文件*
## 许可证

本项目遵循 [MIT 协议](https://opensource.org/licenses/MIT)。

## 增长趋势

<img src="https://api.star-history.com/svg?repos=Done-0/Jank&type=timeline" width="100%" height="65%" alt="GitHub Stats">
