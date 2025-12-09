#!/bin/bash

# 一键编译部署重启脚本
# 使用方法: ./deploy.sh

set -e  # 遇到错误立即退出

echo "=========================================="
echo "管理系统 - 一键部署脚本"
echo "=========================================="

# 检查Docker是否运行
if ! docker info > /dev/null 2>&1; then
    echo "❌ 错误: Docker未运行，请先启动Docker"
    exit 1
fi

# 检查docker-compose是否安装
if command -v docker-compose &> /dev/null; then
    COMPOSE_CMD="docker-compose"
elif docker compose version &> /dev/null; then
    COMPOSE_CMD="docker compose"
else
    echo "❌ 错误: docker-compose（或 docker compose） 未安装"
    exit 1
fi

# 1. 编译后端代码
echo ""
echo "📦 步骤 1/4: 编译后端代码..."
cd backend

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "⚠️  警告: Go未安装，跳过本地编译，将直接使用Docker构建"
    cd ..
else
    echo "正在下载依赖..."
    go env -w GOPROXY=https://goproxy.cn,direct
    go env -w GOSUMDB=sum.golang.google.cn

    go mod download
    
    echo "正在编译..."
    if go build -o main .; then
        echo "✅ 后端代码编译成功"
    else
        echo "❌ 后端代码编译失败"
        exit 1
    fi
    cd ..
fi

# 2. 停止旧服务
echo ""
echo "🛑 步骤 2/4: 停止旧服务..."
$COMPOSE_CMD down

# 3. 构建并启动服务
echo ""
echo "🚀 步骤 3/4: 构建Docker镜像并启动服务..."
$COMPOSE_CMD up -d --build

# 4. 等待服务启动
echo ""
echo "⏳ 步骤 4/4: 等待服务启动..."
sleep 8

# 检查服务状态
echo ""
echo "📊 服务状态:"
$COMPOSE_CMD ps

# 检查后端服务健康状态
echo ""
echo "🔍 检查后端服务..."
max_attempts=10
attempt=0
while [ $attempt -lt $max_attempts ]; do
    if curl -s http://localhost:7701/api/auth/userinfo > /dev/null 2>&1; then
        echo "✅ 后端服务运行正常"
        break
    else
        attempt=$((attempt + 1))
        if [ $attempt -eq $max_attempts ]; then
            echo "⚠️  后端服务可能还在启动中，请稍后检查"
        else
            sleep 2
        fi
    fi
done

echo ""
echo "=========================================="
echo "✅ 部署完成！"
echo "=========================================="
echo "📌 服务地址:"
echo "   后端API: http://localhost:7701"
echo "   MySQL:   localhost:3306"
echo "   Redis:   localhost:6379"
echo ""
echo "📝 常用命令:"
echo "   查看日志:    docker-compose logs -f backend"
echo "   查看所有日志: docker-compose logs -f"
echo "   停止服务:    docker-compose down"
echo "   重启服务:    docker-compose restart"
echo "=========================================="

