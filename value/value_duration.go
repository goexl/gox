package value

import (
	`fmt`
	`time`

	`github.com/goexl/gox`
)

var (
	_           = Duration
	_           = Durationp
	_ gox.Value = (*DurationValue)(nil)
)

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

func (dv *DurationValue) String() string {
	return fmt.Sprintf(`%v`, *dv.value)
}
