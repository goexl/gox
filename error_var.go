package gox

var (
	// 文件相关
	ErrOnCopySelf   = &CodeError{ErrorCode: 1001, Message: "不能复制或者移动自身"}
	ErrSrcIsNotFile = &CodeError{ErrorCode: 1002, Message: "源文件不是文件"}

	// 密钥相关
	ErrKeySize = &CodeError{ErrorCode: 2001, Message: "密钥位数有错，必须是8的倍数"}
)
