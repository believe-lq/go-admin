# 构建阶段
FROM golang:alpine AS builder

# 设置工作目录
WORKDIR /app

# 设置Go环境变量
ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# 安装必要的构建工具
RUN apk add --no-cache gcc g++ libc6-compat

# 复制go.mod文件
COPY go.mod ./

# 执行go mod tidy下载依赖并生成go.sum
RUN go mod tidy

# 复制源代码
COPY . .

# 编译应用
RUN go build -ldflags="-w -s" -o go-admin .

# 运行阶段
FROM alpine

# 安装运行时依赖
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates tzdata
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