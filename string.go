package gox

import (
	"encoding/json"
	"fmt"
	"strconv"
)

var _ = ToStrings

// ToStrings 将任意类型转换成字符串列表
func ToStrings(from ...any) (to []string) {
	to = make([]string, 0, len(from))
	for _, value := range from {
		to = append(to, ToString(value))
	}

	return
}

// ToString 将任意类型转换为字符串
func ToString(from any) (to string) {
	switch val := from.(type) {
	case int8:
		to = strconv.FormatInt(int64(val), 10)
	case *int8:
		to = strconv.FormatInt(int64(*val), 10)
	case uint8:
		to = strconv.FormatUint(uint64(val), 10)
	case *uint8:
		to = strconv.FormatUint(uint64(*val), 10)
	case int:
		to = strconv.FormatInt(int64(val), 10)
	case *int:
		to = strconv.FormatInt(int64(*val), 10)
	case uint:
		to = strconv.FormatUint(uint64(val), 10)
	case *uint:
		to = strconv.FormatUint(uint64(*val), 10)
	case int32:
		to = strconv.FormatInt(int64(val), 10)
	case *int32:
		to = strconv.FormatInt(int64(*val), 10)
	case uint32:
		to = strconv.FormatUint(uint64(val), 10)
	case *uint32:
		to = strconv.FormatUint(uint64(*val), 10)
	case int64:
		to = strconv.FormatInt(val, 10)
	case *int64:
		to = strconv.FormatInt(*val, 10)
	case uint64:
		to = strconv.FormatUint(val, 10)
	case *uint64:
		to = strconv.FormatUint(*val, 10)
	case float32:
		to = strconv.FormatFloat(float64(val), 'f', -1, 64)
	case *float32:
		to = strconv.FormatFloat(float64(*val), 'f', -1, 64)
	case float64:
		to = strconv.FormatFloat(val, 'f', -1, 64)
	case *float64:
		to = strconv.FormatFloat(*val, 'f', -1, 64)
	case bool:
		to = strconv.FormatBool(val)
	case *bool:
		to = strconv.FormatBool(*val)
	case string:
		to = val
	case *string:
		to = *val
	case fmt.Stringer:
		to = val.String()
	case json.Marshaler:
		if bytes, err := val.MarshalJSON(); nil == err {
			to = string(bytes)
		}
	default:
		to = fmt.Sprintf("%v", from)
	}

	return
}
