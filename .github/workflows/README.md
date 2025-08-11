# GitHub Actions 自动构建部署

## 功能说明

这个工作流会在您推送代码到 `main` 或 `master` 分支时，自动：

1. 构建 Docker 镜像
2. 推送到阿里云容器镜像服务
3. 使用 `latest` 标签

## 配置步骤

### 1. 配置 GitHub Secrets

在您的 GitHub 仓库中，进入 `Settings` → `Secrets and variables` → `Actions`，添加以下 secrets：

#### **ALIYUN_USERNAME**
- 阿里云容器镜像服务的用户名
- 通常是您的阿里云账号

#### **ALIYUN_PASSWORD**
- 阿里云容器镜像服务的密码
- 在阿里云控制台 → 容器镜像服务 → 访问凭证中获取

### 2. 修改配置（可选）

如果需要修改镜像仓库地址，编辑 `.github/workflows/docker-build.yml` 文件中的环境变量：

```yaml
env:
  REGISTRY: registry.cn-hangzhou.aliyuncs.com  # 阿里云镜像仓库地址
  NAMESPACE: luoqiangtest                      # 您的命名空间
  IMAGE_NAME: go-admin                         # 镜像名称
```

### 3. 创建镜像仓库

在阿里云容器镜像服务控制台中：
1. 创建命名空间（如果不存在）
2. 创建镜像仓库 `go-admin`

## 使用方法

### 自动触发
- 推送代码到 `main` 或 `master` 分支时自动触发

### 手动触发
- 在 GitHub 仓库页面，进入 `Actions` 标签
- 选择 `Build and Push to Aliyun Registry` 工作流
- 点击 `Run workflow` 手动触发

## 镜像地址

构建完成后，镜像地址为：
```
registry.cn-hangzhou.aliyuncs.com/luoqiangtest/go-admin:latest
```

## 拉取镜像

```bash
docker pull registry.cn-hangzhou.aliyuncs.com/luoqiangtest/go-admin:latest
```

## 运行容器

```bash
docker run -d -p 8000:8000 --name go-admin registry.cn-hangzhou.aliyuncs.com/luoqiangtest/go-admin:latest
``` 