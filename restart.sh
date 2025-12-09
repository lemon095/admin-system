#!/bin/bash

# 快速重启脚本（不重新构建）
# 使用方法: ./restart.sh

echo "=========================================="
echo "管理系统 - 快速重启"
echo "=========================================="

# 检查Docker是否运行
if ! docker info > /dev/null 2>&1; then
    echo "❌ 错误: Docker未运行，请先启动Docker"
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

echo "🛑 停止服务..."
$COMPOSE_CMD stop

echo "🚀 启动服务..."
$COMPOSE_CMD start

# 等待服务启动
echo "⏳ 等待服务启动..."
sleep 5

echo ""
echo "📊 服务状态:"
$COMPOSE_CMD ps

echo ""
echo "=========================================="
echo "✅ 重启完成！"
echo "=========================================="
echo "查看日志: $COMPOSE_CMD logs -f backend"
echo "=========================================="

