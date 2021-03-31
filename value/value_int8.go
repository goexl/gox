package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*Int8Value)(nil)

// Int8Value 整形值
type Int8Value struct {
	value int8
}

// NewInt8Value 整形值
func NewInt8Value(value int8) Int8Value {
	return Int8Value{
		value: value,
	}
}

func (iv *Int8Value) Value() interface{} {
	return iv.value
}
