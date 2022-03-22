package value

import (
	`encoding/json`

	`github.com/goexl/gox`
)

var _ gox.Value = (*ErrorValue)(nil)

// ErrorValue 错误值
type ErrorValue struct {
	value error
}

// Error 错误值
func Error(err error) ErrorValue {
	return ErrorValue{
		value: err,
	}
}

func (ev *ErrorValue) Value() interface{} {
	return ev.value
}

func (ev *ErrorValue) String() (str string) {
	if bytes, err := json.Marshal(ev.value); nil != err {
		str = `error`
	} else {
		str = string(bytes)
	}

	return
}
