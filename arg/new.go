package arg

var _ = New

// New 创建新的参数
func New() *builder {
	return newBuilder(size)
}
