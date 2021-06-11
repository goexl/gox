package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*IntsValue)(nil)

// IntsValue 整形值
type IntsValue struct {
	value []int
}

// Ints 整形数组
func Ints(value ...int) IntsValue {
	return IntsValue{
		value: value,
	}
}

func (iv *IntsValue) Value() interface{} {
	return iv.value
}
