package gox

import (
	`encoding/json`
)

type (
	// ErrorCode 错误码
	ErrorCode int

	// Error 接口，符合条件的错误统一处理
	Error interface {
		// ErrorCode 返回错误码
		ErrorCode() ErrorCode
		// Message 返回错误消息
		Message() string
		// Fields 返回错误实体
		// 在某些错误下，可能需要返回额外的信息给前端处理
		// 比如，认证错误，需要返回哪些字段有错误
		Fields() Fields
	}

	codeError struct {
		// 错误码
		Code ErrorCode `json:"code"`
		// 消息
		Message string `json:"message"`
		// 数据
		Fields Fields `json:"data"`
	}
)

// NewCodeError 创建错误
func NewCodeError(errorCode ErrorCode, message string, fields Fields) *codeError {
	return &codeError{
		Code:    errorCode,
		Message: message,
		Data:    data,
	}
}

// ParseCodeError 从JSON字符串中解析错误
func ParseCodeError(str string) (ec *codeError, err error) {
	err = json.Unmarshal([]byte(str), &ec)

	return
}

func (ce *codeError) ToErrorCode() ErrorCode {
	return ce.ErrorCode
}

func (ce *codeError) ToMessage() string {
	return ce.Message
}

func (ce *codeError) ToData() interface{} {
	return ce.Data
}

func (ce *codeError) Error() (str string) {
	if data, err := json.Marshal(ce); nil != err {
		return
	} else {
		str = string(data)
	}

	return
}
