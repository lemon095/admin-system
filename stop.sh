#!/bin/bash

# 停止所有服务
# 使用方法: ./stop.sh

echo "=========================================="
echo "管理系统 - 停止服务"
echo "=========================================="

# 检查Docker是否运行
if ! docker info > /dev/null 2>&1; then
    echo "❌ 错误: Docker未运行"
    exit 1
fi

# 检查docker-compose是否安装，兼容新旧版本
if command -v docker-compose &> /dev/null; then
    COMPOSE_CMD="docker-compose"
elif docker compose version &> /dev/null 2>&1; then
    COMPOSE_CMD="docker compose"
else
    echo "❌ 错误: docker-compose（或 docker compose）未安装"
    exit 1
fi

echo "🛑 正在停止服务..."
$COMPOSE_CMD down

echo ""
echo "=========================================="
echo "✅ 服务已停止"
echo "=========================================="
echo "启动服务: ./start.sh 或 $COMPOSE_CMD up -d"
echo "=========================================="

