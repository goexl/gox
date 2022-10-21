package value

import (
	"encoding/json"

	"github.com/goexl/gox"
)

var _ gox.Value = (*AnyValue)(nil)

// AnyValue 任意值
type AnyValue struct {
	value interface{}
}

// Any 任意值
func Any(value interface{}) AnyValue {
	return AnyValue{
		value: value,
	}
}

func (av *AnyValue) Value() interface{} {
	return av.value
}

func (av *AnyValue) String() (str string) {
	if bytes, err := json.Marshal(av.value); nil != err {
		str = `error`
	} else {
		str = string(bytes)
	}

	return
}
