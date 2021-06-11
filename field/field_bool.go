package field

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*BoolField)(nil)

// BoolField 描述一个布尔字段
type BoolField struct {
	key.StringKey
	value.BoolValue
}

// Bool 创建一个布尔字段
func Bool(k string, v bool) *BoolField {
	return Boolp(k, &v)
}

// Boolp 创建一个布尔字段
func Boolp(k string, v *bool) *BoolField {
	return &BoolField{
		StringKey: key.String(k),
		BoolValue: value.Boolp(v),
	}
}
