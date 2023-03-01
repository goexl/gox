package http

var _ = New

type builder struct{}

// New 创建操作主对象
func New() *builder {
	return new(builder)
}

func (b *builder) Disposition(filename string) *dispositionBuilder {
	return newDispositionBuilder(filename)
}
