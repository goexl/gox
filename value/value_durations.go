package value

import (
	`encoding/json`
	`time`

	`github.com/goexl/gox`
)

var _ gox.Value = (*DurationsValue)(nil)

// DurationsValue 时间段数组值
type DurationsValue struct {
	value []time.Duration
}

// Durations 时间段数组值
func Durations(value ...time.Duration) DurationsValue {
	return DurationsValue{
		value: value,
	}
}

func (dv *DurationsValue) Value() interface{} {
	return dv.value
}

func (dv *DurationsValue) String() (str string) {
	if bytes, err := json.Marshal(dv.value); nil != err {
		str = `error`
	} else {
		str = string(bytes)
	}

	return
}
