#!/bin/bash

# ============================================
# 管理系统 - 统一服务管理脚本
# ============================================
# 使用方法:
#   ./server.sh start      - 启动服务
#   ./server.sh stop       - 停止服务
#   ./server.sh restart    - 重启服务
#   ./server.sh deploy     - 一键部署（编译+构建+启动）
#   ./server.sh status     - 查看服务状态
#   ./server.sh logs       - 查看服务日志
#   ./server.sh build      - 仅构建镜像
# ============================================

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 项目配置
PROJECT_NAME="admin-system"
CONTAINER_NAME="admin-system-backend"
IMAGE_NAME="admin-system-backend"
COMPOSE_FILE="docker-compose.yml"
BACKEND_DIR="backend"

# 从 .env 文件读取端口，如果没有则使用默认值
if [ -f .env ]; then
    SERVER_PORT=$(grep "^SERVER_PORT=" .env 2>/dev/null | cut -d'=' -f2 | tr -d '"' || echo "7701")
else
    SERVER_PORT=7701
fi

# 检查 Docker 是否运行
check_docker() {
    if ! docker info > /dev/null 2>&1; then
        echo -e "${RED}❌ 错误: Docker未运行，请先启动Docker${NC}"
        exit 1
    fi
}

# 检查并设置 docker-compose 命令
setup_compose_cmd() {
    if command -v docker-compose &> /dev/null; then
        COMPOSE_CMD="docker-compose"
    elif docker compose version &> /dev/null 2>&1; then
        COMPOSE_CMD="docker compose"
    else
        echo -e "${RED}❌ 错误: docker-compose（或 docker compose）未安装${NC}"
        exit 1
    fi
}

# 检查 .env 文件
check_env_file() {
    if [ ! -f .env ]; then
        echo -e "${YELLOW}⚠️  警告: .env 文件不存在${NC}"
        if [ -f .env.example ]; then
            echo -e "${YELLOW}   提示: 可以从 .env.example 复制创建: cp .env.example .env${NC}"
        fi
    fi
}

# 构建 Docker 镜像
build_image() {
    echo -e "${BLUE}🔨 构建 Docker 镜像...${NC}"
    
    # 检查 buildx 是否可用且版本足够
    USE_BUILDX=false
    if docker buildx version &> /dev/null 2>&1; then
        BUILDX_VERSION=$(docker buildx version 2>/dev/null | grep -oE '[0-9]+\.[0-9]+' | head -1 || echo "0.0")
        if [ -n "$BUILDX_VERSION" ]; then
            MAJOR=$(echo $BUILDX_VERSION | cut -d. -f1)
            MINOR=$(echo $BUILDX_VERSION | cut -d. -f2)
            
            if [ "$MAJOR" -gt 0 ] || ([ "$MAJOR" -eq 0 ] && [ "$MINOR" -ge 17 ]); then
                USE_BUILDX=true
                echo -e "${GREEN}✅ 使用 buildx 构建（版本: $BUILDX_VERSION）${NC}"
            else
                echo -e "${YELLOW}⚠️  buildx 版本过低（$BUILDX_VERSION），需要 0.17+${NC}"
            fi
        fi
    fi
    
    # 如果 buildx 不可用，使用传统方式构建
    if [ "$USE_BUILDX" = false ]; then
        echo -e "${YELLOW}⚠️  使用传统方式构建镜像（不使用 buildx）...${NC}"
        cd $BACKEND_DIR
        docker build -t $IMAGE_NAME:latest .
        cd ..
    else
        # 使用 docker-compose 构建
        $COMPOSE_CMD build
    fi
}

# 编译后端代码
build_backend() {
    echo -e "${BLUE}📦 编译后端代码...${NC}"
    cd $BACKEND_DIR
    
    if ! command -v go &> /dev/null; then
        echo -e "${YELLOW}⚠️  警告: Go未安装，跳过本地编译，将直接使用Docker构建${NC}"
        cd ..
        return
    fi
    
    echo "正在下载依赖..."
    go env -w GOPROXY=https://goproxy.cn,direct
    go env -w GOSUMDB=sum.golang.google.cn
    
    go mod download
    
    echo "正在编译..."
    if go build -o main .; then
        echo -e "${GREEN}✅ 后端代码编译成功${NC}"
    else
        echo -e "${RED}❌ 后端代码编译失败${NC}"
        cd ..
        exit 1
    fi
    cd ..
}

