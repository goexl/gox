package size

import (
	"strconv"
	"strings"
	"unsafe"

	"github.com/goexl/gox/internal/constant"
	"github.com/goexl/gox/internal/text"
)

// Bytes 字节大小
type Bytes int64

// ParseBytes 解析字节大小
func ParseBytes(from string) (bytes Bytes, err error) {
	// 逐步解析各个容量
	for _, volume := range strings.Split(from, constant.StringSpace) {
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
		if capacity, err = strconv.ParseFloat(num, constant.Float64Size); nil != err {
			return
		}

		switch strings.ToLower(unit) {
		case "b":
			bytes += Bytes(capacity)
		case "kb", "k":
			bytes += Bytes(capacity) * BytesKB
		case "mb", "m":
			bytes += Bytes(capacity) * BytesMB
		case "gb", "g":
			bytes += Bytes(capacity) * BytesGB
		case "tb", "t":
			bytes += Bytes(capacity) * BytesTB
		case "pb", "p":
			bytes += Bytes(capacity) * BytesPB
		case "eb", "e":
			bytes += Bytes(capacity) * BytesEB
		}
	}

	return
}

func (b Bytes) String() string {
	return b.Formatter().Format()
}

func (b Bytes) Bit() int64 {
	return int64(b) * 8
}

func (b Bytes) Byte() int64 {
	return int64(b)
}

func (b *Bytes) Formatter() *BytesFormatter {
	return NewBytesFormatter(b)
}

func (b Bytes) MarshalJSON() (bytes []byte, err error) {
	bytes = text.NewBuilder(constant.StringQuote, b.String(), constant.StringQuote).Bytes()

	return
}

func (b *Bytes) UnmarshalJSON(bytes []byte) (err error) {
	json := *(*string)(unsafe.Pointer(&bytes))
	*b, err = ParseBytes(strings.ReplaceAll(json, constant.StringQuote, constant.StringEmpty))

	return
}
