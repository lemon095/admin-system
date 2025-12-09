#!/bin/bash

echo "=========================================="
echo "启动管理系统"
echo "=========================================="

# 检查Docker是否运行
if ! docker info > /dev/null 2>&1; then
    echo "错误: Docker未运行，请先启动Docker"
    exit 1
fi

# 检查docker-compose是否安装，兼容新旧版本
if command -v docker-compose &> /dev/null; then
    COMPOSE_CMD="docker-compose"
elif docker compose version &> /dev/null 2>&1; then
    COMPOSE_CMD="docker compose"
else
    echo "错误: docker-compose（或 docker compose）未安装"
    echo ""
    echo "请安装 docker-compose："
    echo "  CentOS/RHEL: sudo yum install docker-compose"
    echo "  或使用新版本: docker compose（Docker 20.10+ 内置）"
    exit 1
fi

# 启动服务
echo "正在启动服务..."
$COMPOSE_CMD up -d

# 等待服务启动
echo "等待服务启动..."
sleep 5

# 检查服务状态
echo ""
echo "服务状态:"
$COMPOSE_CMD ps

echo ""
echo "=========================================="
echo "服务启动完成！"
echo "=========================================="
echo "后端API: http://localhost:9020"
echo "MySQL: localhost:3306"
echo "Redis: localhost:6379"
echo ""
echo "查看日志: $COMPOSE_CMD logs -f backend"
echo "停止服务: $COMPOSE_CMD down"
echo "=========================================="

