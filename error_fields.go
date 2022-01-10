package gox

import (
	`encoding/json`
)

type fieldsError struct {
	// 消息
	message string
	// 数据
	fields Fields
}

// NewFieldsError 创建错误
func NewFieldsError(message string, fields ...Field) *codeError {
	return &codeError{
		message: message,
		fields:  fields,
	}
}

func (fe *fieldsError) Code() int {
	return -1
}

func (fe *fieldsError) Message() string {
	return fe.message
}

func (fe *fieldsError) Fields() Fields {
	return fe.fields
}

func (fe *fieldsError) Error() (error string) {
	output := make(map[string]interface{})
	output[`message`] = fe.message

	if 0 < len(fe.fields) {
		data := make(map[string]interface{})
		for _, field := range fe.fields {
			data[field.Key()] = field.Value()
		}
		output[`data`] = data
	}

	if bytes, err := json.Marshal(fe); nil != err {
		error = err.Error()
	} else {
		error = string(bytes)
	}

	return
}
