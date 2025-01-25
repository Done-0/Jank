# 第一阶段：构建阶段
FROM golang:1.23 AS builder

# 设置环境变量，确保 Go 编译时设置为 Linux 环境，禁用 CGO
ENV CGO_ENABLED=0 GOOS=linux

# 设置工作目录
WORKDIR /jank

# 复制 go.mod 和 go.sum，并下载依赖
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

# 复制项目源码到容器中
COPY . .

# 构建 Go 应用
RUN go build -o main .

# 第二阶段：生产镜像
FROM alpine:latest

# 安装基础依赖，如 CA 证书和时区
RUN apk --no-cache add ca-certificates tzdata

# 设置时区为上海
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的应用
COPY --from=builder /jank/main .

# 开放端口 9010
EXPOSE 9010

# 启动命令
CMD ["./main"]
