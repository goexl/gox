package gox

var _ Key = (*IntField)(nil)

// IntField 描述一个整形字段
type IntField struct {
	StringKey
	IntValue
}

// NewIntField 创建一个整形字段
func NewIntField(key string, value int) *IntField {
	return &IntField{
		StringKey: StringKey{
			key: key,
		},
		IntValue: IntValue{
			value: value,
		},
	}
}
