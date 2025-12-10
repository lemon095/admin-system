#!/bin/bash

# 在服务器上初始化 Git 仓库并配置
# 使用方法: 在服务器上执行 ./init-git-repo.sh

set -e

echo "=========================================="
echo "初始化 Git 仓库配置"
echo "=========================================="

# Git 仓库配置（仅针对当前仓库）
GIT_REPO_URL="git@github.com:lemon095/admin-system.git"
GIT_BRANCH="aki"
GIT_USER_NAME="Server Deploy"
GIT_USER_EMAIL="deploy@admin-system.local"

# 检查当前目录
CURRENT_DIR=$(pwd)
echo "当前目录: ${CURRENT_DIR}"
echo ""

# 1. 检查 Git 是否安装
if ! command -v git &> /dev/null; then
    echo "❌ Git 未安装，正在安装..."
    if command -v yum &> /dev/null; then
        yum install -y git
    elif command -v apt-get &> /dev/null; then
        apt-get update
        apt-get install -y git
    else
        echo "❌ 无法自动安装 Git，请手动安装"
        exit 1
    fi
fi

echo "✅ Git 已安装: $(git --version)"
echo ""

# 2. 初始化 Git 仓库（如果还没有）
if [ ! -d .git ]; then
    echo "📦 初始化 Git 仓库..."
    git init
    echo "✅ Git 仓库初始化完成"
else
    echo "✅ Git 仓库已存在"
fi
echo ""

# 3. 配置远程仓库
echo "🔗 配置远程仓库..."
if git remote get-url origin &> /dev/null; then
    echo "   当前远程地址: $(git remote get-url origin)"
    read -p "   是否更新远程地址? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        git remote set-url origin "${GIT_REPO_URL}"
        echo "✅ 远程地址已更新"
    else
        echo "   保持现有远程地址"
    fi
else
    git remote add origin "${GIT_REPO_URL}"
    echo "✅ 远程仓库已添加"
fi
echo ""

# 4. 配置 Git 用户信息（仅针对当前仓库）
echo "👤 配置 Git 用户信息（仅针对当前仓库）..."
git config user.name "${GIT_USER_NAME}"
git config user.email "${GIT_USER_EMAIL}"
echo "✅ 用户信息已配置:"
echo "   用户名: $(git config user.name)"
echo "   邮箱: $(git config user.email)"
echo ""

# 5. 配置其他 Git 设置（仅针对当前仓库）
echo "⚙️  配置其他 Git 设置..."
git config pull.rebase false  # 使用 merge 策略
git config core.autocrlf false  # 不自动转换换行符
echo "✅ Git 配置完成"
echo ""

# 6. 显示当前配置
echo "=========================================="
echo "当前 Git 配置"
echo "=========================================="
echo "远程仓库:"
git remote -v
echo ""
echo "当前分支: $(git branch --show-current 2>/dev/null || echo '无分支')"
echo ""
echo "仓库配置:"
git config --local --list | grep -E "(user\.|remote\.|branch\.)" || echo "无额外配置"
echo ""

# 7. 尝试拉取代码
echo "=========================================="
echo "尝试拉取代码"
echo "=========================================="

# 检查是否有本地提交
if git rev-parse --verify HEAD &> /dev/null; then
    echo "⚠️  检测到本地提交，先检查远程分支..."
    git fetch origin 2>&1 || {
        echo "⚠️  无法连接到远程仓库，可能原因："
        echo "   1. 网络连接问题"
        echo "   2. SSH 密钥未配置"
        echo "   3. 仓库地址不正确"
        echo ""
        echo "   请检查后手动执行:"
        echo "   git fetch origin"
        echo "   git checkout ${GIT_BRANCH}"
        echo "   git pull origin ${GIT_BRANCH}"
        exit 0
    }
    
    # 检查远程分支是否存在
    if git ls-remote --heads origin ${GIT_BRANCH} &> /dev/null; then
        echo "✅ 远程分支 ${GIT_BRANCH} 存在"
        
        # 切换到目标分支
        if [ "$(git branch --show-current 2>/dev/null)" != "${GIT_BRANCH}" ]; then
            echo "📌 切换到分支 ${GIT_BRANCH}..."
            git checkout -b ${GIT_BRANCH} origin/${GIT_BRANCH} 2>/dev/null || \
            git checkout ${GIT_BRANCH} 2>/dev/null || {
                echo "⚠️  无法自动切换分支，请手动执行:"
                echo "   git checkout ${GIT_BRANCH}"
            }
        fi
        
        echo "📥 拉取最新代码..."
        git pull origin ${GIT_BRANCH} || {
            echo "⚠️  拉取失败，可能需要处理冲突"
            echo "   当前状态:"
            git status
        }
    else
        echo "⚠️  远程分支 ${GIT_BRANCH} 不存在"
        echo "   可用的远程分支:"
        git branch -r | head -10
    fi
else
    echo "📥 首次拉取代码..."
    git fetch origin 2>&1 || {
        echo "⚠️  无法连接到远程仓库"
        echo "   请检查网络和 SSH 配置后手动执行:"
        echo "   git fetch origin"
        echo "   git checkout -b ${GIT_BRANCH} origin/${GIT_BRANCH}"
        exit 0
    }
    
    if git ls-remote --heads origin ${GIT_BRANCH} &> /dev/null; then
        git checkout -b ${GIT_BRANCH} origin/${GIT_BRANCH}
        echo "✅ 代码拉取完成"
    else
        echo "⚠️  远程分支 ${GIT_BRANCH} 不存在"
    fi
fi

echo ""
echo "=========================================="
echo "✅ Git 仓库配置完成！"
echo "=========================================="
echo ""
echo "📝 常用命令:"
echo "   查看状态:    git status"
echo "   拉取代码:    git pull origin ${GIT_BRANCH}"
echo "   查看分支:    git branch -a"
echo "   查看远程:    git remote -v"
echo "   查看配置:    git config --local --list"
echo "=========================================="
