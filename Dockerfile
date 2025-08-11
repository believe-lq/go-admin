# 构建阶段
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

# 设置Go环境变量 - 禁用CGO进行静态编译
ENV GOPROXY=https://goproxy.io,direct
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# 使用国内源安装基本工具
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update --no-cache && \
    apk add --no-cache git

# 复制依赖文件（go.sum可选）
COPY go.mod ./

# 下载依赖
RUN go mod tidy


# 复制源代码
COPY . .

# 静态编译应用（不需要CGO）
RUN CGO_ENABLED=0 go build -v -ldflags="-w -s" -o go-admin ./main.go

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