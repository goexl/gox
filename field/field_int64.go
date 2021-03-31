package field

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*Int64Field)(nil)

// Int64Field 描述一个整形字段
type Int64Field struct {
	key.StringKey
	value.Int64Value
}

// Int64 创建一个整形字段
func Int64(k string, v int64) *Int64Field {
	return &Int64Field{
		StringKey:  key.String(k),
		Int64Value: value.Int64(v),
	}
}
