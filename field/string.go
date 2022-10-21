package field

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/key"
	"github.com/goexl/gox/value"
)

var (
	_         = String
	_ gox.Key = (*StringField)(nil)
)

// StringField 描述一个字符串字段
type StringField struct {
	key.StringKey
	value.StringValue
}

// String 创建一个整形字段
func String(k string, v string) *StringField {
	return Stringp(k, &v)
}

// Stringp 创建一个整形字段
func Stringp(k string, v *string) *StringField {
	return &StringField{
		StringKey:   key.String(k),
		StringValue: value.Stringp(v),
	}
}
