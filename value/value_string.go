package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*StringValue)(nil)

// StringValue 字符串值
type StringValue struct {
	value string
}

// String 字符串值
func String(value string) StringValue {
	return StringValue{
		value: value,
	}
}

func (iv *StringValue) Value() interface{} {
	return iv.value
}
