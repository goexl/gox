package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*Float64Value)(nil)

// Float64Value 浮点值
type Float64Value struct {
	value *float64
}

// Float64 浮点值
func Float64(value float64) Float64Value {
	return Float64p(&value)
}

// Float64p 浮点值
func Float64p(value *float64) Float64Value {
	return Float64Value{
		value: value,
	}
}

func (fv *Float64Value) Value() interface{} {
	return fv.value
}
