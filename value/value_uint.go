package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*UintValue)(nil)

// UintValue 整形值
type UintValue struct {
	value uint
}

// NewUintValue 整形值
func NewUintValue(value uint) UintValue {
	return UintValue{
		value: value,
	}
}

func (iv *UintValue) Value() interface{} {
	return iv.value
}
