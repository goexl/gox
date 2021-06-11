package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*BoolValue)(nil)

// BoolValue 布尔值
type BoolValue struct {
	value *bool
}

// Bool 布尔值
func Bool(value bool) BoolValue {
	return Boolp(&value)
}

// Boolp 布尔值
func Boolp(value *bool) BoolValue {
	return BoolValue{
		value: value,
	}
}

func (bv *BoolValue) Value() interface{} {
	return bv.value
}
