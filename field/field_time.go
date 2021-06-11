package field

import (
	`time`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*TimeField)(nil)

// TimeField 描述一个时间字段
type TimeField struct {
	key.StringKey
	value.TimeValue
}

// Time 创建一个时间字段
func Time(k string, v time.Time) *TimeField {
	return Timep(k, &v)
}

// Timep 创建一个时间字段
func Timep(k string, v *time.Time) *TimeField {
	return &TimeField{
		StringKey: key.String(k),
		TimeValue: value.Timep(v),
	}
}
