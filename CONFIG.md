# 数据库配置说明

## 配置位置

数据库配置有两个位置，根据运行方式选择：

### 1. Docker 环境配置（推荐）

**文件位置**: `docker-compose.yml`

**配置内容**:
```yaml
services:
  mysql:
    environment:
      MYSQL_ROOT_PASSWORD: root123456  # MySQL root密码
      MYSQL_DATABASE: admin_system      # 数据库名
      MYSQL_CHARSET: utf8mb4
      MYSQL_COLLATION: utf8mb4_unicode_ci
    ports:
      - "3306:3306"                     # 端口映射

  backend:
    environment:
      DB_HOST: mysql                    # 数据库主机（Docker网络内使用服务名）
      DB_PORT: 3306                     # 数据库端口
      DB_USER: root                     # 数据库用户名
      DB_PASSWORD: root123456           # 数据库密码
      DB_NAME: admin_system             # 数据库名
      DB_CHARSET: utf8mb4
      
      REDIS_HOST: redis                 # Redis主机
      REDIS_PORT: 6379                  # Redis端口
      REDIS_PASSWORD: ""                # Redis密码（空表示无密码）
      REDIS_DB: 0                        # Redis数据库编号
```

**修改方式**: 直接编辑 `docker-compose.yml` 文件中的环境变量

---

### 2. 本地开发环境配置

**文件位置**: `backend/.env`

**创建方式**:
```bash
cd backend
cp .env.example .env
```

**配置内容** (`.env.example`):
```env
# 数据库配置
DB_HOST=localhost        # 本地开发使用localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root123456
DB_NAME=admin_system
DB_CHARSET=utf8mb4

# Redis配置
REDIS_HOST=localhost     # 本地开发使用localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT配置
JWT_SECRET=your-secret-key-change-in-production
JWT_EXPIRE_HOURS=72
JWT_REFRESH_THRESHOLD_HOURS=24

# 服务配置
SERVER_PORT=7701
SERVER_MODE=debug        # 开发模式使用debug
```

**修改方式**: 编辑 `backend/.env` 文件

---

## 配置说明

### MySQL 配置

| 配置项 | 说明 | 默认值 |
|--------|------|--------|
| `DB_HOST` | 数据库主机地址 | `mysql` (Docker) / `localhost` (本地) |
| `DB_PORT` | 数据库端口 | `3306` |
| `DB_USER` | 数据库用户名 | `root` |
| `DB_PASSWORD` | 数据库密码 | `root123456` |
| `DB_NAME` | 数据库名称 | `admin_system` |
| `DB_CHARSET` | 字符集 | `utf8mb4` |

### Redis 配置

| 配置项 | 说明 | 默认值 |
|--------|------|--------|
| `REDIS_HOST` | Redis主机地址 | `redis` (Docker) / `localhost` (本地) |
| `REDIS_PORT` | Redis端口 | `6379` |
| `REDIS_PASSWORD` | Redis密码 | 空（无密码） |
| `REDIS_DB` | Redis数据库编号 | `0` |

### JWT 配置

| 配置项 | 说明 | 默认值 |
|--------|------|--------|
| `JWT_SECRET` | JWT密钥 | `your-secret-key-change-in-production` |
| `JWT_EXPIRE_HOURS` | Token有效期（小时） | `72` (3天) |
| `JWT_REFRESH_THRESHOLD_HOURS` | 自动刷新阈值（小时） | `24` |

### 服务配置

| 配置项 | 说明 | 默认值 |
|--------|------|--------|
| `SERVER_PORT` | 服务端口 | `7701` |
| `SERVER_MODE` | 运行模式 | `release` (生产) / `debug` (开发) |

---

## 配置优先级

配置读取优先级（从高到低）：
1. 环境变量（系统环境变量）
2. `.env` 文件（本地开发）
3. `docker-compose.yml` 中的环境变量（Docker环境）
4. 代码中的默认值

---

## 修改配置后的操作

### Docker 环境

修改 `docker-compose.yml` 后需要重启服务：

```bash
# 方式1: 使用部署脚本
./deploy.sh

# 方式2: 手动重启
docker-compose down
docker-compose up -d
```

### 本地开发环境

修改 `backend/.env` 后需要重启服务：

```bash
# 停止当前服务（Ctrl+C）
# 重新启动
go run main.go
```

---

## 安全建议

### 生产环境

1. **修改默认密码**: 
   - 修改 `MYSQL_ROOT_PASSWORD` 为强密码
   - 修改 `DB_PASSWORD` 为强密码

2. **修改JWT密钥**:
   - 使用强随机字符串作为 `JWT_SECRET`
   - 建议长度至少32字符

3. **使用环境变量**:
   - 不要将敏感信息提交到代码仓库
   - 使用环境变量或密钥管理服务

4. **限制网络访问**:
   - 生产环境不要暴露MySQL和Redis端口到公网
   - 使用防火墙规则限制访问

---

## 常见问题

### 1. 连接数据库失败

**检查项**:
- 数据库服务是否启动
- 主机地址是否正确（Docker环境使用服务名，本地使用localhost）
- 端口是否正确
- 用户名密码是否正确
- 数据库是否存在

### 2. 连接Redis失败

**检查项**:
- Redis服务是否启动
- 主机地址和端口是否正确
- 密码是否正确（如果设置了密码）

### 3. 配置不生效

**解决方法**:
- 确保修改了正确的配置文件
- 重启服务使配置生效
- 检查环境变量优先级

---

## 配置文件位置总结

```
admin-system/
├── docker-compose.yml          # Docker环境配置（MySQL、Redis、后端环境变量）
└── backend/
    ├── .env.example             # 环境变量示例文件
    ├── .env                     # 本地开发环境变量（需要创建）
    └── config/
        └── config.go            # 配置读取代码
```

