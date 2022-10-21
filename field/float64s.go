package field

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/key"
	"github.com/goexl/gox/value"
)

var (
	_         = Float64s
	_ gox.Key = (*Float64sField)(nil)
)

// Float64sField 描述一个浮点型数组字段
type Float64sField struct {
	key.StringKey
	value.Float64sValue
}

// Float64s 创建一个浮点型数组字段
func Float64s(k string, v ...float64) *Float64sField {
	return &Float64sField{
		StringKey:     key.String(k),
		Float64sValue: value.Float64s(v...),
	}
}
