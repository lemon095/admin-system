#!/bin/bash

# 在服务器上配置 GitHub SSH 密钥或使用 HTTPS
# 使用方法: 在服务器上执行 ./setup-github-ssh.sh

set -e

echo "=========================================="
echo "配置 GitHub 访问"
echo "=========================================="
echo ""

# 方案选择
echo "请选择配置方式："
echo "1) 配置 SSH 密钥（推荐，更安全）"
echo "2) 使用 HTTPS + Personal Access Token（更简单）"
echo ""
read -p "请选择 (1/2，默认1): " -n 1 -r
echo ""

if [[ $REPLY =~ ^[2]$ ]]; then
    # 方案2: 使用 HTTPS
    echo ""
    echo "=========================================="
    echo "配置 HTTPS 方式"
    echo "=========================================="
    echo ""
    echo "📝 使用 HTTPS 方式需要 GitHub Personal Access Token"
    echo ""
    echo "步骤："
    echo "1. 访问 https://github.com/settings/tokens"
    echo "2. 点击 'Generate new token (classic)'"
    echo "3. 设置名称和过期时间"
    echo "4. 勾选 'repo' 权限"
    echo "5. 生成并复制 token"
    echo ""
    read -p "请输入你的 GitHub Personal Access Token: " GITHUB_TOKEN
    echo ""
    
    if [ -z "$GITHUB_TOKEN" ]; then
        echo "❌ Token 不能为空"
        exit 1
    fi
    
    # 修改远程仓库地址为 HTTPS
    if [ -d .git ]; then
        if git remote get-url origin &> /dev/null; then
            CURRENT_URL=$(git remote get-url origin)
            # 如果是 SSH URL，转换为 HTTPS
            if [[ $CURRENT_URL == git@github.com:* ]]; then
                HTTPS_URL=$(echo $CURRENT_URL | sed 's|git@github.com:|https://github.com/|' | sed 's|\.git$||')
                HTTPS_URL="${HTTPS_URL}.git"
                # 在 URL 中嵌入 token
                HTTPS_URL_WITH_TOKEN=$(echo $HTTPS_URL | sed "s|https://|https://${GITHUB_TOKEN}@|")
                git remote set-url origin "$HTTPS_URL_WITH_TOKEN"
                echo "✅ 已更新远程仓库地址为 HTTPS（带 token）"
            else
                echo "✅ 远程仓库已配置为 HTTPS"
            fi
        fi
    fi
    
    echo ""
    echo "✅ HTTPS 配置完成！"
    echo ""
    echo "📝 测试连接："
    echo "   git ls-remote origin"
    echo ""
    
else
    # 方案1: 配置 SSH 密钥（默认）
    echo ""
    echo "=========================================="
    echo "配置 SSH 密钥"
    echo "=========================================="
    echo ""
    
    # 1. 检查是否已有 SSH 密钥
    SSH_DIR="$HOME/.ssh"
    SSH_KEY_FILE="$SSH_DIR/id_rsa"
    SSH_PUB_KEY_FILE="$SSH_DIR/id_rsa.pub"
    
    if [ ! -d "$SSH_DIR" ]; then
        mkdir -p "$SSH_DIR"
        chmod 700 "$SSH_DIR"
        echo "✅ 创建 .ssh 目录"
    fi
    
    # 2. 生成或使用现有密钥
    if [ -f "$SSH_KEY_FILE" ]; then
        echo "✅ 发现现有 SSH 密钥: $SSH_KEY_FILE"
        read -p "是否使用现有密钥? (Y/n): " -n 1 -r
        echo ""
        if [[ ! $REPLY =~ ^[Nn]$ ]]; then
            USE_EXISTING=true
        else
            USE_EXISTING=false
        fi
    else
        USE_EXISTING=false
    fi
    
    if [ "$USE_EXISTING" = false ]; then
        echo ""
        echo "🔑 生成新的 SSH 密钥..."
        read -p "请输入邮箱（用于标识密钥）: " SSH_EMAIL
        if [ -z "$SSH_EMAIL" ]; then
            SSH_EMAIL="deploy@admin-system.local"
        fi
        
        ssh-keygen -t rsa -b 4096 -C "$SSH_EMAIL" -f "$SSH_KEY_FILE" -N "" <<< y
        echo "✅ SSH 密钥生成完成"
    fi
    
    # 3. 设置正确的权限
    chmod 600 "$SSH_KEY_FILE" 2>/dev/null || true
    chmod 644 "$SSH_PUB_KEY_FILE" 2>/dev/null || true
    
    # 4. 显示公钥
    echo ""
    echo "=========================================="
    echo "SSH 公钥（请添加到 GitHub）"
    echo "=========================================="
    echo ""
    if [ -f "$SSH_PUB_KEY_FILE" ]; then
        cat "$SSH_PUB_KEY_FILE"
        echo ""
        echo ""
        echo "📋 公钥已显示在上方，请复制整个内容"
        echo ""
    else
        echo "❌ 公钥文件不存在: $SSH_PUB_KEY_FILE"
        exit 1
    fi
    
    # 5. 添加到 SSH config（可选）
    SSH_CONFIG="$SSH_DIR/config"
    if [ ! -f "$SSH_CONFIG" ] || ! grep -q "Host github.com" "$SSH_CONFIG" 2>/dev/null; then
        echo ""
        echo "📝 配置 SSH config..."
        cat >> "$SSH_CONFIG" << EOF

Host github.com
    HostName github.com
    User git
    IdentityFile $SSH_KEY_FILE
    StrictHostKeyChecking no
EOF
        chmod 600 "$SSH_CONFIG"
        echo "✅ SSH config 已配置"
    fi
    
    # 6. 测试连接
    echo ""
    echo "=========================================="
    echo "测试 GitHub 连接"
    echo "=========================================="
    echo ""
    echo "⚠️  请先完成以下步骤："
    echo ""
    echo "1. 访问 https://github.com/settings/keys"
    echo "2. 点击 'New SSH key'"
    echo "3. 标题填写: Server Deploy Key"
    echo "4. 将上面的公钥内容粘贴到 'Key' 字段"
    echo "5. 点击 'Add SSH key'"
    echo ""
    read -p "完成上述步骤后，按 Enter 继续测试连接..."
    echo ""
    
    echo "🔍 测试 SSH 连接..."
    if ssh -T git@github.com 2>&1 | grep -q "successfully authenticated"; then
        echo "✅ SSH 连接成功！"
    elif ssh -T git@github.com 2>&1 | grep -q "Permission denied"; then
        echo "❌ SSH 认证失败"
        echo "   请检查："
        echo "   1. 公钥是否已正确添加到 GitHub"
        echo "   2. 密钥文件权限是否正确"
        echo ""
        echo "   可以手动测试："
        echo "   ssh -T git@github.com"
    else
        echo "⚠️  连接测试结果："
        ssh -T git@github.com 2>&1 || true
    fi
    
    echo ""
    echo "✅ SSH 配置完成！"
fi

echo ""
echo "=========================================="
echo "配置完成"
echo "=========================================="
echo ""
echo "📝 现在可以尝试克隆或拉取代码："
echo "   git clone git@github.com:lemon095/admin-system.git"
echo "   或"
echo "   git pull origin aki"
echo "=========================================="
