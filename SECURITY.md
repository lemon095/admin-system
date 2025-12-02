# 安全配置说明

## 敏感信息保护

为了保护敏感信息（如数据库密码、JWT密钥等），项目使用环境变量来管理配置。

## 配置文件说明

### 1. docker-compose.yml

**已更新为使用环境变量**，不再包含硬编码的敏感信息。

所有敏感配置都从 `.env` 文件或环境变量中读取：

```yaml
environment:
  DB_PASSWORD: ${DB_PASSWORD:-root123456}  # 从环境变量读取，有默认值
  JWT_SECRET: ${JWT_SECRET:-default-secret}
```

### 2. .env 文件（项目根目录）

**用于 Docker 环境**，包含所有敏感配置。

**重要**：
- ✅ `.env` 文件已添加到 `.gitignore`，**不会被提交到 Git**
- ✅ 只提交 `.env.example` 作为模板
- ⚠️ **生产环境必须修改所有默认密码和密钥**

### 3. backend/.env 文件

**用于本地开发环境**，配置本地 MySQL 和 Redis 连接。

## 安全最佳实践

### 1. 首次部署

```bash
# 1. 复制环境变量模板
cp .env.example .env

# 2. 编辑 .env 文件，修改所有敏感信息
# - 修改 MYSQL_ROOT_PASSWORD（强密码）
# - 修改 DB_PASSWORD（强密码）
# - 修改 JWT_SECRET（至少32字符的随机字符串）
# - 修改 REDIS_PASSWORD（如果启用）

# 3. 确保 .env 文件权限
chmod 600 .env

# 4. 部署
./deploy.sh
```

### 2. 生产环境配置

**必须修改的配置**：

```env
# MySQL - 使用强密码
MYSQL_ROOT_PASSWORD=your-strong-password-here-min-16-chars
DB_PASSWORD=your-strong-password-here-min-16-chars

# JWT - 使用强随机字符串（至少32字符）
JWT_SECRET=your-very-long-random-secret-key-at-least-32-characters-long

# Redis - 如果启用密码
REDIS_PASSWORD=your-redis-password-here
```

**生成强密码的方法**：

```bash
# 生成随机密码（32字符）
openssl rand -base64 32

# 生成随机密码（64字符）
openssl rand -base64 64
```

### 3. 文件权限

确保 `.env` 文件权限正确：

```bash
chmod 600 .env          # 只有所有者可读写
chmod 600 backend/.env  # 本地开发环境
```

### 4. 不要提交敏感信息

**永远不要**：
- ❌ 提交 `.env` 文件到 Git
- ❌ 在代码中硬编码密码
- ❌ 在 `docker-compose.yml` 中硬编码敏感信息
- ❌ 在日志中输出密码

**应该**：
- ✅ 使用环境变量
- ✅ 只提交 `.env.example`
- ✅ 使用密钥管理服务（生产环境）
- ✅ 定期轮换密码和密钥

## 环境变量优先级

配置读取优先级（从高到低）：

1. **系统环境变量** - 最高优先级
2. **`.env` 文件** - 项目根目录（Docker）或 `backend/.env`（本地开发）
3. **docker-compose.yml 中的默认值** - 仅作为开发环境的后备

## 检查清单

部署前检查：

- [ ] `.env` 文件已创建并配置
- [ ] 所有默认密码已修改
- [ ] JWT_SECRET 已设置为强随机字符串
- [ ] `.env` 文件权限设置为 600
- [ ] `.env` 文件已添加到 `.gitignore`
- [ ] 没有在代码中硬编码敏感信息
- [ ] `docker-compose.yml` 使用环境变量

## 密钥管理（生产环境推荐）

### 方案1：使用 Docker Secrets

```yaml
services:
  backend:
    secrets:
      - db_password
      - jwt_secret
secrets:
  db_password:
    file: ./secrets/db_password.txt
  jwt_secret:
    file: ./secrets/jwt_secret.txt
```

### 方案2：使用密钥管理服务

- AWS Secrets Manager
- HashiCorp Vault
- Azure Key Vault
- Google Secret Manager

### 方案3：使用 CI/CD 环境变量

在 CI/CD 平台（如 GitHub Actions、GitLab CI）中配置环境变量，部署时自动注入。

## 紧急情况处理

如果发现敏感信息已泄漏：

1. **立即修改所有密码和密钥**
2. **撤销已泄漏的 JWT token**（如果可能）
3. **检查访问日志**，确认是否有未授权访问
4. **通知相关团队**
5. **更新安全策略**

## 相关文件

- `.env.example` - 环境变量模板（可提交）
- `.env` - 实际环境变量（不提交）
- `docker-compose.yml` - 使用环境变量
- `backend/.env.example` - 本地开发模板
- `backend/.env` - 本地开发配置（不提交）

