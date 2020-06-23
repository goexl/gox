package gox

import (
	`fmt`
)

type (
	// Error 接口，符合条件的错误统一处理
	Error interface {
		// ErrorCode 返回错误码
		ErrorCode() int
		// Message 返回错误消息
		Message() string
		// Data 返回错误实体
		// 在某些错误下，可能需要返回额外的信息给前端处理
		// 比如，认证错误，需要返回哪些字段有错误
		Data() interface{}
	}

	// ErrorCode 错误码
	ErrorCode int

	// CodeError 带错误编号和消息的错误
	CodeError struct {
		// 错误码
		errorCode ErrorCode
		// 消息
		message string
		// 数据
		data interface{}
	}
)

// NewCodeError 创建错误
func NewCodeError(errorCode ErrorCode, message string, data interface{}) *CodeError {
	return &CodeError{
		errorCode: errorCode,
		message:   message,
		data:      data,
	}
}

func (ce *CodeError) ErrorCode() int {
	return int(ce.errorCode)
}

func (ce *CodeError) Message() string {
	return ce.message
}

func (ce *CodeError) Data() interface{} {
	return ce.data
}

func (ce *CodeError) Error() string {
	return fmt.Sprintf("{ErrorCode=%d, message=%s, data=%v}", ce.errorCode, ce.message, ce.data)
}
