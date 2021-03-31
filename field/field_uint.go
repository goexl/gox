package field

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*UintField)(nil)

// UintField 描述一个整形字段
type UintField struct {
	key.StringKey
	value.UintValue
}

// Uint 创建一个整形字段
func Uint(k string, v uint) *UintField {
	return &UintField{
		StringKey: key.String(k),
		UintValue: value.Uint(v),
	}
}
