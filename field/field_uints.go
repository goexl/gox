package field

import (
	`github.com/goexl/gox`
	`github.com/goexl/gox/key`
	`github.com/goexl/gox/value`
)

var _ gox.Key = (*UintsField)(nil)

// UintsField 描述一个整形数组字段
type UintsField struct {
	key.StringKey
	value.UintsValue
}

// Uints 创建一个整形数组字段
func Uints(k string, v ...uint) *UintsField {
	return &UintsField{
		StringKey:  key.String(k),
		UintsValue: value.Uints(v...),
	}
}
