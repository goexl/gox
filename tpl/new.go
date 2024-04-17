package tpl

import (
	"github.com/goexl/gox/tpl/internal"
)

var _ = New

// New 创建模板
func New(input string) *internal.Builder {
	return internal.NewBuilder(input)
}
