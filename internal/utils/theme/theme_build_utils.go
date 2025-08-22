// Package theme 主题构建工具函数
package theme

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Done-0/jank/configs"
)

// ExecuteBuildScript 执行主题构建脚本
func ExecuteBuildScript(themePath string) error {
	cfgs, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	buildScript := filepath.Join(themePath, cfgs.ThemeConfig.BuildScriptDir, cfgs.ThemeConfig.BuildScriptFile)

	// 检查构建脚本是否存在，没有则视为静态主题
	if _, err := os.Stat(buildScript); os.IsNotExist(err) {
		return nil
	}

	timeout := time.Duration(cfgs.ThemeConfig.BuildTimeoutMinutes) * time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 获取绝对路径
	absBuildScript, err := filepath.Abs(buildScript)
	if err != nil {
		return err
	}

	absThemePath, err := filepath.Abs(themePath)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "bash", absBuildScript)
	cmd.Dir = absThemePath
	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Theme build failed: %v\nOutput: %s", err, string(output))
	}

	return err
}
