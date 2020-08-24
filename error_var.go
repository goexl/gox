package gox

var (
	ErrOnCopySelf   = &CodeError{ErrorCode: 1001, Msg: "不能复制或者移动自身"}
	ErrSrcIsNotFile = &CodeError{ErrorCode: 1002, Msg: "源文件不是文件"}
)
