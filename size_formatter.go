package gox

import (
	"strconv"
	"strings"
)

type sizeFormatter struct {
	size      *Size
	unit      Size
	separator rune
}

func newSizeFormatter(size *Size) *sizeFormatter {
	return &sizeFormatter{
		size:      size,
		unit:      SizeMB,
		separator: ' ',
	}
}

func (sf *sizeFormatter) KB() *sizeFormatter {
	return sf.Unit(SizeKB)
}

func (sf *sizeFormatter) MB() *sizeFormatter {
	return sf.Unit(SizeMB)
}

func (sf *sizeFormatter) GB() *sizeFormatter {
	return sf.Unit(SizeGB)
}

func (sf *sizeFormatter) TB() *sizeFormatter {
	return sf.Unit(SizeTB)
}

func (sf *sizeFormatter) PB() *sizeFormatter {
	return sf.Unit(SizePB)
}

func (sf *sizeFormatter) EB() *sizeFormatter {
	return sf.Unit(SizeEB)
}

func (sf *sizeFormatter) Unit(unit Size) *sizeFormatter {
	sf.unit = unit

	return sf
}

func (sf *sizeFormatter) Separator(separator rune) *sizeFormatter {
	sf.separator = separator

	return sf
}

func (sf *sizeFormatter) Format() string {
	var sb strings.Builder
	for {
		if *sf.size < sf.unit {
			break
		}

		switch {
		case *sf.size >= SizeEB:
			unit := *sf.size / SizeEB
			*sf.size -= unit * SizeEB
			sb.WriteString(strconv.Itoa(int(unit)))
			sb.WriteRune('e')
		case *sf.size >= SizePB:
			unit := *sf.size / SizePB
			*sf.size -= unit * SizePB
			sb.WriteString(strconv.Itoa(int(unit)))
			sb.WriteRune('p')
		case *sf.size >= SizeTB:
			unit := *sf.size / SizeTB
			*sf.size -= unit * SizeTB
			sb.WriteString(strconv.Itoa(int(unit)))
			sb.WriteRune('t')
		case *sf.size >= SizeGB:
			unit := *sf.size / SizeGB
			*sf.size -= unit * SizeGB
			sb.WriteString(strconv.Itoa(int(unit)))
			sb.WriteRune('g')
		case *sf.size >= SizeMB:
			unit := *sf.size / SizeMB
			*sf.size -= unit * SizeMB
			sb.WriteString(strconv.Itoa(int(unit)))
			sb.WriteRune('m')
		case *sf.size >= SizeKB:
			unit := *sf.size / SizeKB
			*sf.size -= unit * SizeKB
			sb.WriteString(strconv.Itoa(int(unit)))
			sb.WriteRune('k')
		}

		sb.WriteRune(sf.separator)
	}

	return sb.String()[:sb.Len()-1] // 去掉最后一个分隔符
}
