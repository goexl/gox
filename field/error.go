package field

var _ = Error

// Error 创建错误字段
func Error(err error) *_any[error] {
	return New("error", err)
}
