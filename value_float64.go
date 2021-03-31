package gox

var _ Value = (*Float64Value)(nil)

// Float64Value 浮点值
type Float64Value struct {
	value float64
}

func (iv *Float64Value) Value() interface{} {
	return iv.value
}
