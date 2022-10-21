package field

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/key"
	"github.com/goexl/gox/value"
)

var (
	_         = Bools
	_ gox.Key = (*BoolsField)(nil)
)

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
