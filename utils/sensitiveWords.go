package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/feiin/sensitivewords"
)

var sensitive = sensitivewords.New()

func SensitiveWordInit() error {
	
	SensitiveWordsPath := "./utils/sensitiveWords.txt"
	if err := sensitive.LoadFromFile(SensitiveWordsPath); err != nil {
		return err
	}
	
	return nil
}

// SensitiveWordCheck 敏感词检测
func SensitiveWordCheck(text string, userID int) bool {
	filterText := FilterSpecialChar(text)
	isContain := sensitive.Check(filterText)
	fmt.Printf("isContain: %v\n", isContain)
	if isContain {
		log.Printf("UserID:" + strconv.Itoa(userID) + " | 发表: “" + text + "” | 被视为包含敏感词评论")
	}
	return isContain
}

// 敏感词替换
func SensitiveWordReplace(text string) string {
	filterText := FilterSpecialChar(text)
	replaceText := sensitive.Filter(filterText)
	return replaceText
}

// FilterSpecialChar 过滤特殊字符
func  FilterSpecialChar(text string) string {
	text = strings.ToLower(text)
	text = strings.Replace(text, " ", "", -1) // 去除空格
 
	// 过滤除中英文及数字以外的其他字符
	otherCharReg := regexp.MustCompile("[^\u4e00-\u9fa5a-zA-Z0-9]")
	text = otherCharReg.ReplaceAllString(text, "")
	return text
 }