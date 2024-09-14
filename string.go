package gox

import (
	"github.com/goexl/gox/internal/text"
)

// String 字符串操作
func String(from string) *text.String {
	return text.NewString(from)
}
