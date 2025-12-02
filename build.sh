#!/bin/bash

# 仅构建Docker镜像（不启动）
# 使用方法: ./build.sh

echo "=========================================="
echo "管理系统 - 构建Docker镜像"
echo "=========================================="

# 检查Docker是否运行
if ! docker info > /dev/null 2>&1; then
    echo "❌ 错误: Docker未运行，请先启动Docker"
    exit 1
fi

# 编译后端代码（如果Go已安装）
if command -v go &> /dev/null; then
    echo "📦 编译后端代码..."
    cd backend
    go mod download
    if go build -o main .; then
        echo "✅ 后端代码编译成功"
    else
        echo "❌ 后端代码编译失败"
        exit 1
    fi
    cd ..
fi

# 构建Docker镜像
echo ""
echo "🔨 构建Docker镜像..."
docker-compose build

echo ""
echo "=========================================="
echo "✅ 构建完成！"
echo "=========================================="
echo "启动服务: docker-compose up -d"
echo "=========================================="

