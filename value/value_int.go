package value

import (
	`fmt`

	`github.com/goexl/gox`
)

var (
	_           = Int
	_           = Intp
	_ gox.Value = (*IntValue)(nil)
)

// IntValue 整形值
type IntValue struct {
	value *int
}

// Int 整形值
func Int(value int) IntValue {
	return Intp(&value)
}

// Intp 整形值
func Intp(value *int) IntValue {
	return IntValue{
		value: value,
	}
}

func (iv *IntValue) Value() interface{} {
	return iv.value
}

func (iv *IntValue) String() string {
	return fmt.Sprintf(`%v`, *iv.value)
}
