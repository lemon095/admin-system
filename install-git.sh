#!/bin/bash

# 在远程服务器上安装Git
# 使用方法: ./install-git.sh

set -e

# 服务器配置
SERVER_IP="39.105.136.116"
SERVER_USER="root"
SSH_KEY="${HOME}/Desktop/chuchang/chuchang.pem"

echo "=========================================="
echo "在远程服务器上安装Git"
echo "=========================================="
echo "服务器: ${SERVER_USER}@${SERVER_IP}"
echo ""

# 检查SSH密钥
if [ ! -f "${SSH_KEY}" ]; then
    echo "❌ SSH密钥文件不存在: ${SSH_KEY}"
    exit 1
fi

chmod 600 "${SSH_KEY}"

# 在服务器上安装Git
ssh -i "${SSH_KEY}" -o StrictHostKeyChecking=no ${SERVER_USER}@${SERVER_IP} << 'ENDSSH'
set -e

echo "🔍 检查Git安装..."
if command -v git &> /dev/null; then
    echo "✅ Git已安装: $(git --version)"
    exit 0
fi

echo "📦 检测操作系统..."
if [ -f /etc/os-release ]; then
    . /etc/os-release
    OS=$ID
    VER=$VERSION_ID
else
    echo "❌ 无法检测操作系统"
    exit 1
fi

echo "操作系统: $OS $VER"
echo ""

echo "安装Git..."

if [ "$OS" == "centos" ] || [ "$OS" == "rhel" ] || [ "$OS" == "almalinux" ] || [ "$OS" == "rocky" ]; then
    # CentOS/RHEL/AlmaLinux/Rocky
    echo "使用yum安装..."
    yum install -y git
elif [ "$OS" == "ubuntu" ] || [ "$OS" == "debian" ]; then
    # Ubuntu/Debian
    echo "使用apt-get安装..."
    export DEBIAN_FRONTEND=noninteractive
    apt-get update
    apt-get install -y git
else
    echo "❌ 不支持的操作系统: $OS"
    exit 1
fi

# 验证安装
if command -v git &> /dev/null; then
    echo ""
    echo "✅ Git安装成功: $(git --version)"
    echo ""
    echo "📝 Git配置建议："
    echo "   git config --global user.name 'Your Name'"
    echo "   git config --global user.email 'your.email@example.com'"
else
    echo "❌ Git安装失败"
    exit 1
fi
ENDSSH

echo ""
echo "=========================================="
echo "✅ Git安装完成！"
echo "=========================================="
echo ""
echo "📝 下一步："
echo "   1. 配置Git（可选）："
echo "      ssh -i ${SSH_KEY} ${SERVER_USER}@${SERVER_IP}"
echo "      git config --global user.name 'Your Name'"
echo "      git config --global user.email 'your.email@example.com'"
echo ""
echo "   2. 初始化Git仓库："
echo "      cd /opt/admin-system"
echo "      git init"
echo "      git remote add origin <your-repo-url>"
echo ""
echo "   3. 或直接克隆仓库："
echo "      git clone <your-repo-url> /opt/admin-system"
echo "=========================================="

