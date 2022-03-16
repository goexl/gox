package field

import (
	`github.com/goexl/gox`
	`github.com/goexl/gox/key`
	`github.com/goexl/gox/value`
)

var _ gox.Key = (*Float64Field)(nil)

// Float64Field 描述一个整形字段
type Float64Field struct {
	key.StringKey
	value.Float64Value
}

// Float64 创建一个整形字段
func Float64(k string, v float64) *Float64Field {
	return Float64p(k, &v)
}

// Float64p 创建一个整形字段
func Float64p(k string, v *float64) *Float64Field {
	return &Float64Field{
		StringKey:    key.String(k),
		Float64Value: value.Float64p(v),
	}
}
