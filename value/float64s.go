package value

import (
	"encoding/json"

	"github.com/goexl/gox"
)

var _ gox.Value = (*Float64sValue)(nil)

// Float64sValue 浮点数组值
type Float64sValue struct {
	value []float64
}

// Float64s 浮点数组值
func Float64s(value ...float64) Float64sValue {
	return Float64sValue{
		value: value,
	}
}

func (fv *Float64sValue) Value() interface{} {
	return fv.value
}

func (fv *Float64sValue) String() (str string) {
	if bytes, err := json.Marshal(fv.value); nil != err {
		str = `error`
	} else {
		str = string(bytes)
	}

	return
}
