package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*Int8Value)(nil)

// Int8Value 整形值
type Int8Value struct {
	value *int8
}

// Int8 整形值
func Int8(value int8) Int8Value {
	return Int8p(&value)
}

// Int8p 整形值
func Int8p(value *int8) Int8Value {
	return Int8Value{
		value: value,
	}
}

func (iv *Int8Value) Value() interface{} {
	return iv.value
}
