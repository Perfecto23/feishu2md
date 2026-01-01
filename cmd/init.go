// Package main - 初始化配置文件功能
package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// envTemplate 环境变量配置文件模板
const envTemplate = `# ====================================
# 飞书文档导出工具 - 环境变量配置
# ====================================

# ----------------------------------
# 飞书 API 认证配置（必需）
# ----------------------------------
# 获取方式：https://open.feishu.cn/app
FEISHU_APP_ID=your_app_id_here
FEISHU_APP_SECRET=your_app_secret_here

# ----------------------------------
# 知识库配置（可选）
# ----------------------------------
# 用于 wiki-tree 命令下载知识库子文档

# 知识库空间 ID（必需）
# 从知识库设置页面获取: https://xxx.feishu.cn/wiki/settings/{space_id}
# FEISHU_SPACE_ID=your_space_id_here

# 要下载的文档节点 URL（可选）
# 如果配置了此项，运行 wiki-tree 命令时可以不提供 URL 参数
# FEISHU_FOLDER_TOKEN=https://xxx.feishu.cn/wiki/your_node_token

# ----------------------------------
# 输出配置（可选）
# ----------------------------------
# 文档输出目录
# 默认: ./dist
# OUTPUT_DIR=./dist

# 图片目录（相对于输出目录）
# 默认: img
# IMAGE_DIR=img


# ====================================
# PicGo 图床配置（可选）
# ====================================
# 启用后，下载的图片会通过 PicGo 上传到图床
# 并将 Markdown 中的图片链接替换为图床 URL
#
# 前置条件：
# 1. 安装 PicGo CLI: npm install picgo -g
# 2. 安装压缩插件（可选）: picgo add compress
# 3. 配置图床: picgo set uploader
# 4. 配置压缩（可选）: picgo config plugin compress
#
# PicGo 支持的图床:
# - SM.MS (smms)
# - GitHub (github)
# - 腾讯云 COS (tcyun)
# - 阿里云 OSS (aliyun)
# - 七牛云 (qiniu)
# - 又拍云 (upyun)
# - Imgur (imgur)
# 更多图床可通过 PicGo 插件扩展

# ----------------------------------
# PicGo 开关
# ----------------------------------
# 是否启用 PicGo 图床上传功能
# 值: true/false 或 1/0
PICGO_ENABLED=false


# ----------------------------------
# 使用说明
# ----------------------------------
# 1. 填写上述配置项的值（至少需要填写 FEISHU_APP_ID 和 FEISHU_APP_SECRET）
# 2. 使用配置文件运行:
#    feishu2md document <url> --config .env
#    或者默认会自动加载当前目录的 .env 文件:
#    feishu2md document <url>
# 3. 也可以手动加载环境变量:
#    source .env  (Linux/macOS)
#
# PicGo 图床配置步骤:
# 1. npm install picgo -g           # 安装 PicGo
# 2. picgo add compress             # 安装压缩插件（可选）
# 3. picgo set uploader github      # 配置 GitHub 图床（或其他）
# 4. picgo config plugin compress   # 配置压缩选项（可选）
# 5. 设置 PICGO_ENABLED=true        # 启用 PicGo
#
# 注意: .env 文件包含敏感信息，请勿提交到 Git 仓库
#       本项目的 .gitignore 已默认忽略 .env 文件
`

// handleInitCommand 处理 init 命令
func handleInitCommand(ctx *cli.Context) error {
	force := ctx.Bool("force")
	filename := ".env"

	// 检查文件是否已存在
	if !force {
		if _, err := os.Stat(filename); err == nil {
			return cli.Exit(fmt.Sprintf("❌ 文件 %s 已存在\n"+
				"使用 --force 参数强制覆盖，或手动删除后重试", filename), 1)
		}
	}

	// 写入配置文件
	if err := os.WriteFile(filename, []byte(envTemplate), 0644); err != nil {
		return cli.Exit(fmt.Sprintf("❌ 创建配置文件失败: %v", err), 1)
	}

	// 成功提示
	fmt.Println("✅ 配置文件已创建: " + filename)
	fmt.Println()
	fmt.Println("📝 后续步骤:")
	fmt.Println("  1. 编辑配置文件: vim .env  # 或使用你喜欢的编辑器")
	fmt.Println("  2. 填写必需的配置项（至少需要 FEISHU_APP_ID 和 FEISHU_APP_SECRET）")
	fmt.Println("  3. 开始使用: feishu2md document <url>")
	fmt.Println()
	fmt.Println("💡 提示:")
	fmt.Println("  - 工具会自动加载当前目录的 .env 文件")
	fmt.Println("  - 也可使用 --config 指定其他配置文件: feishu2md --config my.env document <url>")
	fmt.Println("  - 图床功能为可选，不需要可保持 PICGO_ENABLED=false")
	fmt.Println("  - .env 文件已在 .gitignore 中，不会被提交到版本控制")

	return nil
}
