package field

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/key"
	"github.com/goexl/gox/value"
)

var (
	_         = Int8
	_ gox.Key = (*Int8Field)(nil)
)

// Int8Field 描述一个整形字段
type Int8Field struct {
	key.StringKey
	value.Int8Value
}

// Int8 创建一个整形字段
func Int8(k string, v int8) *Int8Field {
	return Int8p(k, &v)
}

// Int8p 创建一个整形字段
func Int8p(k string, v *int8) *Int8Field {
	return &Int8Field{
		StringKey: key.String(k),
		Int8Value: value.Int8p(v),
	}
}
