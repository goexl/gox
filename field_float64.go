package gox

var _ Key = (*Float64Field)(nil)

// Float64Field 描述一个整形字段
type Float64Field struct {
	StringKey
	Float64Value
}

// NewFloat64Field 创建一个整形字段
func NewFloat64Field(key string, value float64) *Float64Field {
	return &Float64Field{
		StringKey: StringKey{
			key: key,
		},
		Float64Value: Float64Value{
			value: value,
		},
	}
}
