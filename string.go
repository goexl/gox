package gox

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

var (
	_ = Strings
	_ = HasAnyPrefix
	_ = HasAllPrefix
	_ = HasAnySuffix
	_ = HasAllSuffix
)

// Strings 将任意类型转换成字符串列表
func Strings(from ...any) (to []string) {
	to = make([]string, 0, len(from))
	for _, value := range from {
		to = append(to, String(value))
	}

	return
}

// String 将任意类型转换为字符串
func String(from any) (to string) {
	switch _value := from.(type) {
	case int8:
		to = strconv.FormatInt(int64(_value), 10)
	case *int8:
		to = strconv.FormatInt(int64(*_value), 10)
	case uint8:
		to = strconv.FormatUint(uint64(_value), 10)
	case *uint8:
		to = strconv.FormatUint(uint64(*_value), 10)
	case int:
		to = strconv.FormatInt(int64(_value), 10)
	case *int:
		to = strconv.FormatInt(int64(*_value), 10)
	case uint:
		to = strconv.FormatUint(uint64(_value), 10)
	case *uint:
		to = strconv.FormatUint(uint64(*_value), 10)
	case int32:
		to = strconv.FormatInt(int64(_value), 10)
	case *int32:
		to = strconv.FormatInt(int64(*_value), 10)
	case uint32:
		to = strconv.FormatUint(uint64(_value), 10)
	case *uint32:
		to = strconv.FormatUint(uint64(*_value), 10)
	case int64:
		to = strconv.FormatInt(_value, 10)
	case *int64:
		to = strconv.FormatInt(*_value, 10)
	case uint64:
		to = strconv.FormatUint(_value, 10)
	case *uint64:
		to = strconv.FormatUint(*_value, 10)
	case float32:
		to = strconv.FormatFloat(float64(_value), 'f', -1, 64)
	case *float32:
		to = strconv.FormatFloat(float64(*_value), 'f', -1, 64)
	case float64:
		to = strconv.FormatFloat(_value, 'f', -1, 64)
	case *float64:
		to = strconv.FormatFloat(*_value, 'f', -1, 64)
	case bool:
		to = strconv.FormatBool(_value)
	case *bool:
		to = strconv.FormatBool(*_value)
	case string:
		to = _value
	case *string:
		to = *_value
	case fmt.Stringer:
		to = _value.String()
	case json.Marshaler:
		if bytes, err := _value.MarshalJSON(); nil == err {
			to = string(bytes)
		}
	default:
		to = fmt.Sprintf("%v", from)
	}

	return
}

// HasAnyPrefix 判断是否包含任意一个前缀
func HasAnyPrefix(str string, checks ...string) (has bool) {
	has = false
	for _, check := range checks {
		if strings.HasPrefix(str, check) {
			has = true
		}

		if has {
			break
		}
	}

	return
}

// HasAllPrefix 判断是否包含所有前缀
func HasAllPrefix(str string, checks ...string) (has bool) {
	has = true
	for _, check := range checks {
		if !strings.HasPrefix(str, check) {
			has = false
		}

		if !has {
			break
		}
	}

	return
}

// HasAnySuffix 判断是否包含任意一个后缀
func HasAnySuffix(str string, checks ...string) (has bool) {
	has = false
	for _, check := range checks {
		if strings.HasSuffix(str, check) {
			has = true
		}

		if has {
			break
		}
	}

	return
}

// HasAllSuffix 判断是否包含所有后缀
func HasAllSuffix(str string, checks ...string) (has bool) {
	has = true
	for _, check := range checks {
		if !strings.HasSuffix(str, check) {
			has = false
		}

		if !has {
			break
		}
	}

	return
}
