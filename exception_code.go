package gox

import (
	`fmt`
)

var _ = NewCodeException

// CodeException 带状态码的错误
type (
	CodeException interface {
		// Code 返回错误码
		Code() int
	}

	codeException struct {
		code int
	}
)

// NewCodeException 创建异常
func NewCodeException(code int) *codeException {
	return &codeException{
		code: code,
	}
}

func (ce *codeException) Code() int {
	return ce.code
}

func (ce codeException) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{code: %d}`, ce.code)), nil
}

func (ce *codeException) Error() string {
	return fmt.Sprintf(`code = %d, `, ce.code)
}
