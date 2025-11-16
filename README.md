# feishu2md

🚀 **强大的飞书文档转 Markdown 工具** - 支持单文档、批量下载和知识库导出，智能处理图片并自动上传到图床。

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

---

## ✨ 核心特性

| 特性 | 说明 |
|------|------|
| 📄 **多种下载模式** | 单文档、文件夹批量、整个知识库、子文档递归下载 |
| 🖼️ **智能图片处理** | 自动下载图片，支持本地保存或上传图床 |
| ☁️ **图床自动上传** | 支持阿里云 OSS、腾讯云 COS，自动替换图片链接 |
| 🌳 **保持文档结构** | 递归下载时保持原有层级结构 |
| ⚡ **高效并发** | 支持多线程并发下载，智能限流 |
| 📝 **友好文件名** | 默认使用文档标题，智能处理特殊字符 |
| 🎯 **格式完整** | 完整支持表格、列表、代码块等 Markdown 格式 |
| 💾 **智能缓存** | 图片和文档去重，避免重复下载和上传 |
| 🔧 **配置管理** | 环境变量配置，一键初始化配置文件 |

---

## 🚀 快速开始

### 1. 安装

```bash
# 克隆仓库
git clone https://github.com/Perfecto23/feishu2md.git
cd feishu2md

# 编译
make build

# 或使用 go build
go build -o feishu2md ./cmd/...
```

### 2. 初始化配置

```bash
# 创建配置文件
./feishu2md init

# 编辑配置文件
vim .env
```

配置文件示例：

```bash
# 飞书 API 认证（必需）
FEISHU_APP_ID=your_app_id
FEISHU_APP_SECRET=your_app_secret

# 知识库配置（wiki-tree 命令需要）
FEISHU_SPACE_ID=your_space_id
FEISHU_FOLDER_TOKEN=https://xxx.feishu.cn/wiki/your_node_token

# 图床配置（可选）
IMGBED_ENABLED=true
IMGBED_PLATFORM=oss  # oss 或 cos
IMGBED_SECRET_ID=your_secret_id
IMGBED_SECRET_KEY=your_secret_key
IMGBED_BUCKET=your-bucket
IMGBED_REGION=oss-cn-hangzhou
IMGBED_PREFIX_KEY=images/
```

### 3. 开始使用

```bash
# 下载单个文档
./feishu2md document https://xxx.feishu.cn/docx/abc123

# 批量下载文件夹
./feishu2md folder https://xxx.feishu.cn/drive/folder/abc123

# 下载整个知识库
./feishu2md wiki https://xxx.feishu.cn/wiki/space/abc123

# 下载知识库子文档（使用配置文件中的设置）
./feishu2md wiki-tree
```

---

## 📖 详细用法

### 命令概览

| 命令 | 别名 | 说明 |
|------|------|------|
| `init` | `i` | 创建配置文件模板 |
| `document` | `doc`, `d` | 下载单个文档 |
| `folder` | `f`, `batch` | 批量下载文件夹 |
| `wiki` | `w` | 下载整个知识库 |
| `wiki-tree` | `wt`, `children` | 下载子文档树 |

### 全局选项

| 参数 | 说明 | 默认值 |
|------|------|--------|
| `--config`, `-c` | 配置文件路径 | `.env` |
| `--out`, `-o` | 输出目录 | `./dist` |
| `--img-dir` | 图片目录 | `img` |
| `--title-name`, `-t` | 使用标题作为文件名 | `true` |
| `--skip-same`, `-s` | 跳过重复文件（MD5检查） | `true` |
| `--force`, `-f` | 强制下载 | `false` |
| `--no-img` | 跳过图片下载 | `false` |
| `--html` | 使用 HTML 而非 Markdown | `false` |
| `--json` | 导出 JSON 响应 | `false` |

---

## 🖼️ 图床功能

### 支持的图床平台

- ✅ **阿里云 OSS** (`oss`)
- ✅ **腾讯云 COS** (`cos`)

### 配置图床

#### 阿里云 OSS

```bash
IMGBED_ENABLED=true
IMGBED_PLATFORM=oss
IMGBED_SECRET_ID=你的AccessKeyID
IMGBED_SECRET_KEY=你的AccessKeySecret
IMGBED_BUCKET=your-bucket-name
IMGBED_REGION=oss-cn-hangzhou  # 或其他区域
IMGBED_PREFIX_KEY=blog/images/  # 可选，上传路径前缀
```

