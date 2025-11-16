// Package core - 环境变量配置文件加载器
package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LoadEnvFile 从指定路径加载环境变量配置文件
// 支持 .env 格式的配置文件
func LoadEnvFile(filepath string) error {
	// 检查文件是否存在
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return fmt.Errorf("配置文件不存在: %s", filepath)
	}

	// 打开文件
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("无法打开配置文件: %w", err)
	}
	defer file.Close()

	// 逐行读取并设置环境变量
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// 跳过空行和注释行
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// 解析 KEY=VALUE 格式
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			// 忽略格式不正确的行，不报错
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// 移除值两端的引号（如果有）
		value = strings.Trim(value, "\"'")

		// 只有当环境变量未设置时才设置（命令行/系统环境变量优先）
		if os.Getenv(key) == "" {
			if err := os.Setenv(key, value); err != nil {
				return fmt.Errorf("设置环境变量失败 %s: %w", key, err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	return nil
}

// LoadEnvFileIfExists 加载配置文件，如果文件不存在则忽略
func LoadEnvFileIfExists(filepath string) error {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil // 文件不存在，不报错
	}
	return LoadEnvFile(filepath)
}

