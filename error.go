package gox

import (
	`encoding/json`
)

type (
	// ToErrorCode 错误码
	ErrorCode int

	// Error 接口，符合条件的错误统一处理
	Error interface {
		// ToErrorCode 返回错误码
		ToErrorCode() ErrorCode
		// ToMessage 返回错误消息
		ToMessage() string
		// ToData 返回错误实体
		// 在某些错误下，可能需要返回额外的信息给前端处理
		// 比如，认证错误，需要返回哪些字段有错误
		ToData() interface{}
	}

	// CodeError 带错误编号和消息的错误
	CodeError struct {
		// ErrorCode 错误码
		ErrorCode ErrorCode `json:"errorCode"`
		// Message 消息
		Message string `json:"message"`
		// Data 数据
		Data interface{} `json:"data"`
	}
)

// NewCodeError 创建错误
func NewCodeError(errorCode ErrorCode, message string, data interface{}) *CodeError {
	return &CodeError{
		ErrorCode: errorCode,
		Message:   message,
		Data:      data,
	}
}

func ParseCodeError(str string) (ec *CodeError, err error) {
	err = json.Unmarshal([]byte(str), &ec)

	return
}

func (ce *CodeError) ToErrorCode() ErrorCode {
	return ce.ErrorCode
}

func (ce *CodeError) ToMessage() string {
	return ce.Message
}

func (ce *CodeError) ToData() interface{} {
	return ce.Data
}

func (ce *CodeError) Error() (str string) {
	if data, err := json.Marshal(ce); nil != err {
		return
	} else {
		str = string(data)
	}

	return
}
