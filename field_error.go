package gox

var _ Key = (*ErrorField)(nil)

// ErrorField 描述一个整形字段
type ErrorField struct {
	StringKey
	ErrorValue
}

// NewErrorField 创建一个整形字段
func NewErrorField(value error) *ErrorField {
	return &ErrorField{
		StringKey: StringKey{
			key: "error",
		},
		ErrorValue: ErrorValue{
			value: value,
		},
	}
}
