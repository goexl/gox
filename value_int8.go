package gox

var _ Value = (*Int8Value)(nil)

// Int8Value 整形值
type Int8Value struct {
	value int
}

func (iv *Int8Value) Value() interface{} {
	return iv.value
}
