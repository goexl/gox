package size

import (
	"strconv"
	"strings"
)

type BytesFormatter struct {
	size      *Bytes
	unit      Bytes
	separator rune
}

func NewBytesFormatter(size *Bytes) *BytesFormatter {
	return &BytesFormatter{
		size:      size,
		unit:      BytesMB,
		separator: ' ',
	}
}

func (b *BytesFormatter) KB() *BytesFormatter {
	return b.Unit(BytesKB)
}

func (b *BytesFormatter) MB() *BytesFormatter {
	return b.Unit(BytesMB)
}

func (b *BytesFormatter) GB() *BytesFormatter {
	return b.Unit(BytesGB)
}

func (b *BytesFormatter) TB() *BytesFormatter {
	return b.Unit(BytesTB)
}

func (b *BytesFormatter) PB() *BytesFormatter {
	return b.Unit(BytesPB)
}

func (b *BytesFormatter) EB() *BytesFormatter {
	return b.Unit(BytesEB)
}

func (b *BytesFormatter) Unit(unit Bytes) *BytesFormatter {
	b.unit = unit

	return b
}

func (b *BytesFormatter) Separator(separator rune) *BytesFormatter {
	b.separator = separator

	return b
}

func (b *BytesFormatter) Format() string {
	builder := new(strings.Builder)
	writed := false
	for {
		if *b.size < b.unit {
			break
		}

		writed = true
		switch {
		case *b.size >= BytesEB:
			unit := *b.size / BytesEB
			*b.size -= unit * BytesEB
			builder.WriteString(strconv.FormatInt(int64(unit), 10))
			builder.WriteRune('e')
		case *b.size >= BytesPB:
			unit := *b.size / BytesPB
			*b.size -= unit * BytesPB
			builder.WriteString(strconv.FormatInt(int64(unit), 10))
			builder.WriteRune('p')
		case *b.size >= BytesTB:
			unit := *b.size / BytesTB
			*b.size -= unit * BytesTB
			builder.WriteString(strconv.FormatInt(int64(unit), 10))
			builder.WriteRune('t')
		case *b.size >= BytesGB:
			unit := *b.size / BytesGB
			*b.size -= unit * BytesGB
			builder.WriteString(strconv.FormatInt(int64(unit), 10))
			builder.WriteRune('g')
		case *b.size >= BytesMB:
			unit := *b.size / BytesMB
			*b.size -= unit * BytesMB
			builder.WriteString(strconv.FormatInt(int64(unit), 10))
			builder.WriteRune('m')
		case *b.size >= BytesKB:
			unit := *b.size / BytesKB
			*b.size -= unit * BytesKB
			builder.WriteString(strconv.FormatInt(int64(unit), 10))
			builder.WriteRune('k')
		}

		builder.WriteRune(b.separator)
	}
	if !writed {
		builder.WriteString(strconv.FormatInt(int64(*b.size), 10))
		builder.WriteRune('b')
		builder.WriteRune(b.separator)
	}

	return builder.String()[:builder.Len()-1] // 去掉最后一个分隔符
}
