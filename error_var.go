package gox

var (
	// ErrSrcIsNotFile 文件相关
	ErrSrcIsNotFile = &codeError{ErrorCode: 1002, Message: "源文件不是文件"}
	// ErrKeySize 密钥相关
	ErrKeySize = &codeError{ErrorCode: 2001, Message: "密钥位数有错，必须是8的倍数"}
)
