package gox

import (
	`encoding/json`
	`fmt`
	`strings`
)

var _ = NewException

type (
	// Exception 接口，符合条件的错误统一处理
	Exception interface {
		CodeException
		MessageException
		FieldsException
	}

	exception struct {
		code int
		message string
		fields Fields
	}
)

// NewException 创建异常
func NewException(code int, message string, fields ...Field) *exception {
	return &exception{
		code:    code,
		message: message,
		fields:  fields,
	}
}

func (e *exception) Code() int {
	return e.code
}

func (e *exception) Message() string {
	return e.message
}

func (e *exception) Fields() Fields {
	return e.fields
}

func (e exception) MarshalJSON() (bytes []byte, err error) {
	output := make(map[string]interface{})
	output[`code`] = e.code
	output[`message`] = e.message

	if 0 < len(e.fields) {
		data := make(map[string]interface{})
		for _, field := range e.fields {
			data[field.Key()] = field.Value()
		}
		output[`data`] = data
	}
	bytes, err = json.Marshal(output)

	return
}

func (e *exception) Error() string {
	var sb strings.Builder
	sb.WriteRune('{')
	sb.WriteString(fmt.Sprintf(`code = %d, `, e.code))
	sb.WriteString(fmt.Sprintf(`message = %s, `, e.message))
	sb.WriteString(fmt.Sprintf(`data = %s`, e.fields.String()))
	sb.WriteRune('}')

	return sb.String()
}
