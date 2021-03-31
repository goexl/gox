package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*Int64Value)(nil)

// Int64Value 整形值
type Int64Value struct {
	value int64
}

// Int64 整形值
func Int64(value int64) Int64Value {
	return Int64Value{
		value: value,
	}
}

func (iv *Int64Value) Value() interface{} {
	return iv.value
}
