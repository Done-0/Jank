// Package snowflake 提供雪花算法ID生成工具
// 创建者：Done-0
// 创建时间：2025-08-05
package snowflake

import (
	"fmt"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)

// GenerateID 生成雪花算法 ID
// 返回值：
//
//	int64: 生成的雪花算法 ID
//	error: 操作过程中的错误
func GenerateID() (int64, error) {
	once.Do(func() {
		var err error
		node, err = snowflake.NewNode(1)
		if err != nil {
			fmt.Printf("failed to initialize snowflake node: %v", err)
		}
	})

	switch {
	case node != nil:
		return node.Generate().Int64(), nil

	default:
		// 雪花格式: 41 位时间戳 + 10 位节点 ID + 12 位序列号
		// 标准雪花纪元，节点 ID 1，序列号使用当前纳秒的低 12 位
		ts := time.Now().UnixMilli() - 1288834974657
		nodeID := int64(1)
		seq := time.Now().UnixNano() & 0xFFF

		return (ts << 22) | (nodeID << 12) | seq, nil
	}
}
