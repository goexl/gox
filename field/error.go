package field

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/key"
	"github.com/goexl/gox/value"
)

var (
	_         = Error
	_ gox.Key = (*ErrorField)(nil)
)

// ErrorField 描述一个整形字段
type ErrorField struct {
	key.StringKey
	value.ErrorValue
}

// Error 创建一个整形字段
func Error(err error) *ErrorField {
	return &ErrorField{
		StringKey:  key.String("error"),
		ErrorValue: value.Error(err),
	}
}
