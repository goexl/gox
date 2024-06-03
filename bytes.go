package gox

import (
	"github.com/goexl/gox/internal/size"
)

var _ = ParseBytes

// Bytes 字节大小
type Bytes = size.Bytes

// ParseBytes 解析字节大小
func ParseBytes(from string) (Bytes, error) {
	return size.ParseBytes(from)
}
