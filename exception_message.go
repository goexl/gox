package gox

import (
	`fmt`
)

var _ = NewMessageException

type (
	// MessageException 带消息的异常
	MessageException interface {
		// Message 返回错误消息
		Message() string
	}

	messageException struct {
		message string
	}
)

// NewMessageException 创建带消息的异常
func NewMessageException(message string) *messageException {
	return &messageException{
		message: message,
	}
}

func (mex *messageException) Message() string {
	return mex.message
}

func (mex messageException) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{message: %s}`, mex.message)), nil
}

func (mex *messageException) Error() string {
	return fmt.Sprintf(`message = %s, `, mex.message)
}
