package value

import (
	`encoding/json`

	`github.com/goexl/gox`
)

var _ gox.Value = (*UintsValue)(nil)

// UintsValue 整形数组值
type UintsValue struct {
	value []uint
}

// Uints 整形数组值
func Uints(value ...uint) UintsValue {
	return UintsValue{
		value: value,
	}
}

func (iv *UintsValue) Value() interface{} {
	return iv.value
}

func (iv *UintsValue) String() (str string) {
	if bytes, err := json.Marshal(iv.value); nil != err {
		str = `error`
	} else {
		str = string(bytes)
	}

	return
}
