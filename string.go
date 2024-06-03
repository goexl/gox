package gox

import (
	"github.com/goexl/gox/internal/text"
)

var (
	_ = ToString
	_ = ToStrings
)

// ToString 转换成字符串
func ToString(from any) string {
	return text.ToString(from)
}

// ToStrings 转换成字符串列表
func ToStrings(from ...any) []string {
	return text.ToStrings(from...)
}
