package gox

var _ Value = (*Float32Value)(nil)

// Float32Value 浮点值
type Float32Value struct {
	value float32
}

func (iv *Float32Value) Value() interface{} {
	return iv.value
}
