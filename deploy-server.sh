#!/bin/bash

# 服务器部署脚本
# 使用方法: ./deploy-server.sh

set -e

# 服务器配置
SERVER_IP="39.105.136.116"
SERVER_USER="root"
SERVER_PASSWORD="huojian123456!@#\$%^"
SERVER_PORT="22"
DEPLOY_PATH="/opt/admin-system"
SSH_KEY="${HOME}/Desktop/chuchang/chuchang.pem"  # SSH密钥文件路径（如果使用密钥认证）

echo "=========================================="
echo "服务器部署脚本"
echo "=========================================="
echo "服务器: ${SERVER_IP}"
echo "部署路径: ${DEPLOY_PATH}"
echo ""

# 检查是否安装了sshpass（用于自动输入密码）
if ! command -v sshpass &> /dev/null; then
    echo "⚠️  sshpass未安装，正在安装..."
    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        if command -v brew &> /dev/null; then
            brew install hudochenkov/sshpass/sshpass
        else
            echo "❌ 请先安装 Homebrew，然后运行: brew install hudochenkov/sshpass/sshpass"
            exit 1
        fi
    elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Linux
        sudo apt-get update && sudo apt-get install -y sshpass
    fi
fi

# 检查SSH连接
echo "🔍 检查服务器连接..."
echo "   尝试连接到: ${SERVER_USER}@${SERVER_IP}"

# 检查是否使用SSH密钥
USE_KEY=false
if [ -f "${SSH_KEY}" ]; then
    echo "   找到SSH密钥文件: ${SSH_KEY}"
    chmod 600 "${SSH_KEY}" 2>/dev/null
    USE_KEY=true
    SSH_CMD="ssh -i ${SSH_KEY}"
else
    echo "   未找到SSH密钥，尝试使用密码认证"
    SSH_CMD="sshpass -p '${SERVER_PASSWORD}' ssh"
fi

# 测试连接
if [ "$USE_KEY" = true ]; then
    if ${SSH_CMD} -o StrictHostKeyChecking=no -o ConnectTimeout=10 ${SERVER_USER}@${SERVER_IP} "echo '连接成功'" 2>/dev/null; then
        echo "✅ 服务器连接正常（使用SSH密钥）"
        CONNECT_SUCCESS=true
    else
        CONNECT_SUCCESS=false
    fi
else
    if sshpass -p "${SERVER_PASSWORD}" ssh -o StrictHostKeyChecking=no -o ConnectTimeout=10 ${SERVER_USER}@${SERVER_IP} "echo '连接成功'" 2>/dev/null; then
        echo "✅ 服务器连接正常（使用密码）"
        CONNECT_SUCCESS=true
    else
        CONNECT_SUCCESS=false
    fi
fi

if [ "$CONNECT_SUCCESS" = false ]; then
    echo ""
    echo "❌ 无法连接到服务器"
    echo ""
    echo "🔍 诊断信息："
    echo "   服务器IP: ${SERVER_IP}"
    echo "   用户名: ${SERVER_USER}"
    echo ""
    
    # 测试端口是否开放
    echo "   测试SSH端口(22)是否开放..."
    if command -v nc &> /dev/null; then
        if nc -z -w 5 ${SERVER_IP} 22 2>/dev/null; then
            echo "   ✅ 端口22已开放"
        else
            echo "   ❌ 端口22无法连接（可能被防火墙阻止）"
        fi
    fi
    
    # 尝试详细连接
    echo ""
    echo "   尝试详细连接（显示错误信息）..."
    sshpass -p "${SERVER_PASSWORD}" ssh -v -o StrictHostKeyChecking=no -o ConnectTimeout=10 ${SERVER_USER}@${SERVER_IP} "echo 'test'" 2>&1 | tail -5
    
    echo ""
    echo "💡 可能的解决方案："
    echo "   1. 检查服务器SSH服务是否运行:"
    echo "      (需要在服务器上执行) systemctl status sshd"
    echo ""
    echo "   2. 检查防火墙规则:"
    echo "      (需要在服务器上执行) ufw status 或 firewall-cmd --list-all"
    echo ""
    echo "   3. 尝试手动连接测试:"
    echo "      ssh ${SERVER_USER}@${SERVER_IP}"
    echo ""
    echo "   4. 如果使用密钥认证，请使用:"
    echo "      ssh -i your-key.pem ${SERVER_USER}@${SERVER_IP}"
    echo ""
    echo "   5. 如果服务器在内网，可能需要VPN或跳板机"
    echo "   6. 服务器可能只允许SSH密钥认证"
    echo "      如果使用密钥，请确保密钥文件在: ${SSH_KEY}"
    echo ""
    read -p "是否继续尝试部署？(y/n) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
    echo "⚠️  继续部署，但可能失败..."
