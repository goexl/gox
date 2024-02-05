package gox

import (
	"encoding/json"
	"strconv"
	"time"
)

var (
	_ json.Marshaler   = (*Second)(nil)
	_ json.Unmarshaler = (*Second)(nil)
)

// Second 秒
type Second time.Time

func ParseSecond(from string) (second *Second, err error) {
	second = new(Second)
	err = second.UnmarshalJSON([]byte(from))

	return
}

func (s *Second) Time() time.Time {
	return time.Time(*s)
}

func (s *Second) UnmarshalJSON(bytes []byte) (err error) {
	if parsed, pie := strconv.ParseInt(string(bytes), 10, 64); nil != pie {
		err = pie
	} else {
		*s = Second(time.Unix(parsed, 0))
	}

	return
}

func (s Second) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(s).UnixMilli())
}
