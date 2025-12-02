#!/bin/bash

# 远程服务器从Git仓库拉取代码并部署
# 使用方法: ./pull-code.sh [backend|frontend|all]

set -e

# 服务器配置
SERVER_IP="39.105.136.116"
SERVER_USER="root"
SSH_KEY="${HOME}/Desktop/chuchang/chuchang.pem"
DEPLOY_PATH="/opt/admin-system"

# Git配置（需要修改为实际的Git仓库地址）
GIT_REPO="${GIT_REPO:-}"  # 如果设置了环境变量，使用环境变量
GIT_BRANCH="${GIT_BRANCH:-main}"  # 默认分支

UPDATE_TYPE="${1:-all}"

echo "=========================================="
echo "从Git仓库拉取代码并部署"
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

# 在服务器上执行拉取和部署
ssh -i "${SSH_KEY}" -o StrictHostKeyChecking=no ${SERVER_USER}@${SERVER_IP} << ENDSSH
set -e

# 检查Git是否安装
if ! command -v git &> /dev/null; then
    echo "⚠️  Git未安装，正在安装..."
    if command -v yum &> /dev/null; then
        yum install -y git
    elif command -v apt-get &> /dev/null; then
        export DEBIAN_FRONTEND=noninteractive
        apt-get update
        apt-get install -y git
    else
        echo "❌ 无法安装Git，请手动安装"
        exit 1
    fi
    echo "✅ Git安装完成"
fi

cd ${DEPLOY_PATH}

# 检查Git仓库配置
if [ -z "${GIT_REPO}" ]; then
    # 如果未设置，尝试从现有目录获取
    if [ -d .git ]; then
        GIT_REPO=\$(git remote get-url origin 2>/dev/null || echo "")
        if [ -z "\${GIT_REPO}" ]; then
            echo "❌ 未找到Git仓库配置"
            echo "   请设置环境变量 GIT_REPO 或在此目录初始化Git仓库"
            exit 1
        fi
    else
        echo "❌ 未找到Git仓库"
        echo "   请设置环境变量: export GIT_REPO=your-git-repo-url"
        echo "   或在此目录初始化Git仓库"
        echo ""
        echo "   快速初始化："
        echo "   git init"
        echo "   git remote add origin <your-repo-url>"
        exit 1
    fi
fi

echo "Git仓库: \${GIT_REPO}"
echo "分支: ${GIT_BRANCH}"
echo ""

# 如果目录不存在或不是Git仓库，克隆
if [ ! -d .git ]; then
    echo "📥 克隆Git仓库..."
    if [ -d "${DEPLOY_PATH}" ] && [ "\$(ls -A ${DEPLOY_PATH} 2>/dev/null)" ]; then
        echo "⚠️  目录不为空，备份现有文件..."
        mv ${DEPLOY_PATH} ${DEPLOY_PATH}.backup.\$(date +%Y%m%d_%H%M%S)
    fi
    git clone \${GIT_REPO} ${DEPLOY_PATH}
    cd ${DEPLOY_PATH}
else
    echo "📥 拉取最新代码..."
    git fetch origin
    git checkout ${GIT_BRANCH}
    git pull origin ${GIT_BRANCH}
fi

echo "✅ 代码拉取完成"
echo ""

# 更新后端
if [ "${UPDATE_TYPE}" == "backend" ] || [ "${UPDATE_TYPE}" == "all" ]; then
    echo "=========================================="
    echo "更新后端"
    echo "=========================================="
    
    cd backend
    
    echo "下载Go依赖..."
    if command -v go &> /dev/null; then
        go mod download
        echo "✅ Go依赖下载完成"
    else
        echo "⚠️  Go未安装，跳过依赖下载"
    fi
    
    echo "停止旧服务..."
    cd ${DEPLOY_PATH}
    docker-compose stop backend 2>/dev/null || true
    
    echo "重新构建Docker镜像..."
    docker-compose build backend
    
    echo "启动服务..."
    docker-compose up -d backend
    
    echo "等待服务启动..."
    sleep 5
    
    echo "检查服务状态..."
    docker-compose ps backend
    
    echo "✅ 后端更新完成"
fi

# 更新前端
if [ "${UPDATE_TYPE}" == "frontend" ] || [ "${UPDATE_TYPE}" == "all" ]; then
    echo ""
    echo "=========================================="
    echo "更新前端"
    echo "=========================================="
    
    cd frontend
    
    echo "安装依赖..."
    if command -v npm &> /dev/null; then
        npm install
        echo "构建前端..."
        npm run build
        echo "✅ 前端构建完成"
    else
        echo "⚠️  npm未安装，跳过前端构建"
    fi
fi

echo ""
echo "=========================================="
echo "✅ 更新完成！"
echo "=========================================="
echo ""
echo "📝 服务状态："
docker-compose ps 2>/dev/null || echo "Docker Compose未运行"
echo ""
echo "📝 查看日志："
echo "   docker-compose logs -f backend"
ENDSSH

echo ""
echo "=========================================="
echo "✅ 远程拉取并部署完成！"
echo "=========================================="
