#!/bin/bash

# 在服务器上直接修复 docker-compose.yml
# 使用方法: 在服务器上执行此脚本

set -e

COMPOSE_FILE="/opt/admin-system/docker-compose.yml"

if [ ! -f "$COMPOSE_FILE" ]; then
    echo "❌ 错误: $COMPOSE_FILE 不存在"
    exit 1
fi

echo "=========================================="
echo "修复 docker-compose.yml"
echo "=========================================="

# 备份原文件
cp "$COMPOSE_FILE" "${COMPOSE_FILE}.backup.$(date +%Y%m%d_%H%M%S)"
echo "✅ 已备份原文件"

# 1. 移除 version 字段（如果存在）
if grep -q "^version:" "$COMPOSE_FILE"; then
    sed -i '/^version:/d' "$COMPOSE_FILE"
    echo "✅ 已移除 version 字段"
fi

# 2. 注释掉所有 env_file 配置
sed -i '/env_file:/,/^[[:space:]]*-/ s/^/# /' "$COMPOSE_FILE"
echo "✅ 已注释掉 env_file 配置"

echo ""
echo "=========================================="
echo "修复完成！"
echo "=========================================="
echo ""
echo "现在可以重新运行部署脚本:"
echo "  ./deploy.sh"
echo ""
