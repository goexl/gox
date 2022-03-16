package value

import (
	`github.com/goexl/gox`
)

var _ gox.Value = (*UintsValue)(nil)

// UintsValue 整形数组值
type UintsValue struct {
	value []uint
}

// Uints 整形数组值
func Uints(value ...uint) UintsValue {
	return UintsValue{
		value: value,
	}
}

func (iv *UintsValue) Value() interface{} {
	return iv.value
}
