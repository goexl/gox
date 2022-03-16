package value

import (
	`github.com/goexl/gox`
)

var _ gox.Value = (*AnyValue)(nil)

// AnyValue 任意值
type AnyValue struct {
	value interface{}
}

// Any 任意值
func Any(value interface{}) AnyValue {
	return AnyValue{
		value: value,
	}
}

func (av *AnyValue) Value() interface{} {
	return av.value
}
