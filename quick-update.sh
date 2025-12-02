#!/bin/bash

# 快速更新脚本 - 只更新代码，不重新构建（最快）
# 使用方法: ./quick-update.sh [backend|frontend|all]

set -e

# 服务器配置
SERVER_IP="39.105.136.116"
SERVER_USER="root"
SSH_KEY="${HOME}/Desktop/chuchang/chuchang.pem"
DEPLOY_PATH="/opt/admin-system"

UPDATE_TYPE="${1:-all}"

echo "=========================================="
echo "快速更新代码（不重新构建）"
echo "=========================================="
echo "服务器: ${SERVER_USER}@${SERVER_IP}"
echo "更新类型: ${UPDATE_TYPE}"
echo ""

# 检查SSH密钥
if [ ! -f "${SSH_KEY}" ]; then
    echo "❌ SSH密钥文件不存在: ${SSH_KEY}"
    exit 1
fi

chmod 600 "${SSH_KEY}"

# 检查rsync是否可用，如果不可用则使用tar+scp
if command -v rsync &> /dev/null; then
    echo "🔄 使用rsync同步文件..."
    USE_RSYNC=true
else
    echo "⚠️  rsync未安装，使用tar+scp方式..."
    USE_RSYNC=false
fi

# 创建临时目录
TEMP_DIR=$(mktemp -d)

if [ "$UPDATE_TYPE" == "backend" ] || [ "$UPDATE_TYPE" == "all" ]; then
    echo ""
    echo "同步后端代码..."
    
    if [ "$USE_RSYNC" = true ]; then
        rsync -avz --delete \
            -e "ssh -i ${SSH_KEY} -o StrictHostKeyChecking=no" \
            --exclude='.env' \
            --exclude='main' \
            --exclude='*.log' \
            --exclude='.git' \
            --exclude='go.sum' \
            backend/ ${SERVER_USER}@${SERVER_IP}:${DEPLOY_PATH}/backend/
    else
        # 使用tar+scp
        cd backend
        tar -czf "${TEMP_DIR}/backend-sync.tar.gz" \
            --exclude='.env' \
            --exclude='main' \
            --exclude='*.log' \
            --exclude='.git' \
            --exclude='go.sum' \
            .
        cd ..
        scp -i "${SSH_KEY}" -o StrictHostKeyChecking=no "${TEMP_DIR}/backend-sync.tar.gz" ${SERVER_USER}@${SERVER_IP}:/tmp/
        ssh -i "${SSH_KEY}" -o StrictHostKeyChecking=no ${SERVER_USER}@${SERVER_IP} "cd ${DEPLOY_PATH}/backend && tar -xzf /tmp/backend-sync.tar.gz && rm -f /tmp/backend-sync.tar.gz"
        rm -f "${TEMP_DIR}/backend-sync.tar.gz"
    fi
    echo "✅ 后端代码同步完成"
fi

if [ "$UPDATE_TYPE" == "frontend" ] || [ "$UPDATE_TYPE" == "all" ]; then
    echo ""
    echo "同步前端代码..."
    
    if [ "$USE_RSYNC" = true ]; then
        rsync -avz --delete \
            -e "ssh -i ${SSH_KEY} -o StrictHostKeyChecking=no" \
            --exclude='node_modules' \
            --exclude='dist' \
            --exclude='.env' \
            --exclude='*.log' \
            --exclude='.git' \
            frontend/ ${SERVER_USER}@${SERVER_IP}:${DEPLOY_PATH}/frontend/
    else
        # 使用tar+scp
        cd frontend
        tar -czf "${TEMP_DIR}/frontend-sync.tar.gz" \
            --exclude='node_modules' \
            --exclude='dist' \
            --exclude='.env' \
            --exclude='*.log' \
            --exclude='.git' \
            .
        cd ..
        scp -i "${SSH_KEY}" -o StrictHostKeyChecking=no "${TEMP_DIR}/frontend-sync.tar.gz" ${SERVER_USER}@${SERVER_IP}:/tmp/
        ssh -i "${SSH_KEY}" -o StrictHostKeyChecking=no ${SERVER_USER}@${SERVER_IP} "cd ${DEPLOY_PATH}/frontend && tar -xzf /tmp/frontend-sync.tar.gz && rm -f /tmp/frontend-sync.tar.gz"
        rm -f "${TEMP_DIR}/frontend-sync.tar.gz"
    fi
    echo "✅ 前端代码同步完成"
fi

# 清理临时目录
rm -rf "${TEMP_DIR}"

echo ""
echo "=========================================="
echo "✅ 代码同步完成！"
echo "=========================================="
echo ""
echo "📝 注意："
echo "   此脚本只更新代码文件，不重新构建"
echo "   如果需要重新构建，请运行:"
echo "   ./update-and-deploy.sh ${UPDATE_TYPE}"
echo "=========================================="

