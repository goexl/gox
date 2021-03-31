package gox

var _ Value = (*IntValue)(nil)

// IntValue 整形值
type IntValue struct {
	value int
}

func (iv *IntValue) Value() interface{} {
	return iv.value
}
