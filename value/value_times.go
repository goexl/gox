package value

import (
	`time`

	`github.com/goexl/gox`
)

var _ gox.Value = (*TimesValue)(nil)

// TimesValue 时间数组值
type TimesValue struct {
	value []time.Time
}

// Times 时间数组值
func Times(value ...time.Time) TimesValue {
	return TimesValue{
		value: value,
	}
}

func (iv *TimesValue) Value() interface{} {
	return iv.value
}
