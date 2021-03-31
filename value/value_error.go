package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*ErrorValue)(nil)

// ErrorValue 错误值
type ErrorValue struct {
	value error
}

// NewErrorValue 错误值
func NewErrorValue(err error) ErrorValue {
	return ErrorValue{
		value: err,
	}
}

func (iv *ErrorValue) Value() interface{} {
	return iv.value
}
