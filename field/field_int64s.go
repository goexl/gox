package field

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*Int64sField)(nil)

// Int64sField 描述一个整形数组字段
type Int64sField struct {
	key.StringKey
	value.Int64sValue
}

// Int64s 创建一个整形数组字段
func Int64s(k string, v ...int64) *Int64sField {
	return &Int64sField{
		StringKey:   key.String(k),
		Int64sValue: value.Int64s(v...),
	}
}
