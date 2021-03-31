package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*Float32Value)(nil)

// Float32Value 浮点值
type Float32Value struct {
	value float32
}

// Float32 浮点值
func Float32(value float32) Float32Value {
	return Float32Value{
		value: value,
	}
}

func (iv *Float32Value) Value() interface{} {
	return iv.value
}
