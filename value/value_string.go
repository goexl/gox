package value

import (
	`fmt`

	`github.com/goexl/gox`
)

var (
	_           = String
	_           = Stringp
	_ gox.Value = (*StringValue)(nil)
)

// StringValue 字符串值
type StringValue struct {
	value *string
}

// String 字符串值
func String(value string) StringValue {
	return Stringp(&value)
}

// Stringp 字符串值
func Stringp(value *string) StringValue {
	return StringValue{
		value: value,
	}
}

func (sv *StringValue) Value() interface{} {
	return sv.value
}

func (sv *StringValue) String() string {
	return fmt.Sprintf(`%v`, *sv.value)
}
