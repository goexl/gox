package gox

import (
	"bytes"
)

var _ = StringBuilder

type stringBuilder struct {
	buffer *bytes.Buffer
}

// 创建连写字符串
func StringBuilder(items ...any) (sb *stringBuilder) {
	sb = new(stringBuilder)
	sb.buffer = new(bytes.Buffer)
	for _, item := range items {
		sb.Append(item)
	}

	return
}

func (sb *stringBuilder) String() string {
	return sb.buffer.String()
}

func (sb *stringBuilder) Bytes() []byte {
	return sb.buffer.Bytes()
}

func (sb *stringBuilder) Append(item any) *stringBuilder {
	switch val := item.(type) {
	case []byte:
		sb.buffer.Write(val)
	case rune:
		sb.buffer.WriteRune(val)
	case string:
		sb.buffer.WriteString(val)
	default:
		sb.buffer.WriteString(ToString(item))
	}

	return sb
}
