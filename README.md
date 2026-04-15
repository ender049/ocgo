# OpenCodeUI Deploy

部署工具，包含：

- **Go服务端** - 静态文件服务 + API代理
- **前端构建包** - 自动从上游构建

## 使用

### 下载工具

从 [Releases](https://github.com/lehhair/opencodeui-deploy/releases) 下载对应平台的二进制。

### 下载前端

从 Releases 下载 `opencodeui-dist-{version}.tar.gz`，解压到 `dist/` 目录。

```bash
tar -xzf opencodeui-dist-v0.5.1.tar.gz
```

### 启动

```bash
./opencodeui serve --backend localhost:4096
```

## 开发

```bash
# 构建工具
go build -o opencodeui .

# 发布
git tag v0.1.0
git push origin v0.1.0
```

## 上游

- OpenCodeUI: https://github.com/lehhair/OpenCodeUI
