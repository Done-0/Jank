#!/bin/bash

set -e

# 检查是否在插件根目录
if [[ ! -f "plugin.json" ]]; then
    echo "Error: plugin.json not found. Please run from plugin root directory." >&2
    exit 1
fi

# 读取插件配置
BINARY_PATH=$(jq -r '.binary' plugin.json)

# 创建输出目录
mkdir -p "$(dirname "$BINARY_PATH")"

# 查找 Go 源文件
if [[ -f "main.go" ]]; then
    MAIN_FILE="main.go"
else
    GO_FILES=(*.go)
    if [[ -f "${GO_FILES[0]}" ]]; then
        MAIN_FILE="${GO_FILES[0]}"
    else
        echo "Error: No Go source files found" >&2
        exit 1
    fi
fi

# 编译插件
go build -o "$BINARY_PATH" "$MAIN_FILE"

# 验证构建结果
if [[ ! -f "$BINARY_PATH" ]]; then
    echo "Error: Build failed" >&2
    exit 1
fi
