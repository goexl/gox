package field

import (
	`github.com/goexl/gox`
	`github.com/goexl/gox/key`
	`github.com/goexl/gox/value`
)

var _ gox.Key = (*IntsField)(nil)

// IntsField 描述一个整形字段
type IntsField struct {
	key.StringKey
	value.IntsValue
}

// Ints 创建一个整形字段
func Ints(k string, v ...int) *IntsField {
	return &IntsField{
		StringKey: key.String(k),
		IntsValue: value.Ints(v...),
	}
}
