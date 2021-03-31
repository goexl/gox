package field

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*Float64Field)(nil)

// Float64Field 描述一个整形字段
type Float64Field struct {
	key.StringKey
	value.Float64Value
}

// NewFloat64Field 创建一个整形字段
func NewFloat64Field(k string, v float64) *Float64Field {
	return &Float64Field{
		StringKey:    key.NewStringKey(k),
		Float64Value: value.NewFloat64Value(v),
	}
}
