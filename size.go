package gox

import (
	"strconv"
	"strings"
)

const (
	SizeB  Size = 1
	SizeKB      = 1024 * SizeB
	SizeMB      = 1024 * SizeKB
	SizeGB      = 1024 * SizeMB
	SizeTB      = 1024 * SizeGB
	SizePB      = 1024 * SizeTB
	SizeEB      = 1024 * SizePB
)

var _ = ParseSize

// Size 字节大小
type Size int64

// ParseSize 解析字节大小
func ParseSize(str string) (size Size, err error) {
	volumes := strings.Split(str, StringSpace)
	if 0 == len(volumes) {
		err = newField(errorInvalidFormat, newStringField(`string`, str))
	}
	if nil != err {
		return
	}

	// 逐步解析各个容量
	for _, volume := range volumes {
		var unit, num string
		length := len(volume)
		check := volume[length-2]
		if '0' <= check && '9' >= check { // 如果倒数第二位是数字，表示是以最小单位
			unit = volume[length-1:]
			num = volume[:length-1]
		} else { // 表示是形如KB，MB，GB等单位
			unit = volume[length-2:]
			num = volume[:length-2]
		}

		// 计算数字大小
		var capacity float64
		if capacity, err = strconv.ParseFloat(num, Float64Size); nil != err {
			return
		}

		switch unit {
		case `b`:
			size += Size(capacity)
		case `kb`, `k`:
			size += Size(capacity) * SizeKB
		case `mb`, `m`:
			size += Size(capacity) * SizeMB
		case `gb`, `g`:
			size += Size(capacity) * SizeGB
		case `tb`, `t`:
			size += Size(capacity) * SizeTB
		case `pb`, `p`:
			size += Size(capacity) * SizePB
		case `eb`, `e`:
			size += Size(capacity) * SizeEB
		}
	}

	return
}

func (s Size) String() string {
	return s.Format()
}

func (s Size) Format(opts ...sizeOption) string {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.applySize(_options)
	}

	var sb strings.Builder
	for {
		if s < _options.unit {
			break
		}

		switch {
		case s >= SizeEB:
			unit := s / SizeEB
			s -= unit * SizeEB
			sb.WriteString(strconv.Itoa(int(unit)))
			sb.WriteRune('e')
		case s >= SizePB:
			unit := s / SizePB
			s -= unit * SizePB
			sb.WriteString(strconv.Itoa(int(unit)))
			sb.WriteRune('p')
		case s >= SizeTB:
			unit := s / SizeTB
			s -= unit * SizeTB
			sb.WriteString(strconv.Itoa(int(unit)))
			sb.WriteRune('t')
		case s >= SizeGB:
			unit := s / SizeGB
			s -= unit * SizeGB
			sb.WriteString(strconv.Itoa(int(unit)))
			sb.WriteRune('g')
		case s >= SizeMB:
			unit := s / SizeMB
			s -= unit * SizeMB
			sb.WriteString(strconv.Itoa(int(unit)))
			sb.WriteRune('m')
		case s >= SizeKB:
			unit := s / SizeKB
			s -= unit * SizeKB
			sb.WriteString(strconv.Itoa(int(unit)))
			sb.WriteRune('k')
		}

		sb.WriteRune(_options.separator)
	}

	return sb.String()[:sb.Len()-1] // 去掉最后一个分隔符
}

func (s Size) MarshalJSON() (bytes []byte, err error) {
	bytes = []byte(s.String())

	return
}

func (s *Size) UnmarshalJSON(bytes []byte) (err error) {
	*s, err = ParseSize(string(bytes))

	return
}
