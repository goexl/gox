package gox

import (
	"encoding/json"
	"strconv"
	"time"
)

var (
	_ json.Marshaler   = (*Nanosecond)(nil)
	_ json.Unmarshaler = (*Nanosecond)(nil)
	_                  = NewNanosecond
)

// Nanosecond 纳秒
type Nanosecond struct {
	time.Time
}

// NewNanosecond 从时间创建毫秒
func NewNanosecond(from time.Time) *Nanosecond {
	return &Nanosecond{
		Time: from,
	}
}

// ParseNanosecond 从字符串解析毫秒
func ParseNanosecond(from string) (nanosecond *Nanosecond, err error) {
	nanosecond = new(Nanosecond)
	err = nanosecond.UnmarshalJSON([]byte(from))

	return
}

func (n *Nanosecond) UnmarshalJSON(bytes []byte) (err error) {
	if parsed, pie := strconv.ParseInt(string(bytes), 10, 64); nil != pie {
		err = pie
	} else {
		*n = Nanosecond{
			Time: time.Unix(0, parsed),
		}
	}

	return
}

func (n Nanosecond) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.UnixMilli())
}
