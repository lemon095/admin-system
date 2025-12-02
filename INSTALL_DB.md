# MySQL和Redis安装指南

## 快速安装

### 方式1: 远程自动安装（推荐）

```bash
./install-db-remote.sh
```

这个脚本会：
- 自动SSH连接到服务器
- 上传安装脚本
- 在服务器上执行安装
- 配置MySQL和Redis允许Docker容器访问

### 方式2: 手动在服务器上安装

**步骤1: SSH登录服务器**
```bash
ssh -i ~/Desktop/chuchang/chuchang.pem root@39.105.136.116
```

**步骤2: 上传安装脚本**
```bash
# 在本地执行
scp -i ~/Desktop/chuchang/chuchang.pem install-mysql-redis.sh root@39.105.136.116:/tmp/
```

**步骤3: 在服务器上执行**
```bash
# 在服务器上执行
chmod +x /tmp/install-mysql-redis.sh
/tmp/install-mysql-redis.sh
```

## 安装脚本功能

`install-mysql-redis.sh` 脚本会自动：

1. **检测操作系统**（CentOS/RHEL/Ubuntu/Debian）
2. **安装MySQL**
   - 自动安装MySQL服务器
   - 启动并设置开机自启
   - 创建数据库 `admin_system`
3. **安装Redis**
   - 自动安装Redis服务器
   - 启动并设置开机自启
4. **配置远程访问**
   - MySQL允许从Docker容器连接
   - Redis允许从Docker容器连接
5. **配置防火墙**
   - 开放MySQL端口3306
   - 开放Redis端口6379

## 安装后配置

### 1. 设置MySQL root密码

如果MySQL是首次安装，需要设置root密码：

```bash
# 在服务器上执行
mysql_secure_installation
```

或者直接设置：

```bash
mysql -u root
ALTER USER 'root'@'localhost' IDENTIFIED BY 'your-password';
FLUSH PRIVILEGES;
```

### 2. 配置.env文件

在项目根目录创建或编辑 `.env` 文件：

```env
# MySQL配置
DB_HOST=host.docker.internal
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your-mysql-password  # 修改为实际密码
DB_NAME=admin_system
DB_CHARSET=utf8mb4

# Redis配置
REDIS_HOST=host.docker.internal
REDIS_PORT=6379
REDIS_PASSWORD=  # 如果Redis设置了密码，填写 here
REDIS_DB=0

# JWT配置
JWT_SECRET=your-secret-key-change-in-production
JWT_EXPIRE_HOURS=72
JWT_REFRESH_THRESHOLD_HOURS=24

# 服务配置
SERVER_PORT=7701
SERVER_MODE=release
```

### 3. 验证安装

**检查MySQL:**
```bash
# 在服务器上执行
systemctl status mysqld  # CentOS/RHEL
# 或
systemctl status mysql   # Ubuntu/Debian

# 测试连接
mysql -u root -p
```

**检查Redis:**
```bash
# 在服务器上执行
systemctl status redis    # CentOS/RHEL
# 或
systemctl status redis-server  # Ubuntu/Debian

# 测试连接
redis-cli ping
# 应该返回: PONG
```

## 常见问题

### 问题1: Docker容器无法连接MySQL

**解决方案:**

1. **检查MySQL是否允许远程连接:**
```sql
mysql -u root -p
SELECT host, user FROM mysql.user;
-- 确保有 'root'@'%' 用户
```

2. **如果不存在，创建:**
```sql
GRANT ALL PRIVILEGES ON admin_system.* TO 'root'@'%' IDENTIFIED BY 'your-password';
FLUSH PRIVILEGES;
```

3. **检查MySQL绑定地址:**
```bash
# 编辑MySQL配置文件
vi /etc/my.cnf  # 或 /etc/mysql/my.cnf
# 确保有: bind-address = 0.0.0.0
```

4. **如果host.docker.internal不可用，使用host网络:**
```yaml
# 在docker-compose.yml中
network_mode: "host"
# 并注释掉ports和extra_hosts
```

### 问题2: Docker容器无法连接Redis

**解决方案:**

1. **检查Redis绑定地址:**
```bash
# 编辑Redis配置文件
vi /etc/redis/redis.conf
# 修改: bind 0.0.0.0
# 或注释掉: #bind 127.0.0.1
```

2. **重启Redis:**
```bash
systemctl restart redis  # 或 redis-server
```

3. **如果设置了密码，在.env中配置:**
```env
REDIS_PASSWORD=your-redis-password
```

### 问题3: 防火墙阻止连接

**解决方案:**

```bash
# CentOS/RHEL (firewalld)
firewall-cmd --permanent --add-port=3306/tcp
firewall-cmd --permanent --add-port=6379/tcp
firewall-cmd --reload

# Ubuntu/Debian (ufw)
ufw allow 3306/tcp
ufw allow 6379/tcp
```

## 安全建议

1. **设置强密码**
   - MySQL root密码应该足够复杂
   - Redis建议设置密码

2. **限制访问**
   - 只允许必要的IP访问
   - 使用防火墙规则限制

3. **定期更新**
   - 保持MySQL和Redis版本更新

## 下一步

安装完成后：

1. 配置 `.env` 文件
2. 运行 `./deploy.sh` 启动后端服务
3. 检查服务日志确认连接成功

