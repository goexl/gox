package value

import (
	`time`

	`github.com/storezhang/gox`
)

var _ gox.Value = (*DurationValue)(nil)

// DurationValue 时间段值
type DurationValue struct {
	value *time.Duration
}

// Duration 时间段值
func Duration(value time.Duration) DurationValue {
	return Durationp(&value)
}

// Durationp 时间段值
func Durationp(value *time.Duration) DurationValue {
	return DurationValue{
		value: value,
	}
}

func (dv *DurationValue) Value() interface{} {
	return dv.value
}
