package rand

var _ = New

type generator struct {}

// New 创建随机字符串
func New() *generator {
	return new(generator)
}

func (g *generator) String() *stringBuilder {
	return newStringBuilder()
}

func (g *generator) Duration() *durationBuilder {
	return newDurationBuilder()
}

func (g *generator) Time() *durationBuilder {
	return newDurationBuilder()
}
