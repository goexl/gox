package gox

import (
	"github.com/goexl/gox/internal/text"
)

var _ = StringBuilder

// StringBuilder 创建字符串构建器
func StringBuilder(items ...any) *text.Builder {
	return text.NewBuilder(items...)
}