**区域代码**：
- `oss-cn-hangzhou` - 华东1（杭州）
- `oss-cn-beijing` - 华北2（北京）
- `oss-cn-shanghai` - 华东2（上海）
- `oss-cn-shenzhen` - 华南1（深圳）

#### 腾讯云 COS

```bash
IMGBED_ENABLED=true
IMGBED_PLATFORM=cos
IMGBED_SECRET_ID=你的SecretId
IMGBED_SECRET_KEY=你的SecretKey
IMGBED_BUCKET=your-bucket-appid
IMGBED_REGION=ap-guangzhou  # 或其他区域
IMGBED_PREFIX_KEY=blog/images/
```

**区域代码**：
- `ap-guangzhou` - 广州
- `ap-beijing` - 北京
- `ap-shanghai` - 上海
- `ap-chengdu` - 成都

### 图床功能特性

- ✅ **智能去重** - 相同图片只上传一次
- ✅ **批量上传** - 并发上传提高效率
- ✅ **自动跳过** - 已上传的图片不会重复上传
- ✅ **链接替换** - 自动将 Markdown 中的图片链接替换为图床 URL
- ✅ **本地缓存** - 保留本地图片副本作为备份

### 使用示例

```bash
# 启用图床下载文档
./feishu2md document https://xxx.feishu.cn/docx/abc123

# 输出示例
📤 图床上传已启用: 阿里云OSS
   ├─ 图床: 上传成功 6 张
   ├─ 图片: 命中缓存 0, 新下载 6
✅ 文档标题

# 第二次运行（图片已缓存）
📤 图床上传已启用: 阿里云OSS
   ├─ 图床: 所有图片均已上传（跳过）
   ├─ 图片: 命中缓存 6, 新下载 0
⏭️  跳过重复文件: 文档标题
```

---

## 📚 使用场景

### 场景 1: 下载单个文档

```bash
# 基础用法
./feishu2md document https://xxx.feishu.cn/docx/abc123

# 指定输出目录
./feishu2md document https://xxx.feishu.cn/docx/abc123 --out ./docs

# 启用图床上传
# 在 .env 中配置 IMGBED_ENABLED=true
./feishu2md document https://xxx.feishu.cn/docx/abc123
```

**输出结构**：
```
dist/
├── 文档标题.md
└── img/
    ├── image1.png
    └── image2.jpg
```

### 场景 2: 批量下载文件夹

```bash
./feishu2md folder https://xxx.feishu.cn/drive/folder/abc123
```

**输出结构**：
```
dist/
├── 子文件夹1/
│   ├── 文档1.md
│   └── img/
├── 子文件夹2/
│   ├── 文档2.md
│   └── img/
└── 文档3.md
```

### 场景 3: 下载知识库

```bash
# 下载整个知识库
./feishu2md wiki https://xxx.feishu.cn/wiki/space/abc123
```

### 场景 4: 下载知识库子文档树

这是最强大的功能，可以下载知识库中某个节点下的所有子文档。

**配置 .env**：
```bash
FEISHU_SPACE_ID=7474915720537620484
FEISHU_FOLDER_TOKEN=https://xxx.feishu.cn/wiki/MekRwTsI9izbqbk
```

**运行**：
```bash
# 使用配置文件中的设置
./feishu2md wiki-tree

# 或指定 URL（会覆盖配置文件）
./feishu2md wiki-tree https://xxx.feishu.cn/wiki/another_node
```

**特性**：
- ✅ 递归获取所有层级的子文档
- ✅ 自动创建文件夹层级结构
- ✅ 智能跳过有子文档的父级文档
- ✅ 并发下载（最大10个并发）
- ✅ 智能去重，避免重复下载

**输出结构**：
```
dist/
├── 一级目录/
│   ├── 二级文档1.md
│   ├── 子目录/
│   │   ├── 三级文档1.md
│   │   └── img/
│   └── img/
└── 其他文档.md
```

---

## 🔧 飞书 API 配置

### 1. 创建飞书应用

