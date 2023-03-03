package gox

import (
	"bytes"
)

var _ = StringBuilder

type stringBuilder struct {
	*bytes.Buffer
}

// 创建连写字符串
func StringBuilder(items ...any) (sb *stringBuilder) {
	sb = new(stringBuilder)
	sb.Buffer = new(bytes.Buffer)
	for _, item := range items {
		sb.Append(item)
	}

	return
}

func (sb *stringBuilder) Append(item any) *stringBuilder {
	switch val := item.(type) {
	case []byte:
		sb.Write(val)
	case rune:
		sb.WriteRune(val)
	case string:
		sb.WriteString(val)
	default:
		sb.WriteString(ToString(item))
	}

	return sb
}
