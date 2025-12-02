# 管理系统 (Admin System)

一个前后端分离的后台管理系统项目。

## 项目结构

```
admin-system/
├── backend/          # Go后端服务
│   ├── config/       # 配置管理
│   ├── controller/   # 控制器层
│   ├── database/     # 数据库相关
│   ├── middleware/   # 中间件
│   ├── model/        # 数据模型
│   ├── router/       # 路由配置
│   ├── service/      # 业务逻辑层
│   ├── utils/        # 工具函数
│   ├── main.go       # 入口文件
│   └── Dockerfile    # Docker构建文件
├── frontend/         # React前端项目
│   ├── src/
│   │   ├── components/  # 组件
│   │   ├── context/     # 上下文
│   │   ├── pages/       # 页面
│   │   └── App.jsx      # 主应用
│   └── package.json
├── docs/             # 文档
│   └── database/     # 数据库设计文档
└── docker-compose.yml # Docker编排文件
```

## 功能特性

- ✅ 用户登录认证（JWT Token）
- ✅ Token自动刷新机制（3天有效期，24小时内自动刷新）
- ✅ 用户权限管理（超级管理员、管理员、普通用户）
- ✅ 默认超级管理员账号自动创建
- ✅ 前后端分离架构
- ✅ Docker容器化部署

## 技术栈

### 后端
- Go 1.21+
- Gin Web框架
- MySQL 8.0
- Redis 7
- JWT认证

### 前端
- React 18
- Vite
- React Router
- Axios

## 快速开始

### 前置要求

- Docker & Docker Compose
- Go 1.21+ (本地开发，可选)
- Node.js 18+ (前端开发)

### 一键部署脚本（推荐）

项目提供了多个便捷脚本：

#### 1. 完整部署（编译+构建+启动）
```bash
./deploy.sh
```
这个脚本会：
- ✅ 编译后端代码
- ✅ 构建Docker镜像
- ✅ 停止旧服务
- ✅ 启动新服务
- ✅ 检查服务状态

#### 2. 快速启动（首次使用）
```bash
./start.sh
```
直接启动所有服务（如果镜像已构建）

#### 3. 快速重启（不重新构建）
```bash
./restart.sh
```
仅重启服务，不重新构建镜像

#### 4. 仅构建镜像
```bash
./build.sh
```
只构建Docker镜像，不启动服务

#### 5. 停止服务
```bash
./stop.sh
```
停止所有服务

### 使用Docker Compose（手动）

1. 启动所有服务（MySQL、Redis、后端）
```bash
docker-compose up -d
```

2. 查看服务状态
```bash
docker-compose ps
```

3. 查看日志
```bash
docker-compose logs -f backend
```

4. 停止服务
```bash
docker-compose down
```

### 本地开发

#### 后端开发

1. 进入后端目录
```bash
cd backend
```

2. 安装依赖
```bash
go mod download
```

3. 复制环境变量文件
```bash
cp .env.example .env
```

4. 修改`.env`文件中的配置（确保MySQL和Redis已启动）

5. 运行服务
```bash
go run main.go
```

后端服务将运行在 `http://localhost:7701`

#### 前端开发

1. 进入前端目录
```bash
cd frontend
```

2. 安装依赖
```bash
npm install
```

3. 启动开发服务器
```bash
npm run dev
```

前端服务将运行在 `http://localhost:3000`

## 默认账号

系统首次启动时会自动创建默认超级管理员账号：

- **用户名**: `chuchangkeji`
- **密码**: `chuchangkeji666`
- **角色**: 超级管理员
- **头像ID**: `default-avatar-001`

> 注意：只有在数据库中没有任何用户时才会创建默认账号。

## API接口

### 登录接口
```
POST /api/auth/login
Content-Type: application/json

{
  "username": "chuchangkeji",
  "password": "chuchangkeji666"
}
```

响应：
```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "chuchangkeji",
      "avatar_id": "default-avatar-001",
      "role": 1,
      "status": 1,
      "created_at": "2024-01-01T00:00:00Z"
    },
    "expire_at": 1704067200
  }
}
```

### 刷新Token接口
```
POST /api/auth/refresh
Authorization: Bearer <token>
```

### 获取用户信息
```
GET /api/auth/userinfo
Authorization: Bearer <token>
```

## 数据库设计

### 用户表 (users)

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 用户ID（主键） |
| username | VARCHAR(50) | 用户名（唯一） |
| avatar_id | VARCHAR(255) | 头像ID/URL |
| role | TINYINT | 用户身份（1-超级管理员，2-管理员，3-普通用户） |
| password | VARCHAR(32) | 密码（MD5值） |
| status | TINYINT | 状态（1-启用，0-禁用） |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |
| created_by | BIGINT | 创建者ID |
| last_login_at | DATETIME | 最后登录时间 |
| last_login_ip | VARCHAR(50) | 最后登录IP |

## 环境变量配置

后端环境变量（`.env`文件）：

```env
# 数据库配置
DB_HOST=mysql
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root123456
DB_NAME=admin_system
DB_CHARSET=utf8mb4

# Redis配置
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT配置
JWT_SECRET=your-secret-key-change-in-production
JWT_EXPIRE_HOURS=72
JWT_REFRESH_THRESHOLD_HOURS=24

# 服务配置
SERVER_PORT=7701
SERVER_MODE=release
```

## Token机制

- Token有效期：72小时（3天）
- 自动刷新阈值：24小时（距离过期时间小于24小时时自动刷新）
- Token存储在Redis中，支持token撤销
- 前端自动检测响应头中的`X-New-Token`并更新本地存储

## 开发说明

### 后端分层架构

- **Controller**: 处理HTTP请求和响应
- **Service**: 业务逻辑处理
- **Model**: 数据模型定义
- **Database**: 数据库操作和初始化
- **Middleware**: 中间件（认证、CORS等）
- **Utils**: 工具函数（JWT、密码加密等）

### 前端架构

- **Pages**: 页面组件
- **Components**: 可复用组件
- **Context**: 全局状态管理（认证状态）
- **Axios拦截器**: 自动处理token刷新

## 注意事项

1. 生产环境请修改JWT_SECRET为强随机字符串
2. 生产环境请修改数据库密码
3. 建议使用HTTPS协议
4. 定期备份数据库数据

## 许可证

MIT
