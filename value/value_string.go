package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*StringValue)(nil)

// StringValue 字符串值
type StringValue struct {
	value *string
}

// String 字符串值
func String(value string) StringValue {
	return Stringp(&value)
}

// Stringp 字符串值
func Stringp(value *string) StringValue {
	return StringValue{
		value: value,
	}
}

func (sv *StringValue) Value() interface{} {
	return sv.value
}
