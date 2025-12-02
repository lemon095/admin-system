#!/bin/bash

# 远程在服务器上安装MySQL和Redis的脚本
# 使用方法: ./install-db-remote.sh

set -e

# 服务器配置
SERVER_IP="39.105.136.116"
SERVER_USER="root"
SSH_KEY="${HOME}/Desktop/chuchang/chuchang.pem"
DEPLOY_PATH="/opt/admin-system"

echo "=========================================="
echo "远程安装MySQL和Redis"
echo "=========================================="
echo "服务器: ${SERVER_USER}@${SERVER_IP}"
echo ""

# 检查SSH密钥
if [ ! -f "${SSH_KEY}" ]; then
    echo "❌ SSH密钥文件不存在: ${SSH_KEY}"
    exit 1
fi

chmod 600 "${SSH_KEY}"

# 上传安装脚本
echo "📤 上传安装脚本到服务器..."
scp -i "${SSH_KEY}" -o StrictHostKeyChecking=no install-mysql-redis.sh ${SERVER_USER}@${SERVER_IP}:/tmp/

# 在服务器上执行安装
echo ""
echo "🚀 在服务器上执行安装..."
ssh -i "${SSH_KEY}" -o StrictHostKeyChecking=no ${SERVER_USER}@${SERVER_IP} << 'ENDSSH'
chmod +x /tmp/install-mysql-redis.sh
/tmp/install-mysql-redis.sh
ENDSSH

echo ""
echo "=========================================="
echo "✅ 安装完成！"
echo "=========================================="
echo ""
echo "📝 下一步："
echo "   1. 在服务器上配置.env文件，设置MySQL密码"
echo "   2. 运行部署脚本: ./deploy-server.sh"
echo "=========================================="

