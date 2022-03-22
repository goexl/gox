package value

import (
	`encoding/json`

	`github.com/goexl/gox`
)

var _ gox.Value = (*IntsValue)(nil)

// IntsValue 整形值
type IntsValue struct {
	value []int
}

// Ints 整形数组
func Ints(value ...int) IntsValue {
	return IntsValue{
		value: value,
	}
}

func (iv *IntsValue) Value() interface{} {
	return iv.value
}

func (iv *IntsValue) String() (str string) {
	if bytes, err := json.Marshal(iv.value); nil != err {
		str = `error`
	} else {
		str = string(bytes)
	}

	return
}
