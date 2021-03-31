package field

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*Int8Field)(nil)

// Int8Field 描述一个整形字段
type Int8Field struct {
	key.StringKey
	value.Int8Value
}

// NewInt8Field 创建一个整形字段
func NewInt8Field(k string, v int8) *Int8Field {
	return &Int8Field{
		StringKey: key.NewStringKey(k),
		Int8Value: value.NewInt8Value(v),
	}
}
