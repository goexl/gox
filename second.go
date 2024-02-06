package gox

import (
	"encoding/json"
	"strconv"
	"time"
)

var (
	_ json.Marshaler   = (*Second)(nil)
	_ json.Unmarshaler = (*Second)(nil)
	_                  = NewSecond
)

// Second 秒
type Second struct {
	time.Time
}

// NewSecond 从时间创建秒
func NewSecond(from time.Time) *Second {
	return &Second{
		Time: from,
	}
}

// ParseSecond 从字符串解析秒
func ParseSecond(from string) (second *Second, err error) {
	second = new(Second)
	err = second.UnmarshalJSON([]byte(from))

	return
}

func (s *Second) UnmarshalJSON(bytes []byte) (err error) {
	if parsed, pie := strconv.ParseInt(string(bytes), 10, 64); nil != pie {
		err = pie
	} else {
		*s = Second{
			Time: time.Unix(parsed, 0),
		}
	}

	return
}

func (s Second) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Unix())
}
