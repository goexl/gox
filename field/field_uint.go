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

// NewUintField 创建一个整形字段
func NewUintField(k string, v uint) *UintField {
	return &UintField{
		StringKey: key.NewStringKey(k),
		UintValue: value.NewUintValue(v),
	}
}
