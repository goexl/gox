package field

import (
	"fmt"

	"github.com/goexl/gox"
	"github.com/goexl/gox/key"
	"github.com/goexl/gox/value"
)

var (
	_         = Stringer
	_ gox.Key = (*StringerField)(nil)
)

// StringerField 描述一个布尔字段
type StringerField struct {
	key.StringKey
	value.StringerValue
}

// Stringer 创建一个字符串字段
func Stringer(k string, v fmt.Stringer) *StringerField {
	return &StringerField{
		StringKey:     key.String(k),
		StringerValue: value.Stringer(v),
	}
}
