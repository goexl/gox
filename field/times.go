package field

import (
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/key"
	"github.com/goexl/gox/value"
)

var (
	_         = Times
	_ gox.Key = (*TimesField)(nil)
)

// TimesField 描述一个时间数组字段
type TimesField struct {
	key.StringKey
	value.TimesValue
}

// Times 创建一个时间数组字段
func Times(k string, v ...time.Time) *TimesField {
	return &TimesField{
		StringKey:  key.String(k),
		TimesValue: value.Times(v...),
	}
}
