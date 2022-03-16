package gox

import (
	`encoding/json`
	`fmt`
	`strings`
)

var _ = NewFieldsException

type (
	FieldsException interface {
		MessageException

		// Fields 返回错误实体
		// 在某些错误下，可能需要返回额外的信息给前端处理
		// 比如，认证错误，需要返回哪些字段有错误
		Fields() Fields
	}

	fieldsException struct {
		message string
		fields  Fields
	}
)

// NewFieldsException 创建错误
func NewFieldsException(message string, fields ...Field) *fieldsException {
	return &fieldsException{
		message: message,
		fields:  fields,
	}
}

func (fe *fieldsException) Message() string {
	return fe.message
}

func (fe *fieldsException) Fields() Fields {
	return fe.fields
}

func (fe fieldsException) MarshalJSON() (bytes []byte, err error) {
	output := make(map[string]interface{})
	output[`message`] = fe.message

	if 0 < len(fe.fields) {
		data := make(map[string]interface{})
		for _, field := range fe.fields {
			data[field.Key()] = field.Value()
		}
		output[`data`] = data
	}
	bytes, err = json.Marshal(output)

	return
}

func (fe *fieldsException) Error() string {
	var sb strings.Builder
	sb.WriteRune('{')
	sb.WriteString(fmt.Sprintf(`message = %s, `, fe.message))
	sb.WriteString(fmt.Sprintf(`data = %s`, fe.fields.String()))
	sb.WriteRune('}')

	return sb.String()
}
