# .env 文件说明

## 文件位置

`.env` 文件应该放在 **`backend/`** 目录下：

```
admin-system/
└── backend/
    ├── .env          ← 放在这里（需要创建）
    ├── .env.example  ← 示例文件（已提供）
    ├── main.go
    └── ...
```

## 创建 .env 文件

### 方式1：复制示例文件（推荐）

```bash
cd backend
cp .env.example .env
```

### 方式2：手动创建

```bash
cd backend
touch .env
# 然后编辑 .env 文件，复制 .env.example 的内容
```

## 配置说明

### 本地开发环境配置

编辑 `backend/.env` 文件：

```env
# 数据库配置（本地开发）
DB_HOST=localhost        # 本地MySQL地址
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root123456   # 修改为你的MySQL密码
DB_NAME=admin_system
DB_CHARSET=utf8mb4

# Redis配置（本地开发）
REDIS_HOST=localhost     # 本地Redis地址
REDIS_PORT=6379
REDIS_PASSWORD=          # 如果Redis有密码，填写在这里
REDIS_DB=0

# JWT配置
JWT_SECRET=your-secret-key-change-in-production  # 生产环境请修改为强密钥
JWT_EXPIRE_HOURS=72
JWT_REFRESH_THRESHOLD_HOURS=24

# 服务配置
SERVER_PORT=7701
SERVER_MODE=debug        # 开发模式使用 debug，生产环境使用 release
```

### Docker 环境

如果使用 Docker，**不需要**创建 `.env` 文件，配置在 `docker-compose.yml` 中。

## 注意事项

1. **不要提交 .env 文件到 Git**
   - `.env` 文件已添加到 `.gitignore`
   - 只提交 `.env.example` 作为模板

2. **本地开发时**
   - 确保 MySQL 和 Redis 服务已启动
   - 修改 `DB_PASSWORD` 为你的实际密码
   - `DB_HOST` 和 `REDIS_HOST` 使用 `localhost`

3. **Docker 环境时**
   - 不需要 `.env` 文件
   - 配置在 `docker-compose.yml` 中
   - `DB_HOST` 和 `REDIS_HOST` 使用服务名（`mysql`、`redis`）

## 验证配置

启动服务后，查看日志确认配置是否正确加载：

```bash
cd backend
go run main.go
```

如果看到 "MySQL连接成功" 和 "Redis连接成功"，说明配置正确。

