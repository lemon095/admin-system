#!/bin/bash

# ============================================
# 管理系统 - 一键部署脚本
# ============================================
# 功能：
#   1. 编译后端代码（可选，如果Go已安装）
#   2. 停止旧服务
#   3. 构建Docker镜像
#   4. 启动服务
#   5. 检查服务状态
# ============================================

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# 配置
PROJECT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BACKEND_DIR="${PROJECT_DIR}/backend"
CONTAINER_NAME="admin-system-backend"
SERVER_PORT="${SERVER_PORT:-9001}"

# 检查Docker是否运行
check_docker() {
    if ! docker info > /dev/null 2>&1; then
        echo -e "${RED}❌ 错误: Docker未运行，请先启动Docker${NC}"
        exit 1
    fi
}

# 检查并设置docker-compose命令
setup_compose_cmd() {
    if docker compose version &> /dev/null 2>&1; then
        COMPOSE_CMD="docker compose"
    elif command -v docker-compose &> /dev/null; then
        COMPOSE_CMD="docker-compose"
    else
        echo -e "${RED}❌ 错误: docker-compose（或 docker compose）未安装${NC}"
        exit 1
    fi
}

# 编译后端代码（可选）
build_backend_optional() {
    if ! command -v go &> /dev/null; then
        echo -e "${YELLOW}⚠️  Go未安装，跳过本地编译（将使用Docker构建）${NC}"
        return 0
    fi

    echo -e "${BLUE}📦 编译后端代码...${NC}"
    cd "${BACKEND_DIR}"
    
    # 设置Go代理
    go env -w GOPROXY=https://goproxy.cn,direct 2>/dev/null || true
    go env -w GOSUMDB=sum.golang.google.cn 2>/dev/null || true
    
    # 下载依赖
    echo "   下载依赖..."
    go mod download
    
    # 编译
    echo "   编译中..."
    if go build -o main .; then
        echo -e "${GREEN}✅ 后端代码编译成功${NC}"
    else
        echo -e "${YELLOW}⚠️  本地编译失败，将使用Docker构建${NC}"
    fi
    
    cd "${PROJECT_DIR}"
}

# 停止旧服务
stop_services() {
    echo -e "${BLUE}🛑 停止旧服务...${NC}"
    $COMPOSE_CMD down 2>/dev/null || true
    echo -e "${GREEN}✅ 旧服务已停止${NC}"
}

# 构建Docker镜像
build_image() {
    echo -e "${BLUE}🔨 构建Docker镜像...${NC}"
    
    # 构建镜像
    $COMPOSE_CMD build --no-cache
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ 镜像构建成功${NC}"
    else
        echo -e "${RED}❌ 镜像构建失败${NC}"
        exit 1
    fi
}

# 启动服务
start_services() {
    echo -e "${BLUE}🚀 启动服务...${NC}"
    $COMPOSE_CMD up -d
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ 服务启动成功${NC}"
    else
        echo -e "${RED}❌ 服务启动失败${NC}"
        exit 1
    fi
}

# 等待服务启动
wait_for_service() {
    echo -e "${BLUE}⏳ 等待服务启动...${NC}"
    sleep 5
    
    local max_attempts=15
    local attempt=0
    
    while [ $attempt -lt $max_attempts ]; do
        if curl -s -f "http://localhost:${SERVER_PORT}/api/auth/userinfo" > /dev/null 2>&1; then
            echo -e "${GREEN}✅ 后端服务运行正常${NC}"
            return 0
        fi
        
        attempt=$((attempt + 1))
        if [ $attempt -lt $max_attempts ]; then
            echo "   等待中... (${attempt}/${max_attempts})"
            sleep 2
        fi
    done
    
    echo -e "${YELLOW}⚠️  服务可能还在启动中，请稍后检查${NC}"
    return 1
}

# 显示服务状态
show_status() {
    echo ""
    echo -e "${BLUE}📊 服务状态:${NC}"
    $COMPOSE_CMD ps
    
    echo ""
    echo -e "${GREEN}=========================================="
    echo "✅ 部署完成！"
    echo "==========================================${NC}"
    echo ""
    echo -e "${BLUE}📌 服务地址:${NC}"
    echo "   后端API: http://localhost:${SERVER_PORT}"
    echo "   MySQL:   localhost:3306"
    echo "   Redis:   localhost:6379"
    echo ""
    echo -e "${BLUE}📝 常用命令:${NC}"
    echo "   查看日志:    $COMPOSE_CMD logs -f backend"
    echo "   查看状态:    $COMPOSE_CMD ps"
    echo "   停止服务:    $COMPOSE_CMD down"
    echo "   重启服务:    $COMPOSE_CMD restart backend"
    echo "=========================================="
}

# 主函数
main() {
    echo -e "${GREEN}=========================================="
    echo "管理系统 - 一键部署脚本"
    echo "==========================================${NC}"
    echo ""
    
    # 检查环境
    check_docker
    setup_compose_cmd
    
    # 步骤1: 可选编译后端代码
    echo ""
    build_backend_optional
    
    # 步骤2: 停止旧服务
    echo ""
    stop_services
    
    # 步骤3: 构建镜像
    echo ""
    build_image
    
    # 步骤4: 启动服务
    echo ""
    start_services
    
    # 步骤5: 等待并检查服务
    echo ""
    wait_for_service
    
    # 显示状态
    show_status
}

# 执行主函数
main "$@"