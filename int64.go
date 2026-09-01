package gox

import (
	"fmt"
	"strconv"
)

// Int64 支持字符串形式
type Int64 int64

func (i *Int64) UnmarshalJSON(bytes []byte) (err error) {
	// 去掉可能的引号
	converted := string(bytes)
	if len(converted) >= 2 && converted[0] == '"' && converted[len(converted)-1] == '"' {
		converted = converted[1 : len(converted)-1]
	}

	if parsed, pue := strconv.ParseInt(converted, 10, 64); pue != nil {
		err = fmt.Errorf("invalid uint64 string %q: %w", converted, pue)
	} else {
		*i = Int64(parsed)
	}

	return
}

func (i Int64) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatUint(uint64(i), 10) + `"`), nil
}
