package field

import (
	`time`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*TimesField)(nil)

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
