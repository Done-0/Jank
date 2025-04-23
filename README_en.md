<p style="text-align: center;">
  <a><img src="https://s2.loli.net/2025/03/14/BnchjpPLeIaoO75.png" alt="Jank"></a>
</p>
<p style="text-align: center;">
  <em>Jank - A lightweight blog system built with Go and Echo framework, emphasizing minimalism, low coupling, and high extensibility</em>
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
  <a href="README.md" style="text-decoration: none;">简体中文</a> | <span style="text-decoration: underline; color: grey;">English</span>
</p>

---

Jank is a lightweight blog system built with Go language and Echo framework. Designed with principles of minimalism, low coupling, and high extensibility, it aims to provide users with a feature-rich, clean-interface, easy-to-use, and secure blogging experience.

> Note: The project currently lacks professional frontend development. We sincerely invite frontend developers to participate in the frontend project development. Your valuable suggestions and contributions are highly appreciated!

## Quick Look

👉 Demo Site: [https://www.jank.org.cn](https://www.jank.org.cn)

👉 [【Jank Blog System】New Tech Stack & UI】](https://www.bilibili.com/video/BV1bjQ8YNEEo/?share_source=copy_web&vd_source=6fd45877cd498bfb9c2b449d1197363c)

👉 Frontend Repository: [https://github.com/Done-0/Jank-website](https://github.com/Done-0/Jank-website)

![home-page.png](https://s2.loli.net/2025/04/07/l1tGYV4WkmoiIHv.png)
![posts-page.png](https://s2.loli.net/2025/04/07/xR62vhWKsmgw3Ht.png)
![post1-page.png](https://s2.loli.net/2025/04/07/DbcJzryKmBNR7vQ.png)
![post2-page.png](https://s2.loli.net/2025/04/07/iNpXyMdkjaDbn92.png)

> Note: As the project is still in its early stages, some configuration files may require adjustments based on actual usage. Please contact the author using the provided contact information or join the developer community for discussions.

## Tech Stack

- **Go Language**: Popular backend development language suitable for building high-concurrency applications.
- **Echo Framework**: High-performance web framework supporting rapid development and flexible route management.
- **Database**: Open-source relational databases supporting Postgres, MySQL, and SQLite.
- **Redis**: Popular caching solution providing fast data access and persistence options.
- **JWT**: Secure user authentication mechanism ensuring data transmission integrity and security.
- **Docker**: Containerization deployment tool simplifying application packaging and distribution.
- **Frontend**: React + Next.js + shadcn/ui + Tailwind CSS.

## Feature Modules

- **Account Module**: JWT authentication supporting user login, registration, logout, password modification, and profile updates.
- **Permission Module**: RBAC (Role-Based Access Control) implementation supporting user-role-permission management.
  - Core functionality implemented but temporarily withheld due to usability considerations.
- **Article Module**: Provides article creation, viewing, updating, and deletion functionalities.
- **Category Module**: Supports recursive queries for category trees and sub-category trees, single category queries, and CRUD operations.
- **Comment Module**: Enables comment creation, viewing, deletion, and replies with tree structure display.
- **Plugin System**: Under active development, coming soon...
- **Other Features**:
  - OpenAPI documentation
  - Air integration for hot reload
  - Logrus implementation for logging
  - CORS support
  - CSRF and XSS protection
  - Server-side Markdown rendering
  - Graphical CAPTCHA integration
  - Email sending support for QQ/Gmail/Outlook
  - **More modules in development** - Welcome your valuable suggestions!

## Local Development

1. **Install Dependencies**:

   ```bash
   # Install swagger tool
   go install github.com/swaggo/swag/cmd/swag@latest

   # Install project dependencies
   go mod tidy
   ```

2. **Configure Database and Email**:
   Modify `configs/config.yaml`:

   ```yaml
   database:
     DB_DIALECT: "postgres" # Database type: postgres, mysql, sqlite
     DB_NAME: "jank_db"
     DB_HOST: "127.0.0.1" # Use "postgres_db" for docker
     DB_PORT: "5432"
     DB_USER: "<DATABASE_USER>"
     DB_PSW: "<DATABASE_PASSWORD>"
     DB_PATH: "./database" # SQLite file path

   # Email configuration
   EMAIL_TYPE: "qq" # qq, gmail, outlook
   FROM_EMAIL: "<FROM_EMAIL>"
   EMAIL_SMTP: "<EMAIL_SMTP>"
   ```

3. **Start Service**:

   ```bash
   go run main.go
   ```

   Or with Air hot reload:

   ```bash
   go install github.com/air-verse/air@latest
   air -c ./configs/.air.toml
   ```

4. **Access API**:
   Visit [http://localhost:9010/ping](http://localhost:9010/ping)

## Docker Deployment (Postgres)

1. Update `configs/config.yaml`:

   ```yaml
   APP_HOST: "0.0.0.0"
   database:
     DB_HOST: "postgres_db"
   ```

2. Modify `docker-compose.yaml` environment variables:

   ```yaml
   environment:
     - POSTGRES_USER=<DATABASE_USER>
     - POSTGRES_PASSWORD=<DATABASE_PASSWORD>
   ```

3. Start containers:
   ```bash
   docker-compose up -d
   ```

## API Documentation

1. **Swagger UI**: Available at [http://localhost:9010/swagger/index.html](http://localhost:9010/swagger/index.html)
2. **Postman Collection**: Import `docs/Jank_blog.postman_collection.json`
3. **README Documentation**: See `docs/README.md`

## Roadmap

![image.png](https://s2.loli.net/2025/03/09/qJrtOeFvD95PV4Y.png)

> Black sections indicate completed features, white sections indicate planned features.

## Official Community

Join our official community for discussions:

<img src="https://s2.loli.net/2025/04/14/qpVkK3Ms9nm6dCj.jpg" alt="Official Community" width="300" />

> Community guidelines:  
> Strictly prohibit discussions involving pornography, gambling, drugs, or politically sensitive content. No advertising permitted.

## Contact

- **QQ**: 927171598
- **Email**: <EMAIL>fenderisfine@outlook.com

## Acknowledgments

Special thanks to our sponsors and contributors:

<p>
  <a href="https://github.com/vxincode">
    <img src="https://github.com/vxincode.png" width="70" height="70" style="border-radius: 50%;" />
  </a>
  <a href="https://github.com/WowDoers">
    <img src="https://github.com/WowDoers.png" width="70" height="70" style="border-radius: 50%;" />
  </a>
</p>

## Contributors

<a href="https://github.com/Done-0/Jank/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Done-0/Jank" alt="Contributors" />
</a>

## Code Statistics

<p align="left">
  <img src="https://img.shields.io/github/languages/top/Done-0/Jank?label=Primary%20Language&color=00ADD8" alt="Primary Language" />
  <img src="https://img.shields.io/github/languages/code-size/Done-0/Jank?label=Code%20Size&color=success" alt="Code Size" />
  <img src="https://img.shields.io/github/last-commit/Done-0/Jank?label=Last%20Commit&color=blue" alt="Last Commit" />
  <img src="https://img.shields.io/github/commit-activity/m/Done-0/Jank?label=Monthly%20Commits&color=orange" alt="Commit Activity" />
</p>

### Detailed Stats

| Language  | Files  | Code Lines | Comments | Blanks  |  Ratio   |
| :-------: | :----: | :--------: | :------: | :-----: | :------: |
|    Go     |   82   |    3343    |   664    |   669   |  92.3%   |
|  Docker   |   1    |     16     |    14    |   13    |   0.4%   |
|   YAML    |   3    |    227     |    21    |   32    |   6.3%   |
| Markdown  |   1    |     1      |    0     |    0    |   0.0%   |
|  Others   |   1    |     36     |    0     |    6    |   1.0%   |
| **Total** | **88** |  **3623**  | **699**  | **720** | **100%** |

_Note: Stats updated automatically via GitHub Actions (Last update: 2025-03-22)_  
_Excludes docs/, tmp/ directories and various config files_

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Star History

<img src="https://api.star-history.com/svg?repos=Done-0/Jank&type=timeline" width="100%" height="65%" alt="GitHub Stats">

### Detailed Statistics
| Language | Files | Code Lines | Comment Lines | Blank Lines | Percentage |
|:--------:|:-----:|:----------:|:-------------:|:-----------:|:----------:|
| Go | 80 | 3478 | 653 | 712 | 92.3% |
| Docker | 1 | 18 | 13 | 12 | 0.5% |
| YAML | 3 | 235 | 21 | 32 | 6.2% |
| Markdown | 1 | 1 | 0 | 0 | 0.0% |
| Others | 1 | 36 | 0 | 6 | 1.0% |
| **Total** | **86** | **3768** | **687** | **762** | **100%** |

*Note: Statistics are automatically updated by GitHub Actions, last updated on 2025-04-23*
*Excluded docs, tmp directories and go.mod, go.sum, LICENSE, .gitignore, .dockerignore, README.md, README_en.md files*
