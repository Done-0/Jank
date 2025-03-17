package utils

import "os"

// MkDir 创建目录
func MkDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}
