package gox

import (
	`encoding/json`
)

type codeError struct {
	// 错误码
	code int
	// 消息
	message string
	// 数据
	fields Fields
}

// NewCodeError 创建错误
func NewCodeError(code int, message string, fields ...Field) *codeError {
	return &codeError{
		code:    code,
		message: message,
		fields:  fields,
	}
}

func (ce *codeError) Code() int {
	return ce.code
}

func (ce *codeError) Message() string {
	return ce.message
}

func (ce *codeError) Fields() Fields {
	return ce.fields
}

func (ce *codeError) Error() (error string) {
	output := make(map[string]interface{})
	output[`code`] = ce.code
	output[`message`] = ce.message

	if 0 < len(ce.fields) {
		data := make(map[string]interface{})
		for _, field := range ce.fields {
			data[field.Key()] = field.Value()
		}
		output[`data`] = data
	}

	if bytes, err := json.Marshal(ce); nil != err {
		error = err.Error()
	} else {
		error = string(bytes)
	}

	return
}
