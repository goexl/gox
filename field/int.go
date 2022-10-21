package field

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/key"
	"github.com/goexl/gox/value"
)

var (
	_         = Int
	_ gox.Key = (*IntField)(nil)
)

// IntField 描述一个整形字段
type IntField struct {
	key.StringKey
	value.IntValue
}

// Int 创建一个整形字段
func Int(k string, v int) *IntField {
	return Intp(k, &v)
}

// Intp 创建一个整形字段
func Intp(k string, v *int) *IntField {
	return &IntField{
		StringKey: key.String(k),
		IntValue:  value.Intp(v),
	}
}
