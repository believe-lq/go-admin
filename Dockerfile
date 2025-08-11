# 构建阶段
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

# 设置Go环境变量 - 禁用CGO进行静态编译
ENV GOPROXY=https://goproxy.io,direct
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

# 使用国内源安装基本工具
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update --no-cache && \
    apk add --no-cache git ca-certificates

# 复制依赖文件
COPY go.mod go.sum ./

# 显示构建环境信息
RUN go version && \
    go env GOPROXY && \
    go env GOSUMDB

# 尝试下载依赖，如果失败则使用tidy
RUN go mod download || (echo "Download failed, trying go mod tidy..." && go mod tidy && go mod download)

# 验证依赖
RUN go mod verify

# 复制源代码
COPY . .

# 显示文件结构
RUN ls -la && echo "=== Source files ===" && find . -name "*.go" | head -10

# 静态编译应用（添加详细错误信息）
RUN CGO_ENABLED=0 go build -v -ldflags="-w -s" -o go-admin ./main.go 2>&1 || \
    (echo "Build failed, trying to get more info..." && \
     go build -v ./main.go 2>&1 && exit 1)

# 运行阶段
FROM alpine:3.18

# 设置时区
ENV TZ=Asia/Shanghai

# 安装运行时依赖（只需要基本工具）
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update --no-cache && \
    apk add --no-cache ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 复制编译结果
COPY --from=builder /app/go-admin /main

# 复制配置文件
COPY ./config/settings.demo.yml /config/settings.yml

# 暴露端口
EXPOSE 8000

# 启动命令
CMD ["/main", "server", "-c", "/config/settings.yml"]