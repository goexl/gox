package field

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*IntField)(nil)

// IntField 描述一个整形字段
type IntField struct {
	key.StringKey
	value.IntValue
}

// NewIntField 创建一个整形字段
func NewIntField(k string, v int) *IntField {
	return &IntField{
		StringKey: key.NewStringKey(k),
		IntValue:  value.NewIntValue(v),
	}
}
