# 部署脚本说明

项目提供了多个便捷的部署和管理脚本，方便快速操作。

## 脚本列表

### 1. deploy.sh - 一键完整部署 ⭐推荐

**功能**: 编译代码 + 构建镜像 + 停止旧服务 + 启动新服务

**使用场景**: 
- 首次部署
- 代码更新后重新部署
- 需要重新构建镜像时

**使用方法**:
```bash
./deploy.sh
```

**执行流程**:
1. 检查Docker环境
2. 编译后端Go代码（如果Go已安装）
3. 停止旧服务
4. 构建Docker镜像
5. 启动所有服务
6. 检查服务健康状态

---

### 2. start.sh - 快速启动

**功能**: 直接启动所有服务（不重新构建）

**使用场景**: 
- 服务已构建，只需启动
- 服务停止后重新启动

**使用方法**:
```bash
./start.sh
```

---

### 3. restart.sh - 快速重启

**功能**: 重启所有服务（不重新构建镜像）

**使用场景**: 
- 服务异常需要重启
- 配置更新后重启

**使用方法**:
```bash
./restart.sh
```

---

### 4. build.sh - 仅构建镜像

**功能**: 只构建Docker镜像，不启动服务

**使用场景**: 
- 只想构建镜像，稍后手动启动
- 测试构建过程

**使用方法**:
```bash
./build.sh
```

---

### 5. stop.sh - 停止服务

**功能**: 停止所有服务

**使用场景**: 
- 需要停止所有服务
- 维护前停止服务

**使用方法**:
```bash
./stop.sh
```

---

## 脚本执行权限

如果脚本没有执行权限，请先添加：

```bash
chmod +x deploy.sh start.sh restart.sh build.sh stop.sh
```

或者一次性添加所有脚本权限：

```bash
chmod +x *.sh
```

## 常用操作流程

### 首次部署
```bash
./deploy.sh
```

### 代码更新后重新部署
```bash
./deploy.sh
```

### 服务异常重启
```bash
./restart.sh
```

### 查看服务日志
```bash
# 查看后端日志
docker-compose logs -f backend

# 查看所有服务日志
docker-compose logs -f

# 查看MySQL日志
docker-compose logs -f mysql

# 查看Redis日志
docker-compose logs -f redis
```

### 查看服务状态
```bash
docker-compose ps
```

### 进入容器调试
```bash
# 进入后端容器
docker-compose exec backend sh

# 进入MySQL容器
docker-compose exec mysql bash

# 进入Redis容器
docker-compose exec redis sh
```

## 故障排查

### 1. 脚本执行失败

**问题**: `Permission denied`

**解决**: 
```bash
chmod +x deploy.sh
```

### 2. Docker未运行

**问题**: `Docker未运行`

**解决**: 启动Docker Desktop或Docker服务

### 3. 端口被占用

**问题**: 端口7701、3306、6379被占用

**解决**: 
- 修改 `docker-compose.yml` 中的端口映射
- 或停止占用端口的其他服务

### 4. 构建失败

**问题**: Docker镜像构建失败

**解决**: 
- 检查网络连接
- 查看详细错误: `docker-compose build --no-cache`
- 检查Dockerfile语法

### 5. 服务启动失败

**问题**: 服务启动后立即退出

**解决**: 
```bash
# 查看详细日志
docker-compose logs backend

# 检查服务状态
docker-compose ps -a
```

## 注意事项

1. **首次部署**: 建议使用 `deploy.sh`，确保所有步骤正确执行
2. **代码更新**: 修改代码后使用 `deploy.sh` 重新部署
3. **快速重启**: 如果只是重启服务，使用 `restart.sh` 更快
4. **数据持久化**: 停止服务不会删除数据卷，数据会保留
5. **完全清理**: 如需完全清理（包括数据），使用 `docker-compose down -v`

## 高级用法

### 仅重新构建后端服务
```bash
docker-compose build backend
docker-compose up -d backend
```

### 查看资源使用情况
```bash
docker stats
```

### 清理未使用的镜像和容器
```bash
docker system prune -a
```

