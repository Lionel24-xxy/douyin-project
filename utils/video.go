package utils

import (
	"fmt"
)

var (
	defaultIP   = ""
	defaultPort = 8080
)

// NewFileName 根据UserID+雪花算法生成的id连接成videoName
func NewFileName(userID int64) string {
	node, _ := NewWorker(1)
	randomID := node.NextId()
	return fmt.Sprintf("%d-%d", userID, randomID)
}

// GetFileUrl 通过filename，将它改为url形式
func GetFileUrl(filename string) string {
	base := fmt.Sprintf("http://%s:%d/static/%s", defaultIP, defaultPort, filename)
	return base
}
