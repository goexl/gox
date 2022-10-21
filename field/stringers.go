package field

import (
	"fmt"

	"github.com/goexl/gox"
	"github.com/goexl/gox/key"
	"github.com/goexl/gox/value"
)

var (
	_         = Stringers
	_ gox.Key = (*StringersField)(nil)
)

// StringersField 字符串列表
type StringersField struct {
	key.StringKey
	value.StringersValue
}

// Stringers 创建一个字符串列表字段
func Stringers(k string, vs ...fmt.Stringer) *StringersField {
	return &StringersField{
		StringKey:      key.String(k),
		StringersValue: value.Stringers(vs...),
	}
}
