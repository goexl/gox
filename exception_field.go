package gox

import (
	`encoding/json`
	`fmt`
	`strings`
)

var _ = NewFieldException

type (
	// FieldException 带字段的异常
	FieldException interface {
		MessageException

		// Field 返回错误实体
		// 在某些错误下，可能需要返回额外的信息给前端处理
		// 比如，认证错误，需要返回哪些字段有错误
		Field() Field
	}

	fieldException struct {
		message string
		field   Field
	}
)

// NewFieldException 创建带字段的异常
func NewFieldException(message string, field Field) *fieldException {
	return &fieldException{
		message: message,
		field:   field,
	}
}

func (fe *fieldException) Message() string {
	return fe.message
}

func (fe *fieldException) Field() Field {
	return fe.field
}

func (fe fieldException) MarshalJSON() (bytes []byte, err error) {
	output := make(map[string]interface{})
	output[`message`] = fe.message
	output[fe.Field().Key()] = fe.Field().Value()
	bytes, err = json.Marshal(output)

	return
}

func (fe *fieldException) Error() string {
	var sb strings.Builder
	sb.WriteRune('{')
	sb.WriteString(fmt.Sprintf(`message = %s, `, fe.message))
	sb.WriteString(fmt.Sprintf(`%s = %v`, fe.field.Key(), fe.field.Value()))
	sb.WriteRune('}')

	return sb.String()
}
