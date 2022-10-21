package field

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/key"
	"github.com/goexl/gox/value"
)

var (
	_         = Uint
	_ gox.Key = (*UintField)(nil)
)

// UintField 描述一个整形字段
type UintField struct {
	key.StringKey
	value.UintValue
}

// Uint 创建一个整形字段
func Uint(k string, v uint) *UintField {
	return Uintp(k, &v)
}

// Uintp 创建一个整形字段
func Uintp(k string, v *uint) *UintField {
	return &UintField{
		StringKey: key.String(k),
		UintValue: value.Uintp(v),
	}
}
