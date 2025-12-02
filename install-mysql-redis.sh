#!/bin/bash

# 在服务器上安装MySQL和Redis的脚本
# 使用方法: 在服务器上执行此脚本，或通过SSH远程执行

set -e

echo "=========================================="
echo "安装MySQL和Redis"
echo "=========================================="

# 检测操作系统
if [ -f /etc/os-release ]; then
    . /etc/os-release
    OS=$ID
    VER=$VERSION_ID
else
    echo "无法检测操作系统"
    exit 1
fi

echo "检测到操作系统: $OS $VER"
echo ""

# 安装MySQL
echo "=========================================="
echo "安装MySQL"
echo "=========================================="

if command -v mysql &> /dev/null; then
    MYSQL_VERSION=$(mysql --version)
    echo "✅ MySQL已安装: ${MYSQL_VERSION}"
else
    echo "正在安装MySQL..."
    
    if [ "$OS" == "centos" ] || [ "$OS" == "rhel" ] || [ "$OS" == "almalinux" ] || [ "$OS" == "rocky" ]; then
        # CentOS/RHEL/AlmaLinux/Rocky
        yum install -y mysql-server
        systemctl start mysqld
        systemctl enable mysqld
        
        # 获取临时密码
        TEMP_PASSWORD=$(grep 'temporary password' /var/log/mysqld.log 2>/dev/null | awk '{print $NF}' | tail -1)
        if [ -n "$TEMP_PASSWORD" ]; then
            echo "⚠️  MySQL临时密码: $TEMP_PASSWORD"
            echo "   请运行: mysql_secure_installation"
        fi
    elif [ "$OS" == "ubuntu" ] || [ "$OS" == "debian" ]; then
        # Ubuntu/Debian
        export DEBIAN_FRONTEND=noninteractive
        apt-get update
        apt-get install -y mysql-server
        systemctl start mysql
        systemctl enable mysql
    else
        echo "❌ 不支持的操作系统: $OS"
        exit 1
    fi
    
    echo "✅ MySQL安装完成"
fi

# 检查MySQL服务状态
if systemctl is-active --quiet mysqld 2>/dev/null || systemctl is-active --quiet mysql 2>/dev/null; then
    echo "✅ MySQL服务运行中"
else
    echo "启动MySQL服务..."
    systemctl start mysqld 2>/dev/null || systemctl start mysql
    systemctl enable mysqld 2>/dev/null || systemctl enable mysql
    echo "✅ MySQL服务已启动"
fi

# 安装Redis
echo ""
echo "=========================================="
echo "安装Redis"
echo "=========================================="

if command -v redis-cli &> /dev/null; then
    REDIS_VERSION=$(redis-cli --version)
    echo "✅ Redis已安装: ${REDIS_VERSION}"
else
    echo "正在安装Redis..."
    
    if [ "$OS" == "centos" ] || [ "$OS" == "rhel" ] || [ "$OS" == "almalinux" ] || [ "$OS" == "rocky" ]; then
        # CentOS/RHEL/AlmaLinux/Rocky
        yum install -y redis
        systemctl start redis
        systemctl enable redis
    elif [ "$OS" == "ubuntu" ] || [ "$OS" == "debian" ]; then
        # Ubuntu/Debian
        apt-get update
        apt-get install -y redis-server
        systemctl start redis-server
        systemctl enable redis-server
    fi
    
    echo "✅ Redis安装完成"
fi

# 检查Redis服务状态
if systemctl is-active --quiet redis 2>/dev/null || systemctl is-active --quiet redis-server 2>/dev/null; then
    echo "✅ Redis服务运行中"
else
    echo "启动Redis服务..."
    systemctl start redis 2>/dev/null || systemctl start redis-server
    systemctl enable redis 2>/dev/null || systemctl enable redis-server
    echo "✅ Redis服务已启动"
fi

# 配置MySQL允许远程连接（Docker容器需要）
echo ""
echo "=========================================="
echo "配置MySQL允许Docker容器访问"
echo "=========================================="

read -p "MySQL root密码（如果已设置，直接回车跳过）: " -s MYSQL_PASSWORD
echo

if [ -z "$MYSQL_PASSWORD" ]; then
    echo "⚠️  未输入密码，跳过MySQL配置"
    echo "   如果需要配置，请手动执行以下SQL："
    echo "   CREATE DATABASE IF NOT EXISTS admin_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
    echo "   GRANT ALL PRIVILEGES ON admin_system.* TO 'root'@'%' IDENTIFIED BY 'your-password';"
    echo "   FLUSH PRIVILEGES;"
else
    # 创建数据库
    mysql -u root -p"${MYSQL_PASSWORD}" << EOF 2>/dev/null || echo "⚠️  数据库可能已存在或密码错误"
CREATE DATABASE IF NOT EXISTS admin_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EOF
    
    # 允许远程连接（从Docker容器）
    mysql -u root -p"${MYSQL_PASSWORD}" << EOF 2>/dev/null || echo "⚠️  配置远程访问可能失败"
GRANT ALL PRIVILEGES ON admin_system.* TO 'root'@'%' IDENTIFIED BY '${MYSQL_PASSWORD}';
FLUSH PRIVILEGES;
EOF
    
    echo "✅ MySQL配置完成"
