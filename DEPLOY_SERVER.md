# 服务器部署指南

## 服务器信息

- **公网IP**: 39.105.136.116
- **内网IP**: 172.28.196.78
- **SSH用户**: root
- **SSH密码**: huojian123456!@#$%^
- **部署路径**: /opt/admin-system

## 快速部署

### 方式1：使用自动部署脚本（推荐）

```bash
# 1. 添加执行权限
chmod +x deploy-server.sh

# 2. 运行部署脚本
./deploy-server.sh
```

脚本会自动：
- 检查服务器连接
- 检查服务器环境（Docker、Docker Compose等）
- 打包项目文件
- 上传到服务器
- 解压到部署目录

### 方式2：手动部署

#### 步骤1：SSH登录服务器

```bash
ssh root@39.105.136.116
# 输入密码: huojian123456!@#$%^
```

#### 步骤2：安装必要软件

```bash
# 更新系统
apt-get update

# 安装Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh

# 安装Docker Compose
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# 启动Docker服务
systemctl start docker
systemctl enable docker
```

#### 步骤3：上传项目文件

**方式A：使用SCP**

```bash
# 在本地项目目录执行
scp -r . root@39.105.136.116:/opt/admin-system/
```

**方式B：使用Git（如果服务器有Git）**

```bash
# 在服务器上执行
cd /opt
git clone <your-repo-url> admin-system
cd admin-system
```

#### 步骤4：配置环境变量

```bash
# SSH登录服务器
ssh root@39.105.136.116

# 进入部署目录
cd /opt/admin-system

# 创建.env文件
cp .env.example .env

# 编辑.env文件（重要：修改所有密码和密钥）
vi .env
```

**必须修改的配置**：

```env
# MySQL - 使用强密码
MYSQL_ROOT_PASSWORD=your-strong-password-here
DB_PASSWORD=your-strong-password-here

# JWT - 使用强随机字符串（至少32字符）
JWT_SECRET=your-very-long-random-secret-key-at-least-32-characters

# 如果需要外网访问，修改端口映射
# 在docker-compose.yml中修改端口映射
```

#### 步骤5：启动服务

```bash
cd /opt/admin-system

# 添加执行权限
chmod +x *.sh

# 部署并启动
./deploy.sh
```

#### 步骤6：检查服务状态

```bash
# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f backend

# 测试API
curl http://localhost:7701/api/auth/userinfo
```

## 服务器配置

### 防火墙配置

如果需要外网访问，需要开放端口：

```bash
# Ubuntu/Debian (ufw)
ufw allow 7701/tcp  # 后端API
ufw allow 3306/tcp  # MySQL（仅内网，不建议开放）
ufw allow 6379/tcp  # Redis（仅内网，不建议开放）

# CentOS/RHEL (firewalld)
firewall-cmd --permanent --add-port=7701/tcp
firewall-cmd --reload
```

### 安全建议

1. **修改SSH端口**（可选）
2. **禁用root密码登录，使用密钥认证**（推荐）
3. **配置防火墙规则**
4. **定期更新系统**
5. **使用强密码**

## 域名配置（可选）

如果需要使用域名访问：

### 1. 配置Nginx反向代理

```bash
# 安装Nginx
apt-get install nginx

# 创建配置文件
vi /etc/nginx/sites-available/admin-system
```

Nginx配置示例：

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:7701;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

```bash
# 启用配置
ln -s /etc/nginx/sites-available/admin-system /etc/nginx/sites-enabled/
nginx -t
systemctl reload nginx
```

### 2. 配置SSL证书（Let's Encrypt）

```bash
# 安装Certbot
apt-get install certbot python3-certbot-nginx

# 获取证书
certbot --nginx -d your-domain.com
```

## 常用操作

### 查看服务状态

```bash
cd /opt/admin-system
docker-compose ps
```

### 查看日志

```bash
# 后端日志
docker-compose logs -f backend

# 所有服务日志
docker-compose logs -f

# MySQL日志
docker-compose logs -f mysql
```

### 重启服务

```bash
cd /opt/admin-system
./restart.sh
# 或
docker-compose restart
```

### 停止服务

```bash
cd /opt/admin-system
./stop.sh
# 或
docker-compose down
```

### 更新代码

```bash
# 1. 在本地更新代码并测试

# 2. 使用部署脚本重新部署
./deploy-server.sh

# 3. 或在服务器上手动更新
ssh root@39.105.136.116
cd /opt/admin-system
git pull  # 如果使用Git
./deploy.sh
```

### 备份数据库

```bash
# 在服务器上执行
docker-compose exec mysql mysqldump -u root -p admin_system > backup_$(date +%Y%m%d_%H%M%S).sql

# 或使用脚本
cat > /opt/admin-system/backup-db.sh << 'EOF'
#!/bin/bash
BACKUP_DIR="/opt/admin-system/backups"
mkdir -p $BACKUP_DIR
docker-compose exec -T mysql mysqldump -u root -p${MYSQL_ROOT_PASSWORD} admin_system | gzip > $BACKUP_DIR/backup_$(date +%Y%m%d_%H%M%S).sql.gz
find $BACKUP_DIR -name "backup_*.sql.gz" -mtime +7 -delete
EOF
chmod +x /opt/admin-system/backup-db.sh
```

## 故障排查

### SSH连接问题

如果遇到 "Connection reset" 错误：

```bash
# 1. 检查服务器SSH服务
ssh root@39.105.136.116 "systemctl status sshd"

# 2. 检查防火墙
ssh root@39.105.136.116 "ufw status"

# 3. 尝试使用详细模式
ssh -v root@39.105.136.116
```

### 服务无法启动

```bash
# 查看详细日志
cd /opt/admin-system
docker-compose logs backend

# 检查配置文件
cat .env

# 检查端口占用
netstat -tulpn | grep 7701
```

### 数据库连接失败

```bash
# 检查MySQL容器
docker-compose ps mysql

# 查看MySQL日志
docker-compose logs mysql

# 测试连接
docker-compose exec mysql mysql -u root -p
```

## 监控和维护

### 设置自动备份（Crontab）

```bash
# 编辑crontab
crontab -e

# 添加每天凌晨2点备份
0 2 * * * /opt/admin-system/backup-db.sh
```

### 监控服务状态

```bash
# 创建监控脚本
cat > /opt/admin-system/check-service.sh << 'EOF'
#!/bin/bash
if ! docker-compose ps | grep -q "Up"; then
    echo "服务异常，正在重启..."
    cd /opt/admin-system
    docker-compose restart
fi
EOF
chmod +x /opt/admin-system/check-service.sh

# 添加到crontab（每5分钟检查一次）
*/5 * * * * /opt/admin-system/check-service.sh
```

## 联系信息

如有问题，请检查：
1. 服务器日志
2. Docker容器状态
3. 网络连接
4. 配置文件

