package text

import (
	"bytes"
)

type Builder struct {
	buffer *bytes.Buffer
}

// NewBuilder 创建连写字符串
func NewBuilder(items ...any) (builder *Builder) {
	builder = new(Builder)
	builder.buffer = new(bytes.Buffer)
	for _, item := range items {
		builder.Append(item)
	}

	return
}

func (b *Builder) String() string {
	return b.buffer.String()
}

func (b *Builder) Bytes() []byte {
	return b.buffer.Bytes()
}

func (b *Builder) Append(item any) *Builder {
	switch typed := item.(type) {
	case []byte:
		b.buffer.Write(typed)
	case rune:
		b.buffer.WriteRune(typed)
	case string:
		b.buffer.WriteString(typed)
	default:
		b.buffer.WriteString(ToString(item))
	}

	return b
}
