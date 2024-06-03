package gox

import (
	"github.com/goexl/gox/internal/size"
)

const (
	BytesB  = size.BytesB
	BytesKB = 1024 * BytesB
	BytesMB = 1024 * BytesKB
	BytesGB = 1024 * BytesMB
	BytesTB = 1024 * BytesGB
	BytesPB = 1024 * BytesTB
	BytesEB = 1024 * BytesPB
)

var (
	_ = ParseBytes
	_ = BytesEB
)

// Bytes 字节大小
type Bytes = size.Bytes

// ParseBytes 解析字节大小
func ParseBytes(from string) (Bytes, error) {
	return size.ParseBytes(from)
}
