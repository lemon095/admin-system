#!/bin/bash

# 在远程服务器上设置Git拉取功能
# 使用方法: ./setup-git-pull.sh

set -e

# 服务器配置
SERVER_IP="39.105.136.116"
SERVER_USER="root"
SSH_KEY="${HOME}/Desktop/chuchang/chuchang.pem"
DEPLOY_PATH="/opt/admin-system"

echo "=========================================="
echo "设置远程服务器Git拉取功能"
echo "=========================================="
echo "服务器: ${SERVER_USER}@${SERVER_IP}"
echo ""

# 检查SSH密钥
if [ ! -f "${SSH_KEY}" ]; then
    echo "❌ SSH密钥文件不存在: ${SSH_KEY}"
    exit 1
fi

chmod 600 "${SSH_KEY}"

# 在服务器上设置
ssh -i "${SSH_KEY}" -o StrictHostKeyChecking=no ${SERVER_USER}@${SERVER_IP} << 'ENDSSH'
set -e

echo "🔍 检查Git安装..."
if ! command -v git &> /dev/null; then
    echo "安装Git..."
    if command -v yum &> /dev/null; then
        yum install -y git
    elif command -v apt-get &> /dev/null; then
        apt-get update
        apt-get install -y git
    else
        echo "❌ 无法安装Git，请手动安装"
        exit 1
    fi
fi

echo "✅ Git已安装: $(git --version)"
echo ""

echo "📝 创建Git拉取脚本..."
cat > /opt/admin-system/pull-and-deploy.sh << 'SCRIPT'
#!/bin/bash
# 在服务器上执行的Git拉取和部署脚本

set -e

DEPLOY_PATH="/opt/admin-system"
GIT_BRANCH="${1:-main}"

cd ${DEPLOY_PATH}

# 检查是否是Git仓库
if [ ! -d .git ]; then
    echo "❌ 当前目录不是Git仓库"
    echo "   请先运行: git clone <your-repo-url> ${DEPLOY_PATH}"
    exit 1
fi

echo "=========================================="
echo "从Git拉取代码"
echo "=========================================="
echo "当前分支: $(git branch --show-current)"
echo "拉取分支: ${GIT_BRANCH}"
echo ""

# 备份当前代码
echo "💾 备份当前代码..."
BACKUP_DIR="/tmp/admin-system-backup-$(date +%Y%m%d_%H%M%S)"
mkdir -p ${BACKUP_DIR}
tar -czf ${BACKUP_DIR}/code-backup.tar.gz . 2>/dev/null || true
echo "备份位置: ${BACKUP_DIR}"
echo ""

# 拉取代码
echo "📥 拉取最新代码..."
git fetch origin
git checkout ${GIT_BRANCH}
git pull origin ${GIT_BRANCH}

echo "✅ 代码拉取完成"
echo ""

# 更新后端
echo "=========================================="
echo "更新后端"
echo "=========================================="

cd backend

# 下载依赖
if command -v go &> /dev/null; then
    echo "下载Go依赖..."
    go mod download
    echo "✅ Go依赖下载完成"
fi

# 重新构建并启动
cd ${DEPLOY_PATH}
echo "停止旧服务..."
docker-compose stop backend 2>/dev/null || true

echo "重新构建Docker镜像..."
docker-compose build backend

echo "启动服务..."
docker-compose up -d backend

echo "等待服务启动..."
sleep 5

echo "检查服务状态..."
docker-compose ps backend

echo ""
echo "=========================================="
echo "✅ 更新完成！"
echo "=========================================="
echo ""
echo "📝 查看日志: docker-compose logs -f backend"
SCRIPT

chmod +x /opt/admin-system/pull-and-deploy.sh

echo "✅ Git拉取脚本已创建: /opt/admin-system/pull-and-deploy.sh"
echo ""

echo "📝 使用方法："
echo "   1. 如果还未初始化Git仓库："
echo "      cd /opt/admin-system"
echo "      git clone <your-repo-url> ."
echo ""
echo "   2. 拉取并部署："
echo "      /opt/admin-system/pull-and-deploy.sh [branch-name]"
echo ""
echo "   3. 或使用本地脚本："
echo "      ./pull-code.sh"
ENDSSH

echo ""
echo "=========================================="
echo "✅ Git拉取功能设置完成！"
echo "=========================================="
echo ""
echo "📝 下一步："
echo "   1. 在服务器上初始化Git仓库（如果还没有）："
echo "      ssh -i ${SSH_KEY} ${SERVER_USER}@${SERVER_IP}"
echo "      cd /opt/admin-system"
echo "      git clone <your-repo-url> ."
echo ""
echo "   2. 使用Git拉取代码："
echo "      ./pull-code.sh"
echo ""
echo "   3. 或在服务器上直接执行："
echo "      /opt/admin-system/pull-and-deploy.sh"
echo "=========================================="
