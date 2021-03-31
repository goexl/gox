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

// Int8 创建一个整形字段
func Int8(k string, v int8) *Int8Field {
	return &Int8Field{
		StringKey: key.String(k),
		Int8Value: value.Int8(v),
	}
}
