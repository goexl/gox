package value

import (
	`encoding/json`

	`github.com/goexl/gox`
)

var _ gox.Value = (*Int64sValue)(nil)

// Int64sValue 整形数组值
type Int64sValue struct {
	value []int64
}

// Int64s 整形数组值
func Int64s(value ...int64) Int64sValue {
	return Int64sValue{
		value: value,
	}
}

func (iv *Int64sValue) Value() interface{} {
	return iv.value
}

func (iv *Int64sValue) String() (str string) {
	if bytes, err := json.Marshal(iv.value); nil != err {
		str = `error`
	} else {
		str = string(bytes)
	}

	return
}
