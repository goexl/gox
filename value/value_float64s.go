package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*Float64sValue)(nil)

// Float64sValue 浮点数组值
type Float64sValue struct {
	value []float64
}

// Float64s 浮点数组值
func Float64s(value ...float64) Float64sValue {
	return Float64sValue{
		value: value,
	}
}

func (fv *Float64sValue) Value() interface{} {
	return fv.value
}