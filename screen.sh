#!/bin/bash

# 设置路径
FRONTEND_DIR="/usr/local/admin-system"
BACKEND_DIR="/usr/local/admin-system/backend"

# -------------------------------
# 重启 frontend
# -------------------------------

# 关闭 frontend screen 会话（如果存在）
screen -S frontend -X quit 2>/dev/null || echo "frontend screen 不存在，直接创建"

# 进入 frontend 目录
cd "$FRONTEND_DIR" || { echo "无法进入 $FRONTEND_DIR"; exit 1; }

# 启动 frontend screen 会话
screen -dmS frontend bash -c "serve -s dist -l tcp://0.0.0.0:9002; exec bash"

echo "frontend 已重启"

# -------------------------------
# 编译并重启 backend
# -------------------------------

# 进入 backend 目录
cd "$BACKEND_DIR" || { echo "无法进入 $BACKEND_DIR"; exit 1; }

# 编译 Go 程序
echo "正在编译 backend..."
GOOS=linux GOARCH=amd64 go build -o main main.go || { echo "Go 编译失败"; exit 1; }

# 关闭 backend screen 会话（如果存在）
screen -S backend -X quit 2>/dev/null || echo "backend screen 不存在，直接创建"

# 启动 backend screen 会话
screen -dmS backend /bin/bash -c "/usr/local/admin-system/backend/main >> /var/log/app.log 2>&1"

echo "backend 已重启"
