package gox

var _ Key = (*StringField)(nil)

// StringField 描述一个字符串字段
type StringField struct {
	StringKey
	StringValue
}

// NewStringField 创建一个整形字段
func NewStringField(key string, value string) *StringField {
	return &StringField{
		StringKey: StringKey{
			key: key,
		},
		StringValue: StringValue{
			value: value,
		},
	}
}
