package value

import (
	"fmt"

	"github.com/goexl/gox"
)

var (
	_           = Stringers
	_ gox.Value = (*StringersValue)(nil)
)

// StringersValue 字符串列表
type StringersValue struct {
	values []fmt.Stringer
}

// Stringers 布尔值
func Stringers(values ...fmt.Stringer) StringersValue {
	return StringersValue{
		values: values,
	}
}

func (sv *StringersValue) Value() interface{} {
	return sv.values
}

func (sv *StringersValue) String() string {
	return fmt.Sprintf(`%v`, sv.values)
}
