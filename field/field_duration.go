package field

import (
	`time`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*DurationField)(nil)

// DurationField 描述一个时间段字段
type DurationField struct {
	key.StringKey
	value.DurationValue
}

// Duration 创建一个时间段字段
func Duration(k string, v time.Duration) *DurationField {
	return Durationp(k, &v)
}

// Durationp 创建一个时间段字段
func Durationp(k string, v *time.Duration) *DurationField {
	return &DurationField{
		StringKey:     key.String(k),
		DurationValue: value.Durationp(v),
	}
}
