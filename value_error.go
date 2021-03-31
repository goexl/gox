package gox

var _ Value = (*ErrorValue)(nil)

// ErrorValue 整形值
type ErrorValue struct {
	value error
}

func (iv *ErrorValue) Value() interface{} {
	return iv.value
}
