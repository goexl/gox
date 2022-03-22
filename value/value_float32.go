package value

import (
	`fmt`

	`github.com/goexl/gox`
)

var (
	_           = Float32
	_           = Float32p
	_ gox.Value = (*Float32Value)(nil)
)

// Float32Value 浮点值
type Float32Value struct {
	value *float32
}

// Float32 浮点值
func Float32(value float32) Float32Value {
	return Float32p(&value)
}

// Float32p 浮点值
func Float32p(value *float32) Float32Value {
	return Float32Value{
		value: value,
	}
}

func (fv *Float32Value) Value() interface{} {
	return fv.value
}

func (fv *Float32Value) String() string {
	return fmt.Sprintf(`%v`, *fv.value)
}
