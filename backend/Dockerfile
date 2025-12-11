# 构建阶段
# 如果无法访问 Docker Hub，请先配置 Docker 镜像加速器
# 或者使用 Dockerfile.aliyun（如果服务器在阿里云）
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /build

# 安装必要的构建工具
RUN apk add --no-cache git make

# 设置Go环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GOSUMDB=sum.golang.google.cn \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 复制go mod文件
COPY go.mod go.sum ./
# 下载依赖（利用Docker缓存层）
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN go build -ldflags="-w -s" -o server main.go

# 确保 resource 目录存在（如果不存在则创建空目录）
RUN mkdir -p resource

# 运行阶段
# 如果无法访问 Docker Hub，请先配置 Docker 镜像加速器
# 或者使用 Dockerfile.aliyun（如果服务器在阿里云）
FROM alpine:latest

# 维护者信息
LABEL MAINTAINER="admin-system"

# 设置时区
ENV TZ=Asia/Shanghai \
    LANG=C.UTF-8

# 安装必要的运行时依赖
RUN apk add --no-cache \
    tzdata \
    ca-certificates \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && update-ca-certificates

# 创建非root用户
RUN addgroup -g 1000 appgroup && \
    adduser -D -u 1000 -G appgroup appuser

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder --chown=appuser:appgroup /build/server ./

# 复制资源文件（resource 目录在构建阶段已确保存在）
COPY --from=builder --chown=appuser:appgroup /build/resource ./resource/

# 切换到非root用户
USER appuser

# 暴露端口（默认7701，可通过环境变量覆盖）
EXPOSE 9001

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=40s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:7701/health || exit 1

# 启动应用
ENTRYPOINT ["./server"]