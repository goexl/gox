package field

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*Float32Field)(nil)

// Float32Field 描述一个整形字段
type Float32Field struct {
	key.StringKey
	value.Float32Value
}

// NewFloat32Field 创建一个整形字段
func NewFloat32Field(k string, v float32) *Float32Field {
	return &Float32Field{
		StringKey:    key.NewStringKey(k),
		Float32Value: value.NewFloat32Value(v),
	}
}
