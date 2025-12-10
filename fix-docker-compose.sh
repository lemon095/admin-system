#!/bin/bash

# 修复服务器上的 docker-compose.yml 文件
# 使用方法: 在服务器上执行 ./fix-docker-compose.sh

set -e

echo "=========================================="
echo "修复 docker-compose.yml 配置"
echo "=========================================="
echo ""

COMPOSE_FILE="docker-compose.yml"

if [ ! -f "$COMPOSE_FILE" ]; then
    echo "❌ 错误: $COMPOSE_FILE 文件不存在"
    exit 1
fi

echo "📝 检查并修复 $COMPOSE_FILE..."
echo ""

# 1. 检查并移除 version 字段
if grep -q "^version:" "$COMPOSE_FILE"; then
    echo "⚠️  发现过时的 version 字段，正在移除..."
    # 使用 sed 删除 version 行及其后的空行
    sed -i '/^version:/d' "$COMPOSE_FILE"
    # 删除可能的空行
    sed -i '/^$/N;/^\n$/d' "$COMPOSE_FILE"
    echo "✅ 已移除 version 字段"
else
    echo "✅ 未发现 version 字段"
fi

# 2. 检查并修复 env_file 配置
if grep -q "env_file:" "$COMPOSE_FILE"; then
    echo "⚠️  发现 env_file 配置，正在检查..."
    
    # 检查是否引用了 backend/.env
    if grep -q "backend/.env" "$COMPOSE_FILE"; then
        echo "⚠️  发现 backend/.env 引用，正在修复..."
        # 注释掉 env_file 配置
        sed -i '/env_file:/,/^[^ ]/ { /env_file:/s/^/# /; /- /s/^/# /; }' "$COMPOSE_FILE"
        # 更精确的替换：注释掉 env_file 块
        sed -i '/env_file:/,/^[[:space:]]*-/ { s/^/# /; }' "$COMPOSE_FILE"
        echo "✅ 已注释掉 env_file 配置（使用 environment 中的默认值）"
    elif grep -q "^\s*-\s*\.env" "$COMPOSE_FILE"; then
        echo "⚠️  发现 .env 引用，正在注释..."
        sed -i '/env_file:/,/^[[:space:]]*-/ { s/^/# /; }' "$COMPOSE_FILE"
        echo "✅ 已注释掉 env_file 配置"
    else
        echo "✅ env_file 配置看起来正常"
    fi
else
    echo "✅ 未发现 env_file 配置（使用 environment 默认值）"
fi

echo ""
echo "=========================================="
echo "修复完成"
echo "=========================================="
echo ""
echo "📋 当前 docker-compose.yml 配置摘要:"
echo ""

# 显示关键配置
if grep -q "# env_file:" "$COMPOSE_FILE"; then
    echo "✅ env_file: 已注释（使用 environment 默认值）"
elif grep -q "env_file:" "$COMPOSE_FILE"; then
    echo "⚠️  env_file: 仍在使用"
    grep "env_file:" "$COMPOSE_FILE" | head -3
fi

if grep -q "^version:" "$COMPOSE_FILE"; then
    echo "⚠️  version: 仍存在"
else
    echo "✅ version: 已移除"
fi

echo ""
echo "📝 如果仍有问题，可以手动检查文件:"
echo "   cat $COMPOSE_FILE | grep -A 5 'env_file:'"
echo ""
