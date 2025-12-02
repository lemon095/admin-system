# 快速启动指南

## 使用Docker一键启动（推荐）

### 1. 启动所有服务

```bash
# 方式1: 使用启动脚本
chmod +x start.sh
./start.sh

# 方式2: 直接使用docker-compose
docker-compose up -d
```

### 2. 查看服务状态

```bash
docker-compose ps
```

### 3. 查看后端日志

```bash
docker-compose logs -f backend
```

### 4. 停止服务

```bash
docker-compose down
```

## 访问系统

### 前端（需要单独启动）

```bash
cd frontend
npm install
npm run dev
```

前端将运行在: http://localhost:3000

### 后端API

后端API运行在: http://localhost:7701

### 测试登录接口

```bash
curl -X POST http://localhost:7701/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "chuchangkeji",
    "password": "chuchangkeji666"
  }'
```

## 默认账号

- **用户名**: chuchangkeji
- **密码**: chuchangkeji666
- **角色**: 超级管理员

## 常见问题

### 1. 端口被占用

如果7701、3306、6379端口被占用，可以修改`docker-compose.yml`中的端口映射。

### 2. 数据库连接失败

确保MySQL容器已完全启动（可能需要等待10-20秒）。

### 3. 前端无法连接后端

检查`frontend/vite.config.js`中的proxy配置，确保后端地址正确。

## 开发模式

### 后端本地开发

```bash
cd backend
go mod tidy
go mod download
cp .env.example .env
# 编辑.env文件，配置数据库连接
go run main.go
```

### 前端本地开发

```bash
cd frontend
npm install
npm run dev
```

