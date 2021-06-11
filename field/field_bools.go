package field

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*BoolsField)(nil)

// BoolsField 描述一个布尔数组字段
type BoolsField struct {
	key.StringKey
	value.BoolsValue
}

// Bools 创建一个布尔数组字段
func Bools(k string, v ...bool) *BoolsField {
	return &BoolsField{
		StringKey:  key.String(k),
		BoolsValue: value.Bools(v...),
	}
}
