#!/bin/bash

# ============================================
# Docker 镜像加速器配置脚本
# ============================================
# 功能：配置 Docker 使用国内镜像源，解决拉取镜像超时问题
# ============================================

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# Docker daemon 配置文件路径
DAEMON_JSON="/etc/docker/daemon.json"
DAEMON_JSON_BACKUP="/etc/docker/daemon.json.backup.$(date +%Y%m%d_%H%M%S)"

# 国内镜像源列表（按优先级排序）
MIRRORS=(
    "https://docker.mirrors.ustc.edu.cn"
    "https://hub-mirror.c.163.com"
    "https://mirror.baidubce.com"
    "https://registry.docker-cn.com"
)

echo -e "${GREEN}=========================================="
echo "Docker 镜像加速器配置"
echo "==========================================${NC}"
echo ""

# 检查是否为 root 用户
if [ "$EUID" -ne 0 ]; then 
    echo -e "${RED}❌ 错误: 请使用 root 权限运行此脚本${NC}"
    echo "   使用命令: sudo $0"
    exit 1
fi

# 检查 Docker 是否安装
if ! command -v docker &> /dev/null; then
    echo -e "${RED}❌ 错误: Docker 未安装${NC}"
    exit 1
fi

# 备份现有配置
if [ -f "$DAEMON_JSON" ]; then
    echo -e "${BLUE}📋 备份现有配置...${NC}"
    cp "$DAEMON_JSON" "$DAEMON_JSON_BACKUP"
    echo -e "${GREEN}✅ 配置已备份到: $DAEMON_JSON_BACKUP${NC}"
fi

# 创建配置目录
mkdir -p /etc/docker

# 生成新的配置
echo -e "${BLUE}🔧 配置镜像加速器...${NC}"

# 读取现有配置（如果存在）
if [ -f "$DAEMON_JSON" ]; then
    # 使用 jq 更新配置（如果安装了 jq）
    if command -v jq &> /dev/null; then
        # 合并现有配置
        jq '.registry-mirrors = $mirrors' --argjson mirrors "$(printf '%s\n' "${MIRRORS[@]}" | jq -R . | jq -s .)" "$DAEMON_JSON" > "${DAEMON_JSON}.tmp" && mv "${DAEMON_JSON}.tmp" "$DAEMON_JSON"
    else
        # 如果没有 jq，创建新配置
        cat > "$DAEMON_JSON" <<EOF
{
  "registry-mirrors": [
    "https://docker.mirrors.ustc.edu.cn",
    "https://hub-mirror.c.163.com",
    "https://mirror.baidubce.com",
    "https://registry.docker-cn.com"
  ],
  "insecure-registries": [],
  "experimental": false
}
EOF
    fi
else
    # 创建新配置
    cat > "$DAEMON_JSON" <<EOF
{
  "registry-mirrors": [
    "https://docker.mirrors.ustc.edu.cn",
    "https://hub-mirror.c.163.com",
    "https://mirror.baidubce.com",
    "https://registry.docker-cn.com"
  ],
  "insecure-registries": [],
  "experimental": false
}
EOF
fi

# 验证 JSON 格式
if ! python3 -m json.tool "$DAEMON_JSON" > /dev/null 2>&1; then
    echo -e "${RED}❌ 错误: 生成的配置文件格式不正确${NC}"
    if [ -f "$DAEMON_JSON_BACKUP" ]; then
        echo -e "${YELLOW}⚠️  恢复备份配置...${NC}"
        mv "$DAEMON_JSON_BACKUP" "$DAEMON_JSON"
    fi
    exit 1
fi

echo -e "${GREEN}✅ 镜像加速器配置完成${NC}"
echo ""
echo -e "${BLUE}📝 配置的镜像源:${NC}"
for mirror in "${MIRRORS[@]}"; do
    echo "   - $mirror"
done

# 重启 Docker 服务
echo ""
echo -e "${BLUE}🔄 重启 Docker 服务...${NC}"
if systemctl restart docker 2>/dev/null; then
    echo -e "${GREEN}✅ Docker 服务已重启${NC}"
elif service docker restart 2>/dev/null; then
    echo -e "${GREEN}✅ Docker 服务已重启${NC}"
else
    echo -e "${YELLOW}⚠️  无法自动重启 Docker，请手动执行:${NC}"
    echo "   systemctl restart docker"
    echo "   或"
    echo "   service docker restart"
fi

# 验证配置
echo ""
echo -e "${BLUE}🔍 验证配置...${NC}"
sleep 2
if docker info | grep -q "Registry Mirrors"; then
    echo -e "${GREEN}✅ 配置已生效${NC}"
    echo ""
    echo -e "${BLUE}📋 当前镜像源配置:${NC}"
    docker info | grep -A 10 "Registry Mirrors" || true
else
    echo -e "${YELLOW}⚠️  配置可能未生效，请检查 Docker 日志${NC}"
fi

echo ""
echo -e "${GREEN}=========================================="
echo "✅ 配置完成！"
echo "==========================================${NC}"
echo ""
echo -e "${BLUE}💡 提示:${NC}"
echo "   现在可以重新运行 ./deploy.sh 进行部署"
echo "   如果仍有问题，可以尝试:"
echo "   1. 检查网络连接"
echo "   2. 尝试其他镜像源"
echo "   3. 使用代理"
