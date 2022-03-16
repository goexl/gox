package value

import (
	`time`

	`github.com/goexl/gox`
)

var _ gox.Value = (*TimeValue)(nil)

// TimeValue 时间值
type TimeValue struct {
	value *time.Time
}

// Time 时间值
func Time(value time.Time) TimeValue {
	return Timep(&value)
}

// Timep 时间值
func Timep(value *time.Time) TimeValue {
	return TimeValue{
		value: value,
	}
}

func (iv *TimeValue) Value() interface{} {
	return iv.value
}
