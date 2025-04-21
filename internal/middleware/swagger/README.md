swagger API 文档组件

- 启动后，默认访问地址：http://localhost:9010/swagger/index.html

> 如果不想使用 swagger，前往 `internal/middleware/middleware.go` 中注释掉 `app.Use(swagger_middleware.InitSwagger())` 函数即可。