1. 访问 [飞书开发者后台](https://open.feishu.cn/app)
2. 创建**企业自建应用**
3. 记录 **App ID** 和 **App Secret**

### 2. 开通 API 权限

在应用后台开通以下权限：

**必需权限**：
- ✅ `drive:drive:readonly` - 查看云空间文件
- ✅ `drive:file:read` - 读取文件内容  
- ✅ `drive:media:download` - **下载媒体文件（重要）**
- ✅ `wiki:wiki:readonly` - 查看知识库

### 3. 添加协作者权限

对于非公开文档，需要额外配置：

**方法一：知识库全局权限**
1. 为应用添加**机器人能力**并发布
2. 创建飞书群，将机器人添加到群中
3. 在知识库设置中，将该群添加为**管理员**

**方法二：单文档权限**
1. 为应用添加**云文档能力**并发布
2. 在文档的协作设置中，将应用添加为**协作者**

---

## ❓ 常见问题

<details>
<summary><b>Q: 如何获取知识库的 space_id？</b></summary>

A: 
1. 打开知识库
2. 点击右上角 **⚙️ 设置**
3. 查看浏览器地址栏：`https://xxx.feishu.cn/wiki/settings/7474915720537620484`
4. 最后的数字就是 space_id

</details>

<details>
<summary><b>Q: 图片下载失败显示 403 错误？</b></summary>

A: 按顺序检查：
1. 确认已开通 `drive:media:download` 权限
2. 检查应用是否为文档/知识库的协作者
3. 参考上方"添加协作者权限"部分

</details>

<details>
<summary><b>Q: 配置文件在哪里？</b></summary>

A: 默认使用当前目录的 `.env` 文件，也可以通过 `--config` 参数指定其他路径：

```bash
./feishu2md --config /path/to/custom.env document <url>
```

</details>

<details>
<summary><b>Q: 如何跳过图片下载？</b></summary>

A: 使用 `--no-img` 参数：

```bash
./feishu2md document <url> --no-img
```

</details>

<details>
<summary><b>Q: 支持哪些文档类型？</b></summary>

A: 仅支持飞书**新版文档 (docx)**，不支持旧版文档 (docs)

</details>

<details>
<summary><b>Q: 图床上传失败怎么办？</b></summary>

A: 检查以下配置：
1. 密钥是否正确（SecretID/AccessKeyID 和 SecretKey）
2. 存储桶名称和区域是否匹配
3. 存储桶是否有公网访问权限
4. 查看错误日志获取详细信息

</details>

---

## 🛠️ 开发

### 项目结构

```
feishu2md/
├── cmd/                # 命令行入口
│   ├── main.go        # 主程序
│   ├── download.go    # 下载逻辑
│   └── init.go        # 初始化命令
├── core/              # 核心功能
│   ├── client.go      # 飞书 API 客户端
│   ├── config.go      # 配置管理
│   ├── parser.go      # Markdown 解析器
│   └── envloader.go   # 环境变量加载
├── imgbed/            # 图床模块
│   ├── interface.go   # 接口定义
│   ├── oss.go         # 阿里云 OSS
│   ├── cos.go         # 腾讯云 COS
│   └── uploader.go    # 上传逻辑
├── utils/             # 工具函数
│   ├── common.go
│   └── url.go
└── .env.example       # 配置文件示例
```

### 构建

```bash
# 开发构建
go build -o feishu2md ./cmd/...

# 生产构建
make build

# 跨平台编译
GOOS=linux GOARCH=amd64 go build -o feishu2md-linux ./cmd/...
GOOS=windows GOARCH=amd64 go build -o feishu2md.exe ./cmd/...
GOOS=darwin GOARCH=arm64 go build -o feishu2md-darwin-arm64 ./cmd/...
```

---

## 📄 开源协议

本项目基于 [MIT](LICENSE) 协议开源。

## 🙏 致谢

- [chyroc/lark](https://github.com/chyroc/lark) - 飞书 Go SDK
- [88250/lute](https://github.com/88250/lute) - Markdown 处理引擎
- [aliyun/aliyun-oss-go-sdk](https://github.com/aliyun/aliyun-oss-go-sdk) - 阿里云 OSS SDK
- [tencentyun/cos-go-sdk-v5](https://github.com/tencentyun/cos-go-sdk-v5) - 腾讯云 COS SDK

---

## 🌟 贡献

欢迎提交 Issue 和 Pull Request！

---

<div align="center">

**如果觉得有用，请给个 ⭐ Star 支持一下！**

Made with ❤️ by [Perfecto23](https://github.com/Perfecto23)

</div>
