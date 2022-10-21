package value

import (
	`encoding/json`

	`github.com/goexl/gox`
)

var _ gox.Value = (*StringsValue)(nil)

// StringsValue 字符串值
type StringsValue struct {
	value []string
}

// Strings 字符串数组
func Strings(value ...string) StringsValue {
	return StringsValue{
		value: value,
	}
}

func (sv *StringsValue) Value() interface{} {
	return sv.value
}

func (sv *StringsValue) String() (str string) {
	if bytes, err := json.Marshal(sv.value); nil != err {
		str = `error`
	} else {
		str = string(bytes)
	}

	return
}
