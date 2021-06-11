package field

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*StringsField)(nil)

// StringsField 描述一个字符串字段
type StringsField struct {
	key.StringKey
	value.StringsValue
}

// Strings 创建一个整形字段
func Strings(k string, v ...string) *StringsField {
	return &StringsField{
		StringKey:    key.String(k),
		StringsValue: value.Strings(v...),
	}
}
