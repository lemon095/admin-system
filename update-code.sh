#!/bin/bash

# 更新远程服务器代码脚本
# 使用方法: ./update-code.sh [backend|frontend|all]

set -e

# 服务器配置
SERVER_IP="39.105.136.116"
SERVER_USER="root"
SSH_KEY="${HOME}/Desktop/chuchang/chuchang.pem"
DEPLOY_PATH="/opt/admin-system"

# 默认更新全部
UPDATE_TYPE="${1:-all}"

echo "=========================================="
echo "更新远程服务器代码"
echo "=========================================="
echo "服务器: ${SERVER_USER}@${SERVER_IP}"
echo "更新类型: ${UPDATE_TYPE}"
echo "部署路径: ${DEPLOY_PATH}"
echo ""

# 检查SSH密钥
if [ ! -f "${SSH_KEY}" ]; then
    echo "❌ SSH密钥文件不存在: ${SSH_KEY}"
    exit 1
fi

chmod 600 "${SSH_KEY}"

# 检查服务器连接
echo "🔍 检查服务器连接..."
if ! ssh -i "${SSH_KEY}" -o StrictHostKeyChecking=no -o ConnectTimeout=10 ${SERVER_USER}@${SERVER_IP} "echo '连接成功'" > /dev/null 2>&1; then
    echo "❌ 无法连接到服务器"
    exit 1
fi
echo "✅ 服务器连接正常"
echo ""

# 创建临时目录
TEMP_DIR=$(mktemp -d)
echo "📦 准备更新文件..."

# 更新后端
if [ "$UPDATE_TYPE" == "backend" ] || [ "$UPDATE_TYPE" == "all" ]; then
    echo ""
    echo "=========================================="
    echo "更新后端代码"
    echo "=========================================="
    
    # 打包后端文件
    echo "打包后端文件..."
    cd backend
    tar -czf "${TEMP_DIR}/backend-update.tar.gz" \
        --exclude='.env' \
        --exclude='main' \
        --exclude='*.log' \
        --exclude='.git' \
        --exclude='go.sum' \
        .
    cd ..
    
    # 上传到服务器
    echo "上传后端文件..."
    scp -i "${SSH_KEY}" -o StrictHostKeyChecking=no "${TEMP_DIR}/backend-update.tar.gz" ${SERVER_USER}@${SERVER_IP}:/tmp/
    
    # 在服务器上更新
    echo "在服务器上更新后端..."
    ssh -i "${SSH_KEY}" -o StrictHostKeyChecking=no ${SERVER_USER}@${SERVER_IP} << ENDSSH
cd ${DEPLOY_PATH}/backend
echo "备份当前代码..."
tar -czf /tmp/backend-backup-\$(date +%Y%m%d_%H%M%S).tar.gz . 2>/dev/null || true

echo "解压新代码..."
tar -xzf /tmp/backend-update.tar.gz

echo "下载Go依赖..."
if command -v go &> /dev/null; then
    go mod download
else
    echo "⚠️  Go未安装，跳过依赖下载"
fi

echo "清理临时文件..."
rm -f /tmp/backend-update.tar.gz

echo "✅ 后端代码更新完成"
ENDSSH
    
    echo "✅ 后端更新完成"
fi

# 更新前端
if [ "$UPDATE_TYPE" == "frontend" ] || [ "$UPDATE_TYPE" == "all" ]; then
    echo ""
    echo "=========================================="
    echo "更新前端代码"
    echo "=========================================="
    
    # 打包前端文件
    echo "打包前端文件..."
    cd frontend
    tar -czf "${TEMP_DIR}/frontend-update.tar.gz" \
        --exclude='node_modules' \
        --exclude='dist' \
        --exclude='.env' \
        --exclude='*.log' \
        --exclude='.git' \
        .
    cd ..
    
    # 上传到服务器
    echo "上传前端文件..."
    scp -i "${SSH_KEY}" -o StrictHostKeyChecking=no "${TEMP_DIR}/frontend-update.tar.gz" ${SERVER_USER}@${SERVER_IP}:/tmp/
    
    # 在服务器上更新
    echo "在服务器上更新前端..."
    ssh -i "${SSH_KEY}" -o StrictHostKeyChecking=no ${SERVER_USER}@${SERVER_IP} << ENDSSH
cd ${DEPLOY_PATH}/frontend
echo "备份当前代码..."
tar -czf /tmp/frontend-backup-\$(date +%Y%m%d_%H%M%S).tar.gz . 2>/dev/null || true

echo "解压新代码..."
tar -xzf /tmp/frontend-update.tar.gz

echo "安装依赖（如果需要）..."
if command -v npm &> /dev/null; then
    npm install
else
    echo "⚠️  npm未安装，跳过依赖安装"
fi

echo "清理临时文件..."
rm -f /tmp/frontend-update.tar.gz

echo "✅ 前端代码更新完成"
ENDSSH
    
    echo "✅ 前端更新完成"
fi

# 清理本地临时文件
rm -rf "${TEMP_DIR}"

echo ""
echo "=========================================="
echo "✅ 代码更新完成！"
echo "=========================================="
echo ""
echo "📝 下一步操作："
if [ "$UPDATE_TYPE" == "backend" ] || [ "$UPDATE_TYPE" == "all" ]; then
    echo "   后端: 需要重新构建Docker镜像并重启服务"
    echo "   ssh -i ${SSH_KEY} ${SERVER_USER}@${SERVER_IP}"
    echo "   cd ${DEPLOY_PATH}"
    echo "   docker-compose build backend"
    echo "   docker-compose up -d backend"
fi
if [ "$UPDATE_TYPE" == "frontend" ] || [ "$UPDATE_TYPE" == "all" ]; then
    echo "   前端: 需要重新构建（如果使用Docker）"
    echo "   或直接运行: npm run build"
fi
echo "=========================================="

