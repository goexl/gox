package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*BoolsValue)(nil)

// BoolsValue 布尔数组
type BoolsValue struct {
	value []bool
}

// Bools 布尔数组
func Bools(value ...bool) BoolsValue {
	return BoolsValue{
		value: value,
	}
}

func (bv *BoolsValue) Value() interface{} {
	return bv.value
}
