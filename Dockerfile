# 构建阶段
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

# 设置Go环境变量 - 使用国内代理
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=sum.golang.google.cn
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

# 使用国内Alpine镜像源并安装必要的构建工具和sqlite3依赖
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update --no-cache && \
    apk add --no-cache gcc g++ libc6-compat sqlite-dev

# 复制go.mod和go.sum文件
COPY go.mod go.sum* ./

# 执行go mod tidy下载依赖并生成go.sum
RUN go mod download && go mod tidy

# 复制源代码
COPY . .

# 显示Go版本和模块信息
RUN go version && go env

# 编译应用（添加详细输出）
RUN go build -v -ldflags="-w -s" -o go-admin . || (echo "Build failed, checking for errors..." && exit 1)

# 运行阶段
FROM alpine:latest

# 安装运行时依赖
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update --no-cache && \
    apk add --no-cache ca-certificates tzdata sqlite
ENV TZ Asia/Shanghai

# 从构建阶段复制编译好的应用
COPY --from=builder /app/go-admin /main

# 复制配置文件
COPY ./config/settings.demo.yml /config/settings.yml

# 复制数据库文件（如果需要）
COPY ./go-admin-db.db /go-admin-db.db

# 暴露端口
EXPOSE 8000

# 给可执行文件添加执行权限
RUN chmod +x /main

# 启动命令
CMD ["/main","server","-c", "/config/settings.yml"]