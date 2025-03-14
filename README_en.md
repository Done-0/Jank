<p style="text-align: center;">
  <a><img src="https://s2.loli.net/2025/03/14/BnchjpPLeIaoO75.png" alt="Jank"></a>
</p>
<p style="text-align: center;">
  <em>Jank, a lightweight blogging system developed with Go and the Echo framework, emphasizing minimalism, low coupling, and high extensibility</em>
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
  <span style="text-decoration: underline; color: grey;">English</span> | <a href="README.md" style="text-decoration: none;">简体中文</a>
</p>

---

Jank is a lightweight blogging system developed using Go and the Echo framework. Its design philosophy emphasizes minimalism, low coupling, and high extensibility, aiming to provide users with a feature-rich, clean, simple, and secure blogging experience.

> **Note:** The current project lacks a front-end implementation. We warmly invite developers interested in front-end development to join and collaborate. Your valuable suggestions and contributions are greatly appreciated!

## Quick Overview

- **Demo Site:** [https://fenderisfine.icu](https://fenderisfine.icu)
- **Bilibili Preview Video:** [Have You Seen a Blogging System Developed in Go?](https://www.bilibili.com/video/BV1W5wdeZEoY/?share_source=copy_web&vd_source=6fd45877cd498bfb9c2b449d1197363c)
- **Deployment Documentation:** [Docker Deployment Guide](https://fenderisfine.icu/posts/2)
- **Front-end Repository:** [https://github.com/Done-0/Jank-website](https://github.com/Done-0/Jank-website)

![home-page.png](https://s2.loli.net/2025/03/14/veaGZ4hwfjpbN9g.png)
![home-white.png](https://raw.githubusercontent.com/Done-0/Jank-website/main/public/images/home-white.png)
![home-black.png](https://raw.githubusercontent.com/Done-0/Jank-website/main/public/images/home-black.png)

> **Note:** As the project is still in its early stages, some configuration files might require adjustments according to your specific environment. Please use the contact information below or join our developer community for further discussion.

## Tech Stack

- **Go:** A popular backend programming language ideal for building high-concurrency applications.
- **Echo Framework:** A high-performance web framework that supports rapid development and flexible routing.
- **PostgreSQL:** An open-source relational database offering high performance and reliability.
- **Redis:** A widely used caching solution that provides fast data access and persistence options.
- **JWT:** A secure authentication mechanism ensuring data integrity and safety during transmission.
- **Docker:** A containerization tool that simplifies application packaging and distribution.
- **Frontend:** react + umi + shadcn/ui + tailwindcss.

## Feature Modules

- **Account Module:**  
  Implements JWT-based authentication, supporting user login, registration, logout, password changes, and updates to personal information.

- **Permission Module:**  
  Implements RBAC (Role-Based Access Control) for managing user roles and permissions, including CRUD operations for users, roles, and permissions.  
  _Note: Basic functionality has been implemented, but due to complexity and potential usability issues, this feature will not be included in the first release._

- **Article Module:**  
  Provides functionalities for creating, viewing, updating, and deleting articles.

- **Category Module:**  
  Supports recursive queries for category trees and subcategories, as well as single category queries and operations for creating, updating, and deleting categories.

- **Comment Module:**  
  Offers features for creating, viewing, deleting, and replying to comments, and supports displaying comments in a tree structure.

- **Plugin System:**  
  Under active development and coming soon…

- **Other Features:**
  - Provides OpenAPI documentation.
  - Integrates Air for hot reloading.
  - Utilizes Logrus for logging.
  - Supports CORS for cross-origin requests.
  - Includes CSRF and XSS protection.
  - Supports server-side Markdown rendering.
  - **Additional modules are under development**—your feedback and suggestions are welcome!

## Local Development

1. **Install Dependencies:**

   ```bash
   # Install the swagger tool
   go install github.com/swaggo/swag/cmd/swag@latest

   # Install dependency packages
   go mod tidy
   ```

2. **Configure Database and Email:**  
   Modify the `configs/config.yaml` file with your database and email configurations. For example:

   ```yaml
   # PostgreSQL database configuration
   DB_USER: "<DATABASE_USER>"
   DB_PSW: "<DATABASE_PASSWORD>"

   # QQ email and SMTP authorization code (optional)
   QQ_SMTP: "<QQ_SMTP>"
   FROM_EMAIL: "<FROM_QQ_EMAIL>"
   ```

3. **Start the Service:**  
   Run the following command to start the application:

   ```bash
   go run main.go
   ```

   Or use Air for hot reloading:

   > This method is very convenient but requires that your GOPATH environment variable is configured in advance.

   ```bash
   # Install Air (requires Go 1.22 or later)
   go install github.com/air-verse/air@latest

   # Start with hot reloading
   air -c ./configs/.air.toml
   ```

4. **Access the API:**  
   Once the application is running locally, open your browser and visit [http://localhost:9010/ping](http://localhost:9010/ping).

## Docker Container Deployment

1. **Update Configuration:**  
   Modify the database and email settings in the `configs/config.yaml` file. For example:

   ```yaml
   # Application configuration
   APP_HOST: "0.0.0.0"

   # PostgreSQL database configuration
   DB_USER: "<DATABASE_USER>"
   DB_PSW: "<DATABASE_PASSWORD>"

   # QQ email and SMTP authorization code (optional)
   QQ_SMTP: "<QQ_SMTP>"
   FROM_EMAIL: "<FROM_QQ_EMAIL>"
   ```

2. **Adjust Environment Variables:**  
   In the `docker-compose.yaml` file, update the environment variables as shown:

   ```yaml
   environment:
     - POSTGRES_USER=<DATABASE_USER>
     - POSTGRES_PASSWORD=<DATABASE_PASSWORD>
   ```

3. **Start the Containers:**

   ```bash
   docker-compose up -d
   ```

## API Documentation

1. **Local Swagger Documentation:**  
   After launching the application locally, visit [http://localhost:9010/swagger/index.html](http://localhost:9010/swagger/index.html) in your browser.

2. **README Documentation:**  
   Open the `README.md` file located in the `docs` directory for more details.

3. **Postman Documentation:**  
   Import the `docs/Jank_blog.postman_collection.json` file into Postman to view the API details.

## Roadmap (Newly Launched)

![image.png](https://s2.loli.net/2025/03/09/qJrtOeFvD95PV4Y.png)

> **Note:** The black areas represent completed features, while the white areas indicate features that are pending.

## Architecture Diagram (To Be Updated)

**Architecture Diagram and Visual API Documentation:**  
Open the `docs/jank_blog_architecture.drawio` file located in the project root.

> **Note:** This document was created using [draw.io](https://app.diagrams.net/). You will need the draw.io tool to view it.

## Official Community

If you have any questions or suggestions, feel free to join our official community for discussion.

<img src="https://s2.loli.net/2025/01/25/L9BspuHnrIeim7S.jpg" alt="Official Community" width="300" />

## Special Thanks

Many thanks to everyone for your support!

<p>
  <a href="https://github.com/vxincode">
    <img src="https://github.com/vxincode.png" width="80" height="80" style="border-radius: 50%;" />
  </a>
  <a href="https://github.com/WowDoers">
    <img src="https://github.com/WowDoers.png" width="80" height="80" style="border-radius: 50%;" />
  </a>
</p>

## Contact & Collaboration

- **QQ:** 927171598
- **Email:** fenderisfine@outlook.com
- **Developer Community (QQ):** 828270460

## Contributors

<a href="https://github.com/Done-0/Jank/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Done-0/Jank" alt="贡献者名单" />
</a>

## Code Statistics

<p align="left">
  <img src="https://img.shields.io/github/languages/top/Done-0/Jank?label=Language&color=00ADD8" alt="Language" />
  <img src="https://img.shields.io/github/languages/code-size/Done-0/Jank?label=Code%20Size&color=success" alt="Code Size" />
  <img src="https://img.shields.io/github/last-commit/Done-0/Jank?label=Last%20Commit&color=blue" alt="Last Commit" />
  <img src="https://img.shields.io/github/commit-activity/m/Done-0/Jank?label=Monthly%20Commits&color=orange" alt="Monthly Commits" />
</p>

### Detailed Statistics
| Language | Files | Code Lines | Comment Lines | Blank Lines | Percentage |
|:--------:|:-----:|:----------:|:-------------:|:-----------:|:----------:|
| Go | 100 | 3989 | 917 | 820 | 93.9% |
| Docker | 1 | 16 | 14 | 13 | 0.4% |
| YAML | 3 | 206 | 21 | 31 | 4.8% |
| Markdown | 1 | 1 | 0 | 0 | 0.0% |
| Others | 1 | 36 | 0 | 6 | 0.8% |
| **Total** | **106** | **4248** | **952** | **870** | **100%** |

*Note: Statistics are automatically updated by GitHub Actions, last updated on 2025-03-14*
*Excluded docs, tmp directories and go.mod, go.sum, LICENSE, .gitignore, .dockerignore, README.md, README_en.md files*
## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Growth Trends

<img src="https://api.star-history.com/svg?repos=Done-0/Jank&type=timeline" width="100%" height="65%" alt="GitHub Growth Trends">
