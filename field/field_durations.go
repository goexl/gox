package field

import (
	`time`

	`github.com/goexl/gox`
	`github.com/goexl/gox/key`
	`github.com/goexl/gox/value`
)

var _ gox.Key = (*DurationsField)(nil)

// DurationsField 描述一个时间字段数组字段
type DurationsField struct {
	key.StringKey
	value.DurationsValue
}

// Durations 创建一个时间字段数组字段
func Durations(k string, v ...time.Duration) *DurationsField {
	return &DurationsField{
		StringKey:      key.String(k),
		DurationsValue: value.Durations(v...),
	}
}
