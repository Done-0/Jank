// Package plugin 插件工具函数
// 创建者：Done-0
// 创建时间：2025-08-08
package plugin

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Done-0/jank/configs"
)

// GenerateBinaryPath 生成插件二进制文件路径
func GenerateBinaryPath(pluginPath, pluginID, configBinary string) string {
	if configBinary == "" {
		cfgs, err := configs.GetConfig()
		if err != nil {
			log.Fatalf("failed to get config: %v", err)
		}

		return filepath.Join(pluginPath, cfgs.PluginConfig.PluginBinDir, pluginID)
	}
	return filepath.Join(pluginPath, configBinary)
}

// CheckBinaryExists 检查二进制文件是否存在
func CheckBinaryExists(binaryPath string) bool {
	_, err := os.Stat(binaryPath)
	return err == nil
}

// ExecuteBuildScript 执行插件构建脚本
func ExecuteBuildScript(pluginPath string) error {
	cfgs, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	buildScript := filepath.Join(pluginPath, cfgs.PluginConfig.BuildScriptDir, cfgs.PluginConfig.BuildScriptFile)

	// 检查构建脚本是否存在，没有则视为无需构建
	if _, err := os.Stat(buildScript); os.IsNotExist(err) {
		return nil
	}

	// 设置脚本执行权限
	if err := os.Chmod(buildScript, 0755); err != nil {
		return err
	}

	timeout := time.Duration(cfgs.PluginConfig.BuildTimeoutMinutes) * time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 获取绝对路径
	absBuildScript, err := filepath.Abs(buildScript)
	if err != nil {
		return err
	}

	absPluginPath, err := filepath.Abs(pluginPath)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "bash", absBuildScript)
	cmd.Dir = absPluginPath
	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Plugin build failed: %v\nOutput: %s", err, string(output))
	}

	return err
}
