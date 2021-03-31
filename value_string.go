package gox

var _ Value = (*StringValue)(nil)

// StringValue 字符串值
type StringValue struct {
	value string
}

func (iv *StringValue) Value() interface{} {
	return iv.value
}
