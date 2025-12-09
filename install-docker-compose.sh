#!/bin/bash

# Docker Compose 安装脚本
# 适用于 CentOS/RHEL 系统

echo "=========================================="
echo "安装 Docker Compose"
echo "=========================================="

# 检查是否已安装
if command -v docker-compose &> /dev/null; then
    echo "✅ docker-compose 已安装: $(docker-compose --version)"
    exit 0
fi

if docker compose version &> /dev/null 2>&1; then
    echo "✅ docker compose 已安装: $(docker compose version)"
    exit 0
fi

# 检查Docker是否安装
if ! command -v docker &> /dev/null; then
    echo "❌ 错误: Docker未安装，请先安装Docker"
    exit 1
fi

echo ""
echo "检测到 Docker 版本: $(docker --version)"

# 检查Docker版本，新版本（20.10+）内置了 docker compose
DOCKER_VERSION=$(docker --version | grep -oE '[0-9]+\.[0-9]+' | head -1)
MAJOR_VERSION=$(echo $DOCKER_VERSION | cut -d. -f1)
MINOR_VERSION=$(echo $DOCKER_VERSION | cut -d. -f2)

if [ "$MAJOR_VERSION" -gt 20 ] || ([ "$MAJOR_VERSION" -eq 20 ] && [ "$MINOR_VERSION" -ge 10 ]); then
    echo "✅ 检测到 Docker 20.10+，已内置 docker compose 插件"
    echo "   使用方式: docker compose (注意是空格，不是连字符)"
    echo ""
    echo "测试命令:"
    docker compose version
    exit 0
fi

# 安装旧版本的 docker-compose
echo ""
echo "正在安装 docker-compose..."

# 下载最新版本的 docker-compose
COMPOSE_VERSION=$(curl -s https://api.github.com/repos/docker/compose/releases/latest | grep 'tag_name' | cut -d\" -f4)
if [ -z "$COMPOSE_VERSION" ]; then
    COMPOSE_VERSION="v2.24.0"  # 备用版本
fi

echo "下载 docker-compose $COMPOSE_VERSION..."

# 下载到 /usr/local/bin
sudo curl -L "https://github.com/docker/compose/releases/download/${COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

# 添加执行权限
sudo chmod +x /usr/local/bin/docker-compose

# 创建软链接（可选）
if [ ! -f /usr/bin/docker-compose ]; then
    sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
fi

# 验证安装
if command -v docker-compose &> /dev/null; then
    echo ""
    echo "✅ docker-compose 安装成功！"
    echo "   版本: $(docker-compose --version)"
    echo ""
    echo "使用方法:"
    echo "  docker-compose up -d"
    echo "  docker-compose ps"
    echo "  docker-compose down"
else
    echo "❌ 安装失败，请手动安装"
    exit 1
fi
