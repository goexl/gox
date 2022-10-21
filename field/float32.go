package field

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/key"
	"github.com/goexl/gox/value"
)

var (
	_         = Float32
	_ gox.Key = (*Float32Field)(nil)
)

// Float32Field 描述一个整形字段
type Float32Field struct {
	key.StringKey
	value.Float32Value
}

// Float32 创建一个整形字段
func Float32(k string, v float32) *Float32Field {
	return Float32p(k, &v)
}

// Float32p 创建一个整形字段
func Float32p(k string, v *float32) *Float32Field {
	return &Float32Field{
		StringKey:    key.String(k),
		Float32Value: value.Float32p(v),
	}
}
