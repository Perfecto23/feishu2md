// Package core 为 feishu2md 提供核心配置和客户端功能
// 此文件处理配置管理，包括从环境变量和CLI参数加载配置
package core

import (
	"os"
)

// Config 表示 feishu2md 应用程序的完整配置
type Config struct {
	Feishu FeishuConfig // 飞书 API 配置
	Output OutputConfig // 输出格式配置
	PicGo  PicGoConfig  // PicGo 图床配置
}

// FeishuConfig 包含飞书/LarkSuite API 凭据
type FeishuConfig struct {
	AppId     string // 飞书应用ID
	AppSecret string // 飞书应用密钥
}

// OutputConfig 包含文档输出格式设置
type OutputConfig struct {
	OutputDir       string // 文档输出目录
	ImageDir        string // 存储下载图片的目录
	TitleAsFilename bool   // 使用文档标题作为文件名而不是令牌
	UseHTMLTags     bool   // 使用HTML标签而不是markdown进行某些格式化
	SkipImgDownload bool   // 跳过下载图片并保留原始链接
	NoBodyTitle     bool   // 禁用正文开头的 H1 标题（因为 frontmatter 已包含 title）
}

// PicGoConfig 包含 PicGo 图床配置
type PicGoConfig struct {
	Enabled bool // 是否启用 PicGo 图床上传
}

// NewConfig 使用提供的应用凭据和默认输出设置创建新配置
func NewConfig(appId, appSecret string) *Config {
	return &Config{
		Feishu: FeishuConfig{
			AppId:     appId,
			AppSecret: appSecret,
		},
		Output: OutputConfig{
			OutputDir:       "./dist", // 默认输出目录
			ImageDir:        "img",    // 默认图片目录
			TitleAsFilename: true,     // 默认使用文档标题作为文件名
			UseHTMLTags:     false,    // 默认使用markdown格式
			SkipImgDownload: false,    // 默认下载图片
		},
	}
}

// LoadConfig 加载配置，优先级：CLI参数 > 环境变量 > 默认值
// 此函数实现一个级联配置系统，每个源可以覆盖前一个源的设置
func LoadConfig(appId, appSecret string) (*Config, error) {
	// 从默认配置开始
	config := NewConfig("", "")

	// 使用环境变量覆盖默认值
	if envAppId := os.Getenv("FEISHU_APP_ID"); envAppId != "" {
		config.Feishu.AppId = envAppId
	}
	if envAppSecret := os.Getenv("FEISHU_APP_SECRET"); envAppSecret != "" {
		config.Feishu.AppSecret = envAppSecret
	}

	// 使用CLI参数覆盖（最高优先级）
	if appId != "" {
		config.Feishu.AppId = appId
	}
	if appSecret != "" {
		config.Feishu.AppSecret = appSecret
	}

	// 加载输出配置（从环境变量）
	loadOutputConfig(config)

	// 加载 PicGo 配置（从环境变量）
	loadPicGoConfig(config)

	return config, nil
}

// loadOutputConfig 从环境变量加载输出配置
func loadOutputConfig(config *Config) {
	// 输出目录
	if outputDir := os.Getenv("OUTPUT_DIR"); outputDir != "" {
		config.Output.OutputDir = outputDir
	}
	// 图片目录
	if imageDir := os.Getenv("IMAGE_DIR"); imageDir != "" {
		config.Output.ImageDir = imageDir
	}
}

// loadPicGoConfig 从环境变量加载 PicGo 配置
func loadPicGoConfig(config *Config) {
	// 检查是否启用 PicGo
	if enabled := os.Getenv("PICGO_ENABLED"); enabled == "true" || enabled == "1" {
		config.PicGo.Enabled = true
	}
}
