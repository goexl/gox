package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*Float64Value)(nil)

// Float64Value 浮点值
type Float64Value struct {
	value float64
}

// NewFloat64Value 浮点值
func NewFloat64Value(value float64) Float64Value {
	return Float64Value{
		value: value,
	}
}

func (iv *Float64Value) Value() interface{} {
	return iv.value
}
