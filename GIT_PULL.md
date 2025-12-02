# Git拉取代码部署指南

现在支持远程服务器直接从Git仓库拉取代码，无需本地传输文件。

## 快速开始

### 方式1: 使用本地脚本（推荐）

```bash
# 首次设置（只需一次）
./setup-git-pull.sh

# 拉取代码并部署
./pull-code.sh
```

### 方式2: 在服务器上直接执行

```bash
# SSH登录服务器
ssh -i ~/Desktop/chuchang/chuchang.pem root@39.105.136.116

# 执行拉取脚本
cd /opt/admin-system
./pull-and-deploy.sh
```

## 设置步骤

### 步骤1: 初始化Git仓库（首次）

如果服务器上还没有Git仓库：

```bash
# SSH登录服务器
ssh -i ~/Desktop/chuchang/chuchang.pem root@39.105.136.116

# 进入部署目录
cd /opt/admin-system

# 如果目录为空，克隆仓库
git clone <your-git-repo-url> .

# 或如果已有代码，初始化Git
git init
git remote add origin <your-git-repo-url>
git add .
git commit -m "Initial commit"
git push -u origin main
```

### 步骤2: 运行设置脚本

```bash
./setup-git-pull.sh
```

这个脚本会：
- 检查并安装Git（如果未安装）
- 在服务器上创建 `pull-and-deploy.sh` 脚本
- 配置自动拉取和部署功能

### 步骤3: 配置Git仓库地址（如果需要）

如果需要在脚本中指定Git仓库：

```bash
# 方式1: 设置环境变量
export GIT_REPO=https://github.com/your-username/your-repo.git
./pull-code.sh

# 方式2: 修改pull-code.sh中的GIT_REPO变量
```

## 使用方法

### 拉取并部署全部

```bash
./pull-code.sh
# 或
./pull-code.sh all
```

### 只更新后端

```bash
./pull-code.sh backend
```

### 只更新前端

```bash
./pull-code.sh frontend
```

### 指定分支

```bash
# 在服务器上执行
cd /opt/admin-system
./pull-and-deploy.sh develop
```

## 工作流程

1. **拉取代码**: 从Git仓库拉取最新代码
2. **备份**: 自动备份当前代码到 `/tmp/`
3. **更新依赖**: 
   - 后端: 运行 `go mod download`
   - 前端: 运行 `npm install`
4. **重新构建**: 重新构建Docker镜像
5. **重启服务**: 自动重启后端服务
6. **检查状态**: 显示服务状态和日志

## 优势

✅ **无需本地传输**: 直接从Git拉取，不占用本地带宽  
✅ **版本控制**: 可以指定分支、标签  
✅ **自动备份**: 每次更新前自动备份  
✅ **快速部署**: 一条命令完成所有操作  
✅ **支持回滚**: 可以切换到任意提交  

## 常见问题

### 问题1: Git仓库未配置

**解决**:
```bash
# 在服务器上执行
cd /opt/admin-system
git remote -v  # 查看当前配置
git remote set-url origin <your-repo-url>  # 设置仓库地址
```

### 问题2: 需要认证

**解决**:
```bash
# 使用SSH密钥（推荐）
ssh-keygen -t rsa -b 4096
# 将公钥添加到Git仓库的SSH keys

# 或使用HTTPS + 个人访问令牌
git remote set-url origin https://token@github.com/user/repo.git
```

### 问题3: 拉取失败

**解决**:
```bash
# 检查网络连接
ping github.com  # 或你的Git服务器

# 检查Git配置
git config --list

# 手动拉取测试
git pull origin main
```

### 问题4: 冲突处理

**解决**:
```bash
# 在服务器上执行
cd /opt/admin-system
git status  # 查看冲突
git stash   # 暂存本地修改
git pull origin main
# 或强制拉取（会覆盖本地修改）
git fetch origin
git reset --hard origin/main
```

## 与现有脚本对比

| 脚本 | 传输方式 | 需要Git | 速度 | 适用场景 |
|------|----------|---------|------|----------|
| `quick-update.sh` | rsync/tar+scp | ❌ | 快 | 本地开发 |
| `update-code.sh` | tar+scp | ❌ | 中 | 手动部署 |
| `update-and-deploy.sh` | tar+scp | ❌ | 慢 | 生产环境 |
| `pull-code.sh` | Git拉取 | ✅ | 快 | **推荐** |

## 最佳实践

1. **使用Git分支管理**
   - `main`: 生产环境
   - `develop`: 开发环境
   - `feature/*`: 功能分支

2. **使用标签发布**
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   # 在服务器上
   git checkout v1.0.0
   ```

3. **定期备份**
   - 脚本自动备份到 `/tmp/`
   - 建议定期清理旧备份

4. **监控部署**
   ```bash
   # 查看部署日志
   docker-compose logs -f backend
   ```

## 自动化部署（可选）

可以设置定时任务自动拉取：

```bash
# 在服务器上设置crontab
crontab -e

# 每天凌晨2点自动拉取并部署
0 2 * * * cd /opt/admin-system && ./pull-and-deploy.sh main >> /var/log/deploy.log 2>&1
```

