package value

import (
	`time`

	`github.com/storezhang/gox`
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
