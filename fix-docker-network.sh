#!/bin/bash

# ============================================
# Docker 网络问题修复脚本
# ============================================
# 功能：修复 Docker 无法拉取镜像的网络问题
# ============================================

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${GREEN}=========================================="
echo "Docker 网络问题修复脚本"
echo "==========================================${NC}"
echo ""

# 检查是否为 root 用户
if [ "$EUID" -ne 0 ]; then 
    echo -e "${RED}❌ 错误: 请使用 root 权限运行此脚本${NC}"
    echo "   使用命令: sudo $0"
    exit 1
fi

# 1. 检查 DNS 配置
echo -e "${BLUE}🔍 步骤 1/4: 检查 DNS 配置...${NC}"
if [ -f /etc/resolv.conf ]; then
    echo "当前 DNS 配置:"
    cat /etc/resolv.conf | grep nameserver || echo "未找到 nameserver 配置"
    echo ""
    
    # 检查是否能解析域名
    if ! nslookup docker.io > /dev/null 2>&1; then
        echo -e "${YELLOW}⚠️  无法解析 docker.io，DNS 可能有问题${NC}"
        echo "建议添加以下 DNS 服务器："
        echo "  8.8.8.8"
        echo "  114.114.114.114"
        echo "  223.5.5.5"
    fi
else
    echo -e "${YELLOW}⚠️  /etc/resolv.conf 不存在${NC}"
fi

# 2. 配置 Docker 镜像加速器（阿里云专用）
echo ""
echo -e "${BLUE}🔧 步骤 2/4: 配置 Docker 镜像加速器...${NC}"

DAEMON_JSON="/etc/docker/daemon.json"
DAEMON_JSON_BACKUP="/etc/docker/daemon.json.backup.$(date +%Y%m%d_%H%M%S)"

# 备份现有配置
if [ -f "$DAEMON_JSON" ]; then
    cp "$DAEMON_JSON" "$DAEMON_JSON_BACKUP"
    echo -e "${GREEN}✅ 配置已备份到: $DAEMON_JSON_BACKUP${NC}"
fi

# 创建配置目录
mkdir -p /etc/docker

# 创建新的配置（使用阿里云镜像加速器）
cat > "$DAEMON_JSON" <<EOF
{
  "registry-mirrors": [
    "https://registry.cn-hangzhou.aliyuncs.com",
    "https://dockerhub.azk8s.cn",
    "https://reg-mirror.qiniu.com"
  ],
  "insecure-registries": [],
  "experimental": false,
  "dns": ["8.8.8.8", "114.114.114.114", "223.5.5.5"]
}
EOF

echo -e "${GREEN}✅ Docker 镜像加速器配置完成${NC}"

# 3. 测试网络连接
echo ""
echo -e "${BLUE}🌐 步骤 3/4: 测试网络连接...${NC}"

# 测试 DNS
echo "测试 DNS 解析..."
if nslookup registry.cn-hangzhou.aliyuncs.com > /dev/null 2>&1; then
    echo -e "${GREEN}✅ 可以解析 registry.cn-hangzhou.aliyuncs.com${NC}"
else
    echo -e "${YELLOW}⚠️  无法解析 registry.cn-hangzhou.aliyuncs.com${NC}"
fi

# 测试连接
echo "测试网络连接..."
if timeout 5 curl -s https://registry.cn-hangzhou.aliyuncs.com > /dev/null 2>&1; then
    echo -e "${GREEN}✅ 可以连接到阿里云容器镜像服务${NC}"
else
    echo -e "${YELLOW}⚠️  无法连接到阿里云容器镜像服务${NC}"
    echo "   可能原因："
    echo "   1. 服务器网络限制"
    echo "   2. 防火墙阻止"
    echo "   3. 需要使用代理"
fi

# 4. 重启 Docker 服务
echo ""
echo -e "${BLUE}🔄 步骤 4/4: 重启 Docker 服务...${NC}"
if systemctl restart docker 2>/dev/null; then
    echo -e "${GREEN}✅ Docker 服务已重启${NC}"
elif service docker restart 2>/dev/null; then
    echo -e "${GREEN}✅ Docker 服务已重启${NC}"
else
    echo -e "${YELLOW}⚠️  无法自动重启 Docker，请手动执行:${NC}"
    echo "   systemctl restart docker"
fi

# 等待 Docker 启动
sleep 3

# 验证配置
echo ""
echo -e "${BLUE}🔍 验证配置...${NC}"
if docker info | grep -q "Registry Mirrors"; then
    echo -e "${GREEN}✅ 镜像加速器配置已生效${NC}"
    echo ""
    echo -e "${BLUE}📋 当前镜像源配置:${NC}"
    docker info | grep -A 10 "Registry Mirrors" || true
else
    echo -e "${YELLOW}⚠️  配置可能未生效，请检查 Docker 日志${NC}"
fi

# 测试拉取镜像
echo ""
echo -e "${BLUE}🧪 测试拉取镜像...${NC}"
if timeout 30 docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/alpine:latest > /dev/null 2>&1; then
    echo -e "${GREEN}✅ 可以拉取阿里云镜像${NC}"
else
    echo -e "${YELLOW}⚠️  无法拉取镜像，可能需要：${NC}"
    echo "   1. 检查网络连接"
    echo "   2. 使用代理"
    echo "   3. 联系网络管理员"
fi

echo ""
echo -e "${GREEN}=========================================="
echo "✅ 配置完成！"
echo "==========================================${NC}"
echo ""
echo -e "${BLUE}💡 提示:${NC}"
echo "   如果仍然无法拉取镜像，请："
echo "   1. 检查服务器是否在阿里云，如果是，确保使用阿里云内网镜像加速器"
echo "   2. 检查防火墙和安全组设置"
echo "   3. 考虑使用代理或 VPN"
echo "   4. 直接修改 Dockerfile 使用阿里云镜像地址"
