package token

import (
	"regexp"
	"strings"
)

func Named(from string) (to []string) {
	// 用正则表达式处理下划线和中划线
	re := regexp.MustCompile(`[-_]+`)
	processed := re.ReplaceAllString(from, " ")

	// 用正则表达式处理驼峰命名
	re = regexp.MustCompile(`([a-z])([A-Z])`)
	processed = re.ReplaceAllString(processed, `${1} ${2}`)

	// 分割字符串为单词
	to = strings.Fields(processed)

	return
}
