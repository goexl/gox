package value

import (
	"fmt"

	"github.com/goexl/gox"
)

var (
	_           = Stringer
	_ gox.Value = (*StringerValue)(nil)
)

// StringerValue 布尔值
type StringerValue struct {
	value fmt.Stringer
}

// Stringer 布尔值
func Stringer(value fmt.Stringer) StringerValue {
	return StringerValue{
		value: value,
	}
}

func (sv *StringerValue) Value() interface{} {
	return sv.value
}

func (sv *StringerValue) String() string {
	return sv.value.String()
}
