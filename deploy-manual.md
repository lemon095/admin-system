# 手动部署指南（SSH连接问题时的替代方案）

如果自动部署脚本无法连接服务器，可以使用以下手动方法：

## 方案1: 手动SSH连接并部署

### 步骤1: 测试SSH连接

```bash
# 运行测试脚本
./test-ssh.sh

# 或直接尝试连接
ssh root@39.105.136.116
# 输入密码: huojian123456!@#$%^
```

### 步骤2: 如果SSH连接成功，手动上传文件

**在本地执行：**

```bash
# 打包项目（排除不需要的文件）
tar -czf admin-system.tar.gz \
  --exclude='.git' \
  --exclude='node_modules' \
  --exclude='.env' \
  --exclude='backend/.env' \
  --exclude='*.log' \
  --exclude='.DS_Store' \
  --exclude='frontend/dist' \
  --exclude='backend/main' \
  .

# 上传到服务器
scp admin-system.tar.gz root@39.105.136.116:/opt/
```

**在服务器上执行：**

```bash
# SSH登录
ssh root@39.105.136.116

# 解压文件
cd /opt
tar -xzf admin-system.tar.gz
mv admin-system/* admin-system/
rm -rf admin-system admin-system.tar.gz

# 进入目录
cd /opt/admin-system

# 创建.env文件
cp .env.example .env
vi .env  # 编辑配置文件，修改密码和密钥

# 部署
chmod +x *.sh
./deploy.sh
```

## 方案2: 使用Git部署（如果服务器有Git）

**在服务器上执行：**

```bash
# SSH登录
ssh root@39.105.136.116

# 安装Git（如果未安装）
apt-get update
apt-get install -y git

# 克隆项目（如果使用Git仓库）
cd /opt
git clone <your-repo-url> admin-system
cd admin-system

# 创建.env文件
cp .env.example .env
vi .env

# 部署
chmod +x *.sh
./deploy.sh
```

## 方案3: 使用FTP/SFTP客户端

1. 使用FileZilla、WinSCP等工具连接服务器
2. 上传项目文件到 `/opt/admin-system`
3. 在服务器上执行部署命令

## 方案4: 服务器端直接下载

**在服务器上执行：**

```bash
# SSH登录
ssh root@39.105.136.116

# 安装必要工具
apt-get update
apt-get install -y wget unzip

# 如果项目在Git仓库，直接克隆
cd /opt
git clone <your-repo-url> admin-system
cd admin-system

# 或从其他位置下载压缩包
# wget <your-file-url>
# unzip admin-system.zip

# 配置和部署
cp .env.example .env
vi .env
chmod +x *.sh
./deploy.sh
```

## 常见SSH连接问题解决

### 问题1: Connection reset by peer

**可能原因：**
- 服务器SSH服务未启动
- 防火墙阻止连接
- SSH配置问题

**解决方法：**

```bash
# 在服务器上检查SSH服务
systemctl status sshd
systemctl start sshd
systemctl enable sshd

# 检查防火墙
ufw status
# 如果防火墙开启，允许SSH
ufw allow 22/tcp
ufw reload

# 检查SSH配置
cat /etc/ssh/sshd_config | grep -E "Port|PermitRootLogin"
```

### 问题2: 端口不是22

如果SSH使用非标准端口：

```bash
# 使用指定端口连接
ssh -p <端口号> root@39.105.136.116
```

### 问题3: 需要VPN或跳板机

如果服务器在内网：

```bash
# 先连接跳板机
ssh user@jump-server

# 再从跳板机连接目标服务器
ssh root@172.28.196.78  # 使用内网IP
```

## 服务器环境准备

如果服务器还没有安装Docker，需要先安装：

```bash
# SSH登录服务器后执行

# 更新系统
apt-get update

# 安装Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh

# 启动Docker服务
systemctl start docker
systemctl enable docker

# 安装Docker Compose
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# 验证安装
docker --version
docker-compose --version
```

## 部署后验证

```bash
# 检查服务状态
cd /opt/admin-system
docker-compose ps

# 查看日志
docker-compose logs -f backend

# 测试API
curl http://localhost:7701/api/auth/userinfo
```

## 需要帮助？

如果以上方法都无法解决，请检查：

1. **服务器状态**：确认服务器是否正常运行
2. **网络连接**：确认能否ping通服务器
3. **防火墙规则**：确认SSH端口是否开放
4. **SSH服务**：确认SSH服务是否运行
5. **密码正确性**：确认密码是否正确