fi

# 配置MySQL绑定地址
MYSQL_CONF=""
if [ -f /etc/my.cnf ]; then
    MYSQL_CONF="/etc/my.cnf"
elif [ -f /etc/mysql/my.cnf ]; then
    MYSQL_CONF="/etc/mysql/my.cnf"
elif [ -f /etc/mysql/mysql.conf.d/mysqld.cnf ]; then
    MYSQL_CONF="/etc/mysql/mysql.conf.d/mysqld.cnf"
fi

if [ -n "$MYSQL_CONF" ]; then
    if ! grep -q "^bind-address" "$MYSQL_CONF"; then
        echo "配置MySQL绑定地址..."
        sed -i '/\[mysqld\]/a bind-address = 0.0.0.0' "$MYSQL_CONF"
        systemctl restart mysqld 2>/dev/null || systemctl restart mysql
        echo "✅ MySQL已配置为允许远程连接"
    else
        echo "✅ MySQL绑定地址已配置"
    fi
fi

# 配置Redis允许远程连接
echo ""
echo "=========================================="
echo "配置Redis允许Docker容器访问"
echo "=========================================="

REDIS_CONF=""
if [ -f /etc/redis/redis.conf ]; then
    REDIS_CONF="/etc/redis/redis.conf"
elif [ -f /etc/redis.conf ]; then
    REDIS_CONF="/etc/redis.conf"
fi

if [ -n "$REDIS_CONF" ]; then
    # 注释掉bind 127.0.0.1，允许所有IP访问
    if grep -q "^bind 127.0.0.1" "$REDIS_CONF"; then
        echo "配置Redis绑定地址..."
        sed -i 's/^bind 127.0.0.1/bind 0.0.0.0/' "$REDIS_CONF"
        # 或者注释掉
        # sed -i 's/^bind 127.0.0.1/#bind 127.0.0.1/' "$REDIS_CONF"
        
        # 如果配置了保护模式，需要设置密码或允许无密码访问
        if grep -q "^protected-mode yes" "$REDIS_CONF"; then
            # 选项1: 关闭保护模式（不推荐生产环境）
            # sed -i 's/^protected-mode yes/protected-mode no/' "$REDIS_CONF"
            # 选项2: 设置密码（推荐）
            if ! grep -q "^requirepass" "$REDIS_CONF"; then
                echo "⚠️  Redis未设置密码，建议设置密码以提高安全性"
            fi
        fi
        
        systemctl restart redis 2>/dev/null || systemctl restart redis-server
        echo "✅ Redis已配置为允许远程连接"
    else
        echo "✅ Redis绑定地址已配置"
    fi
else
    echo "⚠️  未找到Redis配置文件"
fi

# 配置防火墙（如果需要）
echo ""
echo "=========================================="
echo "配置防火墙"
echo "=========================================="

if command -v firewall-cmd &> /dev/null; then
    # firewalld (CentOS/RHEL)
    echo "配置firewalld..."
    firewall-cmd --permanent --add-port=3306/tcp 2>/dev/null && echo "✅ 已开放MySQL端口3306"
    firewall-cmd --permanent --add-port=6379/tcp 2>/dev/null && echo "✅ 已开放Redis端口6379"
    firewall-cmd --reload 2>/dev/null
elif command -v ufw &> /dev/null; then
    # ufw (Ubuntu/Debian)
    echo "配置ufw..."
    ufw allow 3306/tcp 2>/dev/null && echo "✅ 已开放MySQL端口3306"
    ufw allow 6379/tcp 2>/dev/null && echo "✅ 已开放Redis端口6379"
else
    echo "⚠️  未检测到防火墙管理工具"
fi

# 测试连接
echo ""
echo "=========================================="
echo "测试连接"
echo "=========================================="

# 测试MySQL
if command -v mysql &> /dev/null; then
    if mysql -u root -e "SELECT 1;" 2>/dev/null || mysql -u root -p"${MYSQL_PASSWORD}" -e "SELECT 1;" 2>/dev/null; then
        echo "✅ MySQL连接测试成功"
    else
        echo "⚠️  MySQL连接测试失败，请检查密码"
    fi
fi

# 测试Redis
if command -v redis-cli &> /dev/null; then
    if redis-cli ping 2>/dev/null | grep -q PONG; then
        echo "✅ Redis连接测试成功"
    else
        echo "⚠️  Redis连接测试失败"
    fi
fi

echo ""
echo "=========================================="
echo "✅ 安装完成！"
echo "=========================================="
echo ""
echo "📝 服务状态："
systemctl status mysqld --no-pager -l 2>/dev/null | head -3 || systemctl status mysql --no-pager -l 2>/dev/null | head -3
echo ""
systemctl status redis --no-pager -l 2>/dev/null | head -3 || systemctl status redis-server --no-pager -l 2>/dev/null | head -3
echo ""
echo "📝 下一步："
echo "   1. 确保.env文件中配置了正确的MySQL密码"
echo "   2. 运行 ./deploy.sh 启动后端服务"
echo "=========================================="

