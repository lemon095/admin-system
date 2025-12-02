#!/bin/bash

echo "=========================================="
echo "启动管理系统"
echo "=========================================="

# 检查Docker是否运行
if ! docker info > /dev/null 2>&1; then
    echo "错误: Docker未运行，请先启动Docker"
    exit 1
fi

# 启动服务
echo "正在启动服务..."
docker-compose up -d

# 等待服务启动
echo "等待服务启动..."
sleep 5

# 检查服务状态
echo ""
echo "服务状态:"
docker-compose ps

echo ""
echo "=========================================="
echo "服务启动完成！"
echo "=========================================="
echo "后端API: http://localhost:7701"
echo "MySQL: localhost:3306"
echo "Redis: localhost:6379"
echo ""
echo "查看日志: docker-compose logs -f backend"
echo "停止服务: docker-compose down"
echo "=========================================="

