package gox

import (
	"encoding/json"
	"strconv"
	"time"
)

var (
	_ json.Marshaler   = (*Millisecond)(nil)
	_ json.Unmarshaler = (*Millisecond)(nil)
)

// Millisecond 毫秒
type Millisecond time.Time

func ParseMillisecond(from string) (millisecond *Millisecond, err error) {
	millisecond = new(Millisecond)
	err = millisecond.UnmarshalJSON([]byte(from))

	return
}

func (m *Millisecond) Time() time.Time {
	return time.Time(*m)
}

func (m *Millisecond) UnmarshalJSON(bytes []byte) (err error) {
	if parsed, pie := strconv.ParseInt(string(bytes), 10, 64); nil != pie {
		err = pie
	} else {
		*m = Millisecond(time.UnixMilli(parsed))
	}

	return
}

func (m Millisecond) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(m).UnixMilli())
}
