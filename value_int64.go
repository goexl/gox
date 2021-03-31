package gox

var _ Value = (*Int64Value)(nil)

// Int64Value 整形值
type Int64Value struct {
	value int64
}

func (iv *Int64Value) Value() interface{} {
	return iv.value
}
