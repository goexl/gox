package value

import (
	`encoding/json`
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

func (sv *TimesValue) Value() interface{} {
	return sv.value
}

func (sv *TimesValue) String() (str string) {
	if bytes, err := json.Marshal(sv.value); nil != err {
		str = `error`
	} else {
		str = string(bytes)
	}

	return
}
