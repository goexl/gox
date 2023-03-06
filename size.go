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
func ParseSize(from string) (size Size, err error) {
	// 逐步解析各个容量
	for _, volume := range strings.Split(from, stringSpace) {
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
		if capacity, err = strconv.ParseFloat(num, float64Size); nil != err {
			return
		}

		switch strings.ToLower(unit) {
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

func (s *Size) String() string {
	return s.Formatter().Format()
}

func (s *Size) Bit() int64 {
	return int64(*s) * 8
}

func (s *Size) Bit32() int32 {
	return int32(*s) * 8
}

func (s *Size) Formatter() *sizeFormatter {
	return newSizeFormatter(s)
}

func (s *Size) MarshalJSON() (bytes []byte, err error) {
	bytes = []byte(s.String())

	return
}

func (s *Size) UnmarshalJSON(bytes []byte) (err error) {
	*s, err = ParseSize(string(bytes))

	return
}
