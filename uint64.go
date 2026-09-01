package gox

import (
	"fmt"
	"strconv"
)

// Uint64 支持字符串形式
type Uint64 uint64

func (u *Uint64) UnmarshalJSON(bytes []byte) (err error) {
	// 去掉可能的引号
	converted := string(bytes)
	if len(converted) >= 2 && converted[0] == '"' && converted[len(converted)-1] == '"' {
		converted = converted[1 : len(converted)-1]
	}

	if parsed, pue := strconv.ParseUint(converted, 10, 64); pue != nil {
		err = fmt.Errorf("invalid uint64 string %q: %w", converted, pue)
	} else {
		*u = Uint64(parsed)
	}

	return
}

func (u Uint64) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatUint(uint64(u), 10) + `"`), nil
}
