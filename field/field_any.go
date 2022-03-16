package field

import (
	`github.com/goexl/gox`
	`github.com/goexl/gox/key`
	`github.com/goexl/gox/value`
)

var _ gox.Key = (*AnyField)(nil)

// AnyField 描述一个任意字段
type AnyField struct {
	key.StringKey
	value.AnyValue
}

// Any 创建一个任意字段
func Any(k string, v interface{}) *AnyField {
	return &AnyField{
		StringKey: key.String(k),
		AnyValue:  value.Any(v),
	}
}
