package value

import (
	`fmt`

	`github.com/goexl/gox`
)

var (
	_           = Uint
	_           = Uintp
	_ gox.Value = (*UintValue)(nil)
)

// UintValue 整形值
type UintValue struct {
	value *uint
}

// Uint 整形值
func Uint(value uint) UintValue {
	return Uintp(&value)
}

// Uintp 整形值
func Uintp(value *uint) UintValue {
	return UintValue{
		value: value,
	}
}

func (uv *UintValue) Value() interface{} {
	return uv.value
}

func (uv *UintValue) String() string {
	return fmt.Sprintf(`%v`, *uv.value)
}
