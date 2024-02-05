package gox

import (
	"encoding/json"
	"strconv"
	"time"
)

var (
	_ json.Marshaler   = (*Microsecond)(nil)
	_ json.Unmarshaler = (*Microsecond)(nil)
)

// Microsecond 微秒
type Microsecond time.Time

func ParseMicrosecond(from string) (microsecond *Microsecond, err error) {
	microsecond = new(Microsecond)
	err = microsecond.UnmarshalJSON([]byte(from))

	return
}

func (m *Microsecond) Time() time.Time {
	return time.Time(*m)
}

func (m *Microsecond) UnmarshalJSON(bytes []byte) (err error) {
	if parsed, pie := strconv.ParseInt(string(bytes), 10, 64); nil != pie {
		err = pie
	} else {
		*m = Microsecond(time.UnixMicro(parsed))
	}

	return
}

func (m Microsecond) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(m).UnixMicro())
}
