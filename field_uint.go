package gox

var _ Key = (*UintField)(nil)

// UintField 描述一个整形字段
type UintField struct {
	StringKey
	UintValue
}

// NewUintField 创建一个整形字段
func NewUintField(key string, value uint) *UintField {
	return &UintField{
		StringKey: StringKey{
			key: key,
		},
		UintValue: UintValue{
			value: value,
		},
	}
}
