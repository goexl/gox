package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*UintValue)(nil)

// UintValue 整形值
type UintValue struct {
	value *uint
}

// Uint 整形值
func Uint(value uint) UintValue {
	return Uintp(&value)
}

// Uintp 整形值
func Uintp(value *uint) UintValue {
	return UintValue{
		value: value,
	}
}

func (iv *UintValue) Value() interface{} {
	return iv.value
}
