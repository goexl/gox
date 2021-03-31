package gox

var _ Key = (*Int8Field)(nil)

// Int8Field 描述一个整形字段
type Int8Field struct {
	StringKey
	Int8Value
}

// NewInt8Field 创建一个整形字段
func NewInt8Field(key string, value int8) *Int8Field {
	return &Int8Field{
		StringKey: StringKey{
			key: key,
		},
		Int8Value: Int8Value{
			value: value,
		},
	}
}
