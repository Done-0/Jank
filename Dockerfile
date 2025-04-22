# 第一阶段：构建阶段
FROM golang:1.23.0 AS builder

# 设置环境变量
ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /jank

# 复制 go.mod 和 go.sum，并下载依赖
COPY go.mod go.sum ./
RUN go mod download && go mod tidy

# 安装 swag 工具
RUN go install github.com/swaggo/swag/cmd/swag@latest

# 复制项目源码到容器中
COPY . .

# 构建 Go 应用
RUN go build -o main .

# 第二阶段：生产镜像
FROM alpine:3.18

# 安装基础依赖并设置时区
RUN apk --no-cache add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 设置工作目录
WORKDIR /app

# 从构建阶段复制必要文件
COPY --from=builder /jank/main .
COPY --from=builder /go/bin/swag /usr/local/bin/swag
COPY --from=builder /jank/pkg /app/pkg

# 开放端口 9010
EXPOSE 9010

# 启动命令
CMD ["./main"]
