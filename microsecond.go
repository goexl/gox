package gox

import (
	"encoding/json"
	"strconv"
	"time"
)

var (
	_ json.Marshaler   = (*Microsecond)(nil)
	_ json.Unmarshaler = (*Microsecond)(nil)
	_                  = NewMicrosecond
)

// Microsecond 微秒
type Microsecond struct {
	time.Time
}

// NewMicrosecond 从时间创建微秒
func NewMicrosecond(from time.Time) *Microsecond {
	return &Microsecond{
		Time: from,
	}
}

// ParseMicrosecond 从字符串解析微秒
func ParseMicrosecond(from string) (microsecond *Microsecond, err error) {
	microsecond = new(Microsecond)
	err = microsecond.UnmarshalJSON([]byte(from))

	return
}

func (m *Microsecond) UnmarshalJSON(bytes []byte) (err error) {
	if parsed, pie := strconv.ParseInt(string(bytes), 10, 64); nil != pie {
		err = pie
	} else {
		*m = Microsecond{
			Time: time.UnixMicro(parsed),
		}
	}

	return
}

func (m Microsecond) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.UnixMicro())
}
