#!/bin/bash

# SSH连接测试脚本

SERVER_IP="39.105.136.116"
SERVER_USER="root"
SERVER_PASSWORD="huojian123456!@#\$%^"
SSH_KEY="${HOME}/Desktop/chuchang/chuchang.pem"  # SSH密钥文件路径

echo "=========================================="
echo "SSH连接测试"
echo "=========================================="
echo "服务器: ${SERVER_USER}@${SERVER_IP}"
echo ""

# 检查SSH密钥文件
if [ -f "${SSH_KEY}" ]; then
    echo "✅ 找到SSH密钥文件: ${SSH_KEY}"
    chmod 600 "${SSH_KEY}" 2>/dev/null
else
    echo "⚠️  未找到SSH密钥文件: ${SSH_KEY}"
    echo "   请确保密钥文件存在，或修改脚本中的SSH_KEY路径"
fi

echo ""

# 方法1: 使用SSH密钥（推荐）
echo "方法1: 使用SSH密钥测试..."
if [ -f "${SSH_KEY}" ]; then
    if ssh -i "${SSH_KEY}" -o StrictHostKeyChecking=no -o ConnectTimeout=10 ${SERVER_USER}@${SERVER_IP} "echo '连接成功'; uname -a" 2>&1; then
        echo "✅ SSH密钥连接成功"
    else
        echo "❌ SSH密钥连接失败"
        echo "   请检查："
        echo "   1. 密钥文件路径是否正确"
        echo "   2. 密钥文件权限是否正确（应该是600）"
        echo "   3. 密钥是否匹配服务器"
    fi
else
    echo "⚠️  跳过（密钥文件不存在）"
fi

echo ""

# 方法2: 使用sshpass（如果服务器允许密码认证）
echo "方法2: 使用密码测试（如果服务器允许）..."
if command -v sshpass &> /dev/null; then
    if sshpass -p "${SERVER_PASSWORD}" ssh -o StrictHostKeyChecking=no -o ConnectTimeout=10 ${SERVER_USER}@${SERVER_IP} "echo '连接成功'" 2>&1 | grep -q "连接成功"; then
        echo "✅ 密码认证成功"
    else
        echo "❌ 密码认证失败（服务器可能只允许密钥认证）"
    fi
else
    echo "⚠️  sshpass未安装"
fi

echo ""

# 方法3: 手动SSH连接
echo "方法3: 手动SSH连接..."
if [ -f "${SSH_KEY}" ]; then
    echo "使用密钥连接:"
    echo "  ssh -i ${SSH_KEY} ${SERVER_USER}@${SERVER_IP}"
    echo ""
    read -p "是否现在尝试？(y/n) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        ssh -i "${SSH_KEY}" ${SERVER_USER}@${SERVER_IP}
    fi
else
    echo "执行命令: ssh ${SERVER_USER}@${SERVER_IP}"
    echo "（需要密钥文件或密码）"
fi

echo ""
echo "方法3: 测试端口连通性..."
if command -v nc &> /dev/null; then
    echo "测试端口22..."
    if nc -z -w 5 ${SERVER_IP} 22 2>/dev/null; then
        echo "✅ 端口22可访问"
    else
        echo "❌ 端口22无法访问"
    fi
else
    echo "⚠️  nc (netcat) 未安装，无法测试端口"
fi

echo ""
echo "方法4: 使用telnet测试（如果可用）..."
if command -v telnet &> /dev/null; then
    echo "测试端口22..."
    timeout 5 telnet ${SERVER_IP} 22 2>&1 | head -3
else
    echo "⚠️  telnet未安装"
fi

echo ""
echo "=========================================="
echo "如果连接失败，可能的原因："
echo "1. 服务器SSH服务未启动"
echo "2. 防火墙阻止了SSH连接"
echo "3. 服务器IP地址错误"
echo "4. 网络问题（需要VPN或跳板机）"
echo "5. SSH端口不是22"
echo "=========================================="