fi

# 设置SSH命令变量供后续使用
if [ "$USE_KEY" = true ]; then
    SSH_CMD_PREFIX="ssh -i ${SSH_KEY} -o StrictHostKeyChecking=no"
    SCP_CMD_PREFIX="scp -i ${SSH_KEY} -o StrictHostKeyChecking=no"
else
    SSH_CMD_PREFIX="sshpass -p '${SERVER_PASSWORD}' ssh -o StrictHostKeyChecking=no"
    SCP_CMD_PREFIX="sshpass -p '${SERVER_PASSWORD}' scp -o StrictHostKeyChecking=no"
fi

# 在服务器上创建部署目录
echo ""
echo "📁 创建部署目录..."
${SSH_CMD_PREFIX} ${SERVER_USER}@${SERVER_IP} << 'ENDSSH'
mkdir -p /opt/admin-system
mkdir -p /opt/admin-system/backend
mkdir -p /opt/admin-system/frontend
mkdir -p /opt/admin-system/docs
ENDSSH

# 检查服务器环境
echo ""
echo "🔍 检查服务器环境..."
${SSH_CMD_PREFIX} ${SERVER_USER}@${SERVER_IP} << 'ENDSSH'
echo "检查Docker..."
if command -v docker &> /dev/null; then
    docker --version
else
    echo "⚠️  Docker未安装"
fi

echo "检查Docker Compose..."
if command -v docker-compose &> /dev/null; then
    docker-compose --version
else
    echo "⚠️  Docker Compose未安装"
fi

echo "检查Go..."
if command -v go &> /dev/null; then
    go version
fi
ENDSSH

# 打包项目文件（排除不需要的文件）
echo ""
echo "📦 打包项目文件..."
TEMP_DIR=$(mktemp -d)
TAR_FILE="${TEMP_DIR}/admin-system.tar.gz"

# 创建临时目录并复制文件
rsync -av --exclude='.git' \
          --exclude='node_modules' \
          --exclude='.env' \
          --exclude='backend/.env' \
          --exclude='*.log' \
          --exclude='.DS_Store' \
          --exclude='frontend/dist' \
          --exclude='backend/main' \
          ./ "${TEMP_DIR}/admin-system/"

# 打包
cd "${TEMP_DIR}"
tar -czf admin-system.tar.gz admin-system/
cd - > /dev/null

# 上传到服务器
echo ""
echo "📤 上传文件到服务器..."
${SCP_CMD_PREFIX} "${TAR_FILE}" ${SERVER_USER}@${SERVER_IP}:${DEPLOY_PATH}/

# 在服务器上解压和部署
echo ""
echo "🚀 在服务器上部署..."
${SSH_CMD_PREFIX} ${SERVER_USER}@${SERVER_IP} << ENDSSH
cd ${DEPLOY_PATH}
echo "解压文件..."
tar -xzf admin-system.tar.gz
mv admin-system/* .
rm -rf admin-system admin-system.tar.gz

echo "创建.env文件..."
if [ ! -f .env ]; then
    cp .env.example .env
    echo "⚠️  请编辑 .env 文件，修改生产环境的密码和密钥！"
fi

echo "设置文件权限..."
chmod +x deploy.sh start.sh restart.sh build.sh stop.sh

echo "✅ 文件部署完成"
ENDSSH

# 清理临时文件
rm -rf "${TEMP_DIR}"

echo ""
echo "=========================================="
echo "✅ 部署完成！"
echo "=========================================="
echo ""
echo "📝 下一步操作："
echo "   1. SSH登录服务器:"
echo "      ssh ${SERVER_USER}@${SERVER_IP}"
echo ""
echo "   2. 进入部署目录:"
echo "      cd ${DEPLOY_PATH}"
echo ""
echo "   3. 编辑配置文件:"
echo "      vi .env"
echo "      修改所有密码和密钥为生产环境值"
echo ""
echo "   4. 启动服务:"
echo "      ./deploy.sh"
echo ""
echo "   5. 查看日志:"
echo "      docker-compose logs -f backend"
echo ""
echo "=========================================="




