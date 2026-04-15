# OpenCodeUI Deploy

一个小工具，用于部署 [OpenCodeUI](https://github.com/lehhair/OpenCodeUI)。

## 为什么

OpenCodeUI 推荐使用 Docker 部署，但每次更新需要 `docker pull` 重新启动容器，感觉有点麻烦。

这个小工具可以直接下载前端静态文件，用原生 HTTP 服务替代 Docker，减少维护负担。

## 使用

### 下载工具

从 [Releases](https://github.com/ender049/opencodeui/releases) 下载对应平台的二进制。

### 启动服务

```bash
# 默认监听本机 127.0.0.1:3000，后端 127.0.0.1:4096
./opencodeui start

# 公开监听
./opencodeui start --ip 0.0.0.0 -p 8080

# 指定后端
./opencodeui start --backend remote:4096
```

### 管理

```bash
./opencodeui status   # 查看状态
./opencodeui stop     # 停止服务
./opencodeui version  # 查看版本
./opencodeui update   # 更新工具和前端
```

### 下载前端

前端会自动下载，也可手动下载 `opencodeui-dist-{version}.tar.gz` 解压到 `dist/` 目录。

## 开发

```bash
# 构建
go build -o opencodeui .

# 发布工具（创建 git tag）
git tag v0.1.0
git push
```

## 上游

- [lehhair/OpenCodeUI](https://github.com/lehhair/OpenCodeUI) - 本项目只同步上游 release，追踪其版本号
