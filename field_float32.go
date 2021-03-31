package gox

var _ Key = (*Float32Field)(nil)

// Float32Field 描述一个整形字段
type Float32Field struct {
	StringKey
	Float32Value
}

// NewFloat32Field 创建一个整形字段
func NewFloat32Field(key string, value float32) *Float32Field {
	return &Float32Field{
		StringKey: StringKey{
			key: key,
		},
		Float32Value: Float32Value{
			value: value,
		},
	}
}
