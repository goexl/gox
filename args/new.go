package args

var _ = New

// New 创建新的参数
func New() *creator {
	return newCreator()
}
