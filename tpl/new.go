package tpl

var _ = New

// New 创建模板
func New(input string) *builder {
	return newBuilder(input)
}
