package field

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/key`
	`github.com/storezhang/gox/value`
)

var _ gox.Key = (*ErrorField)(nil)

// ErrorField 描述一个整形字段
type ErrorField struct {
	key.StringKey
	value.ErrorValue
}

// NewErrorField 创建一个整形字段
func NewErrorField(err error) *ErrorField {
	return &ErrorField{
		StringKey:  key.NewStringKey("error"),
		ErrorValue: value.NewErrorValue(err),
	}
}
