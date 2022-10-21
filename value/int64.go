package value

import (
	"fmt"

	"github.com/goexl/gox"
)

var (
	_           = Int64
	_           = Int64p
	_ gox.Value = (*Int64Value)(nil)
)

// Int64Value 整形值
type Int64Value struct {
	value *int64
}

// Int64 整形值
func Int64(value int64) Int64Value {
	return Int64p(&value)
}

// Int64p 整形值
func Int64p(value *int64) Int64Value {
	return Int64Value{
		value: value,
	}
}

func (iv *Int64Value) Value() interface{} {
	return iv.value
}

func (iv *Int64Value) String() string {
	return fmt.Sprintf(`%v`, *iv.value)
}