# 启动服务
start_service() {
    echo -e "${BLUE}🚀 启动服务...${NC}"
    $COMPOSE_CMD up -d
    
    echo -e "${BLUE}⏳ 等待服务启动...${NC}"
    sleep 5
    
    show_status
}

# 停止服务
stop_service() {
    echo -e "${BLUE}🛑 停止服务...${NC}"
    $COMPOSE_CMD down
    echo -e "${GREEN}✅ 服务已停止${NC}"
}

# 重启服务
restart_service() {
    echo -e "${BLUE}🔄 重启服务...${NC}"
    $COMPOSE_CMD restart
    
    echo -e "${BLUE}⏳ 等待服务启动...${NC}"
    sleep 5
    
    show_status
}

# 查看服务状态
show_status() {
    echo ""
    echo -e "${BLUE}📊 服务状态:${NC}"
    $COMPOSE_CMD ps
    
    echo ""
    echo -e "${BLUE}🔍 检查后端服务健康状态...${NC}"
    max_attempts=10
    attempt=0
    while [ $attempt -lt $max_attempts ]; do
        if curl -s http://localhost:$SERVER_PORT/api/auth/userinfo > /dev/null 2>&1; then
            echo -e "${GREEN}✅ 后端服务运行正常${NC}"
            break
        else
            attempt=$((attempt + 1))
            if [ $attempt -eq $max_attempts ]; then
                echo -e "${YELLOW}⚠️  后端服务可能还在启动中，请稍后检查${NC}"
            else
                sleep 2
            fi
        fi
    done
    
    echo ""
    echo -e "${GREEN}=========================================="
    echo "✅ 服务信息"
    echo "==========================================${NC}"
    echo "后端API: http://localhost:$SERVER_PORT"
    echo "查看日志: $COMPOSE_CMD logs -f backend"
    echo "停止服务: ./server.sh stop"
    echo "=========================================="
}

# 查看日志
show_logs() {
    echo -e "${BLUE}📋 查看服务日志（按 Ctrl+C 退出）...${NC}"
    $COMPOSE_CMD logs -f backend
}

# 一键部署
deploy_service() {
    echo -e "${GREEN}=========================================="
    echo "管理系统 - 一键部署"
    echo "==========================================${NC}"
    
    check_docker
    setup_compose_cmd
    check_env_file
    
    # 1. 编译后端代码
    echo ""
    build_backend
    
    # 2. 停止旧服务
    echo ""
    echo -e "${BLUE}🛑 停止旧服务...${NC}"
    $COMPOSE_CMD down 2>/dev/null || true
    
    # 3. 构建镜像
    echo ""
    build_image
    
    # 4. 启动服务
    echo ""
    start_service
    
    echo ""
    echo -e "${GREEN}=========================================="
    echo "✅ 部署完成！"
    echo "==========================================${NC}"
}

# 主函数
main() {
    check_docker
    setup_compose_cmd
    
    case "${1:-}" in
        start)
            check_env_file
            start_service
            ;;
        stop)
            stop_service
            ;;
        restart)
            check_env_file
            restart_service
            ;;
        deploy)
            deploy_service
            ;;
        status)
            show_status
            ;;
        logs)
            show_logs
            ;;
        build)
            build_image
            ;;
        *)
            echo -e "${YELLOW}使用方法: ./server.sh {start|stop|restart|deploy|status|logs|build}${NC}"
            echo ""
            echo "命令说明:"
            echo "  start   - 启动服务"
            echo "  stop    - 停止服务"
            echo "  restart - 重启服务"
            echo "  deploy  - 一键部署（编译+构建+启动）"
            echo "  status  - 查看服务状态"
            echo "  logs    - 查看服务日志"
            echo "  build   - 仅构建镜像"
            exit 1
            ;;
    esac
}

main "$@"
