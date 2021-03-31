package gox

var _ Key = (*Int64Field)(nil)

// Int64Field 描述一个整形字段
type Int64Field struct {
	StringKey
	Int64Value
}

// NewInt64Field 创建一个整形字段
func NewInt64Field(key string, value int64) *Int64Field {
	return &Int64Field{
		StringKey: StringKey{
			key: key,
		},
		Int64Value: Int64Value{
			value: value,
		},
	}
}
