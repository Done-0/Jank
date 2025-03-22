<p style="text-align: center;">
  <a><img src="https://s2.loli.net/2025/03/14/BnchjpPLeIaoO75.png" alt="Jank"></a>
</p>
<p style="text-align: center;">
  <em>Jank, a lightweight blogging system, developed using Go language and the Echo framework, emphasizes minimalism, low coupling, and high scalability.</em>
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
  <span style="text-decoration: underline; color: grey;">Simplified Chinese</span> | <a href="README_en.md" style="text-decoration: none;">English</a>
</p>

---

Jank is a lightweight blogging system, developed using Go language and the Echo framework, designed with a focus on minimalism, low coupling, and high scalability. It aims to provide users with a feature-rich, simple, and secure blogging experience.

> Note: This project currently lacks a front-end component. We invite developers interested in front-end development to join and contribute to the project. We look forward to your valuable feedback and contributions!

## Quick Overview

👉 Demo Site: [https://fenderisfine.icu](https://fenderisfine.icu)

👉 [【Jank Blog System】New Technology Stack and UI](https://www.bilibili.com/video/BV1bjQ8YNEEo/?share_source=copy_web&vd_source=6fd45877cd498bfb9c2b449d1197363c)

👉 Front-end Repository: [https://github.com/Done-0/Jank-website](https://github.com/Done-0/Jank-website)

![home-page.png](https://s2.loli.net/2025/03/18/CVYwRJOaXtH4nb8.png)
![posts-page.png](https://s2.loli.net/2025/03/18/s6WH3BVmlbyarRS.png)
![post1-page.png](https://s2.loli.net/2025/03/18/IEpPOhcfRdKbk4J.png)
![post2-page.png](https://s2.loli.net/2025/03/18/TS1j9Zr7UpnVPOY.png)

> Note: Since the project is in its early stages, some configuration files may need to be adjusted according to actual requirements. Please contact the author using the information below, or join the developer community for further discussions.

## Technology Stack

- **Go Language**: A popular back-end development language, suitable for building high-concurrency applications.
- **Echo Framework**: A high-performance web framework that supports rapid development and flexible routing management.
- **PostgreSQL**: An open-source relational database offering high performance and reliability for data storage.
- **Redis**: A popular caching solution that offers fast data access and persistence options.
- **JWT**: A secure user authentication mechanism to ensure the integrity and security of data transmission.
- **Docker**: A containerization deployment tool that simplifies application packaging and distribution.
- **Frontend**: React + Umi + Shadcn/ui + TailwindCSS.

## Feature Modules

- **Account Module**: Implements JWT authentication, supporting user login, registration, logout, password modification, and personal information updates.
- **Permission Module**: Implements RBAC (Role-Based Access Control) for role permission management, supporting CRUD operations for users, roles, and permissions.
  - Basic features are implemented, but due to the complexity and user-friendliness concerns, this feature is not yet released.
- **Article Module**: Provides functionality for creating, viewing, updating, and deleting articles.
- **Category Module**: Supports recursive queries for category trees and subcategory trees, single category queries, and category creation, updating, and deletion.
- **Comment Module**: Provides functionality for creating, viewing, deleting, and replying to comments, supporting tree structures for comment display.
- **Plugin System**: Actively under development, coming soon...
- **Other Features**:
  - Provides OpenAPI documentation
  - Integrates Air for hot reloading
  - Uses Logrus for logging
  - Supports CORS for cross-origin requests
  - Provides CSRF and XSS protection
  - Supports server-side rendering for Markdown
  - **Other modules are under development**, feedback and suggestions are welcome!

## Local Development

1. **Install Dependencies**:

   ```bash
   # Install swagger tool
   go install github.com/swaggo/swag/cmd/swag@latest

   # Install dependencies
   go mod tidy
   ```

2. **Configure Database and Email**:  
   Modify the database and email configurations in the `configs/config.yaml` file as shown below:

   ```yaml
   database:
     DB_DIALECT: "postgres" # Database type, options: postgres, mysql, sqlite
     DB_NAME: "jank_db"
     DB_HOST: "127.0.0.1" # Change to "postgres_db" if using Docker
     DB_PORT: "5432"
     DB_USER: "<DATABASE_USER>"
     DB_PSW: "<DATABASE_PASSWORD>"
     DB_PATH: "./database" # SQLite database file path

   # Email type and SMTP authorization code (optional)
   EMAIL_TYPE: "qq" # Email type, options: qq, gmail, outlook
   FROM_EMAIL: "<FROM_EMAIL>" # Sender email
   EMAIL_SMTP: "<EMAIL_SMTP>" # SMTP authorization code
   ```

3. **Start the Service**:  
   Use the following command to start the application:

   ```bash
   go run main.go
   ```

   Or use Air for hot reloading:

   > This method is the most convenient, but make sure to configure the GOPATH environment variable beforehand.

   ```bash
   # Install Air (requires Go 1.22 or higher)
   go install github.com/air-verse/air@latest

   # Start with hot reloading
   air -c ./configs/.air.toml
   ```

4. **Access the Interface**:  
   After starting the application locally, access it via the browser at [http://localhost:9010/ping](http://localhost:9010/ping)

## Docker Container Deployment (PostgreSQL)

1. Modify the database and email configurations in the `configs/config.yaml` file as shown below:

   ```yaml
   APP_HOST: "0.0.0.0" # Change to "0.0.0.0" if using Docker

   database:
     DB_DIALECT: "postgres" # Database type, options: postgres, mysql, sqlite
     DB_NAME: "jank_db"
     DB_HOST: "postgres_db" # Change to "postgres_db" if using Docker
     DB_PORT: "5432"
     DB_USER: "<DATABASE_USER>"
     DB_PSW: "<DATABASE_PASSWORD>"
     DB_PATH: "./database" # SQLite database file path

   # Email type and SMTP authorization code (optional)
   EMAIL_TYPE: "qq" # Email type, options: qq, gmail, outlook
   FROM_EMAIL: "<FROM_EMAIL>" # Sender email
   EMAIL_SMTP: "<EMAIL_SMTP>" # SMTP authorization code
   ```

2. Modify the environment variables in the `docker-compose.yaml` file as shown below:

   ```yaml
   environment:
     - POSTGRES_USER=<DATABASE_USER>
     - POSTGRES_PASSWORD=<DATABASE_PASSWORD>
   ```

3. Start the container:

   ```bash
   docker-compose up -d
   ```

## API Documentation

1. **View Swagger Documentation Locally**: After starting the application locally, access the Swagger documentation at [http://localhost:9010/swagger/index.html](http://localhost:9010/swagger/index.html)

2. **README.md Documentation**: Open the `README.md` file in the `docs` directory.

3. **Postman Documentation**: Import the `docs/Jank_blog.postman_collection.json` file into Postman to view.

## Roadmap (New Release)

![image.png](https://s2.loli.net/2025/03/09/qJrtOeFvD95PV4Y.png)

> Note: Black represents completed features, white represents pending features.

## Architecture Diagram (To Be Updated)

**Architecture Diagram and Visual API Documentation**: Open the `docs/jank_blog_architecture.drawio` file in the project root directory.

> Note: This document was created using `draw.io` and requires the [draw.io](https://app.diagrams.net/) tool to view.

## Official Community

If you have any questions or suggestions, feel free to join the official community for discussions.

<img src="https://s2.loli.net/2025/01/25/L9BspuHnrIeim7S.jpg" alt="Official Community" width="300" />

## Special Thanks

We sincerely thank every sponsor for their support of the open-source community!

<p>
  <a href="https://github.com/vxincode">
    <img src="https://github.com/vxincode.png" width="80" height="80" style="border-radius: 50%;" />
  </a>
  <a href="https://github.com/WowDoers">
    <img src="https://github.com/WowDoers.png" width="80" height="80" style="border-radius: 50%;" />
  </a>
</p>

## Contact and Cooperation

- **QQ**: 927171598
- **Email**: <EMAIL>fenderisfine@outlook.com
- **Developer Community (QQ)**: 828270460

## Contributors List

<a href="https://github.com/Done-0/Jank/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Done-0/Jank" alt="Contributors List" />
</a>

## Code Statistics

<p align="left">
  <img src="https://img.shields.io/github/languages/top/Done-0/Jank?label=Primary%20Language&color=00ADD8" alt="Primary Language" />
  <img src="https://img.shields.io/github/languages/code-size/Done-0/Jank?label=Code%20Size&color=success" alt="Code Size" />
  <img src="https://img.shields.io/github/last-commit/Done-0/Jank?label=Last%20Commit&color=blue" alt="Last Commit" />
  <img src="https://img.shields.io/github/commit-activity/m/Done-0/Jank?label=Monthly%20Commits&color=orange" alt="Commit Frequency" />
</p>

### Detailed Statistics
| Language | Files | Code Lines | Comment Lines | Blank Lines | Percentage |
|:--------:|:-----:|:----------:|:-------------:|:-----------:|:----------:|
| Go | 82 | 3342 | 664 | 669 | 92.3% |
| Docker | 1 | 16 | 14 | 13 | 0.4% |
| YAML | 3 | 227 | 21 | 32 | 6.3% |
| Markdown | 1 | 1 | 0 | 0 | 0.0% |
| Others | 1 | 36 | 0 | 6 | 1.0% |
| **Total** | **88** | **3622** | **699** | **720** | **100%** |

*Note: Statistics are automatically updated by GitHub Actions, last updated on 2025-03-22*
*Excluded docs, tmp directories and go.mod, go.sum, LICENSE, .gitignore, .dockerignore, README.md, README_en.md files*
## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Growth Trends

<img src="https://api.star-history.com/svg?repos=Done-0/Jank&type=timeline" width="100%" height="65%" alt="GitHub Stats">