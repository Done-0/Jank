<p style="text-align: center;">
  <a><img src="https://p.ipic.vip/6idwb0.PNG" alt="Jank"></a>
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
<p style="text-align: center; margin: 0; padding: 0; position: relative; top: -5px;">
  <span style="text-decoration: underline; color: grey;">English</span> | <a href="README.md" style="text-decoration: none;">简体中文</a>
</p>


---

Jank is a lightweight blogging system developed using Go and the Echo framework, with a design philosophy that emphasizes minimalism, low coupling, and high extensibility. It aims to provide users with a rich, clean, simple, and secure blogging experience.

> The project is currently lacking front-end development. We welcome developers who are interested to contact me for collaboration. We look forward to your valuable input!

## Tech Stack

- **Go**: A popular backend programming language suitable for building high-concurrency applications.
- **Echo Framework**: A high-performance web framework that supports rapid development and flexible routing management.
- **MySQL**: A mature relational database management system that supports complex queries and transaction processing.
- **Redis**: A popular caching solution offering fast data access and persistence options.
- **JWT**: A secure user authentication mechanism ensuring data integrity and security during transmission.
- **Docker**: A containerization tool that simplifies packaging and distributing applications.
- **Frontend**: Vue 3 + Nuxt + Shadcn-vue (The original project is no longer maintained; contributions from frontend experts are welcome).

## Features

- **Account Module**: Implements JWT authentication, supporting user login, registration, logout, password changes, and personal information updates.
- **Article Module**: Provides features for creating, viewing, updating, and deleting articles.
- **Category Module**: Supports recursive category tree queries, single category queries, as well as category creation, update, and deletion.
- **Comment Module**: Provides functionality for comment creation, viewing, deletion, and replies, supporting the display of comment tree structures.
- **Other Features**:
  - Provides OpenAPI documentation.
  - Integrates Air for hot reloading.
  - Uses Logrus for logging.
  - Supports CORS for cross-origin requests.
  - Provides CSRF and XSS protection.
  - Supports server-side rendering of Markdown.

> **Other modules are under development**, and suggestions and feedback are welcome!

## Local Development

1. **Install Dependencies**:

   ```bash
   # install swagger utils
   go install github.com/swaggo/swag/cmd/swag@latest
   
   # install dependencies
   go mod tidy
   ```

2. **Configure Database and Email**:
   Modify the database and email configurations in the `configs/config.yaml` file as shown below:

   ```yaml
   # MySQL Database password (required)
   DB_PSW: "<DATABASE_PASSWORD>"
   
   # QQ email and SMTP authorization code (optional)
   QQ_SMTP: "<QQ_SMTP>"
   FROM_EMAIL: "<FROM_QQ_EMAIL>"
   ```

3. **Start the Service**:
   Use the following command to start the application:

   ```bash
   go run main.go
   ```

   Or use Air for hot reloading:

   ```bash
   # install air, requires go 1.22+
   go install github.com/air-verse/air@latest
   
   air -c ./configs/.air.toml
   ```

4. **Access the API**:
   Open your browser and go to [http://localhost:9010/ping](http://localhost:9010/ping) when you start the system locally.

## Docker Deployment

```bash
docker-compose up -d
```

## API Documentation
1. **Local Swagger Documentation**: After launching the application locally, visit [http://localhost:9010/swagger/index.html](http://localhost:9010/swagger/index.html) in your browser.
2. **Postman Documentation**: In the project root directory, import the `docs/Jank_blog.postman_collection.json` file into Postman to view.

## Architecture Diagram
Architecture Diagram and Visualized API Documentation: Open the `docs/jank_blog_architecture.drawio` file in the project root directory.

> note: This document is drawn by draw.io, which requires the installation of the draw.io software to view.

## Official Community

If you have any questions or suggestions, feel free to join the official community for discussion.

<img src="https://s2.loli.net/2025/01/25/L9BspuHnrIeim7S.jpg" alt="Official Community" width="300" />

## Contact Information

- **QQ**: 927171598
- **Email**: <EMAIL>fenderisfine@outlook.com

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## GitHub Statistics

<img src="https://github-readme-stats.vercel.app/api?username=Done-0&show_icons=true&hide_title=true&theme=radical" width="100%" height="65%">

## Growth Trends

<img src="https://api.star-history.com/svg?repos=Done-0/Jank&type=timeline" width="100%" height="65%">
