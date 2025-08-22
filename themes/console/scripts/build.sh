#!/bin/bash

set -e

# 检查是否在主题根目录
if [[ ! -f "theme.json" ]]; then
    echo "Error: theme.json not found. Please run from theme root directory." >&2
    exit 1
fi

if [[ ! -f "package.json" ]]; then
    echo "Error: package.json not found" >&2
    exit 1
fi

# 读取主题配置
INDEX_FILE_PATH=$(jq -r '.index_file_path' theme.json)

# 提取构建输出目录
BUILD_DIR=$(echo "$INDEX_FILE_PATH" | sed 's|^/||' | cut -d'/' -f1)

# 检测包管理器
if command -v pnpm >/dev/null 2>&1 && [[ -f "pnpm-lock.yaml" ]]; then
    PACKAGE_MANAGER="pnpm"
elif command -v yarn >/dev/null 2>&1 && [[ -f "yarn.lock" ]]; then
    PACKAGE_MANAGER="yarn"
elif command -v npm >/dev/null 2>&1; then
    PACKAGE_MANAGER="npm"
else
    echo "Error: No package manager found" >&2
    exit 1
fi

# 安装依赖
if [[ ! -d "node_modules" ]]; then
    $PACKAGE_MANAGER install
fi

# 构建项目
$PACKAGE_MANAGER run build

# 验证构建结果
if [[ ! -d "$BUILD_DIR" || ! -f "$BUILD_DIR/index.html" ]]; then
    echo "Error: Build failed" >&2
    exit 1
fi
