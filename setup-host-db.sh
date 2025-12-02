#!/bin/bash

# 在服务器上设置MySQL和Redis的脚本
# 如果MySQL和Redis未安装，此脚本可以帮助安装

set -e

echo "=========================================="
echo "设置宿主机MySQL和Redis"
echo "=========================================="

# 检查MySQL
echo ""
echo "🔍 检查MySQL..."
if command -v mysql &> /dev/null; then
    MYSQL_VERSION=$(mysql --version)
    echo "✅ MySQL已安装: ${MYSQL_VERSION}"
else
    echo "⚠️  MySQL未安装"
    read -p "是否安装MySQL? (y/n) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "正在安装MySQL..."
        # CentOS/RHEL
        if command -v yum &> /dev/null; then
            yum install -y mysql-server
            systemctl start mysqld
            systemctl enable mysqld
        # Ubuntu/Debian
        elif command -v apt-get &> /dev/null; then
            apt-get update
            apt-get install -y mysql-server
            systemctl start mysql
            systemctl enable mysql
        fi
        echo "✅ MySQL安装完成"
        echo "⚠️  请设置MySQL root密码: mysql_secure_installation"
    fi
fi

# 检查MySQL服务状态
if systemctl is-active --quiet mysqld || systemctl is-active --quiet mysql; then
    echo "✅ MySQL服务运行中"
else
    echo "⚠️  MySQL服务未运行，正在启动..."
    systemctl start mysqld 2>/dev/null || systemctl start mysql
    systemctl enable mysqld 2>/dev/null || systemctl enable mysql
fi

# 检查Redis
echo ""
echo "🔍 检查Redis..."
if command -v redis-cli &> /dev/null; then
    REDIS_VERSION=$(redis-cli --version)
    echo "✅ Redis已安装: ${REDIS_VERSION}"
else
    echo "⚠️  Redis未安装"
    read -p "是否安装Redis? (y/n) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "正在安装Redis..."
        # CentOS/RHEL
        if command -v yum &> /dev/null; then
            yum install -y redis
            systemctl start redis
            systemctl enable redis
        # Ubuntu/Debian
        elif command -v apt-get &> /dev/null; then
            apt-get update
            apt-get install -y redis-server
            systemctl start redis-server
            systemctl enable redis-server
        fi
        echo "✅ Redis安装完成"
    fi
fi

# 检查Redis服务状态
if systemctl is-active --quiet redis || systemctl is-active --quiet redis-server; then
    echo "✅ Redis服务运行中"
else
    echo "⚠️  Redis服务未运行，正在启动..."
    systemctl start redis 2>/dev/null || systemctl start redis-server
    systemctl enable redis 2>/dev/null || systemctl enable redis-server
fi

# 创建数据库
echo ""
echo "🔍 检查数据库..."
read -p "MySQL root密码: " -s MYSQL_PASSWORD
echo

if mysql -u root -p"${MYSQL_PASSWORD}" -e "USE admin_system;" 2>/dev/null; then
    echo "✅ 数据库 admin_system 已存在"
else
    echo "⚠️  数据库不存在，正在创建..."
    mysql -u root -p"${MYSQL_PASSWORD}" << EOF
CREATE DATABASE IF NOT EXISTS admin_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EOF
    echo "✅ 数据库创建成功"
fi

# 检查MySQL是否允许远程连接（Docker容器需要）
echo ""
echo "🔍 检查MySQL远程连接配置..."
echo "⚠️  确保MySQL允许从Docker容器连接"
echo "   如果连接失败，可能需要执行："
echo "   GRANT ALL PRIVILEGES ON admin_system.* TO 'root'@'%' IDENTIFIED BY 'your-password';"
echo "   FLUSH PRIVILEGES;"

# 检查Redis绑定配置
echo ""
echo "🔍 检查Redis配置..."
if grep -q "^bind 127.0.0.1" /etc/redis/redis.conf 2>/dev/null || grep -q "^bind 127.0.0.1" /etc/redis.conf 2>/dev/null; then
    echo "⚠️  Redis可能只绑定到127.0.0.1，Docker容器可能无法连接"
    echo "   建议修改Redis配置，允许从Docker网络连接"
    echo "   或使用: bind 0.0.0.0"
fi

echo ""
echo "=========================================="
echo "✅ 设置完成！"
echo "=========================================="
echo ""
echo "📝 下一步："
echo "   1. 确保MySQL和Redis服务正在运行"
echo "   2. 配置.env文件，设置正确的数据库密码"
echo "   3. 运行 ./deploy.sh 启动后端服务"
echo "=========================================="

