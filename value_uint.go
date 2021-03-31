package gox

var _ Value = (*UintValue)(nil)

// UintValue 整形值
type UintValue struct {
	value uint
}

func (iv *UintValue) Value() interface{} {
	return iv.value
}
