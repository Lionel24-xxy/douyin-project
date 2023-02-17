package utils

import (
	"fmt"
)

// NewFileName 根据UserID+雪花算法生成的id连接成videoName
func NewFileName(userID int64) string {
	node, _ := NewWorker(1)
	randomID := node.NextId()
	return fmt.Sprintf("%d-%d", userID, randomID)
}
