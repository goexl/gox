package rand

var _ = New

type generator struct {
	_string *_string
	length  int
	bytes   string
}

// New 创建随机字符串
func New() *generator {
	return &generator{
		_string: newString(),
		length:  8,
		bytes:   letters,
	}
}

func (g *generator) Length(length int) *generator {
	g.length = length

	return g
}

func (g *generator) Code() *generator {
	g.length = 6
	g.bytes = numbers

	return g
}

func (g *generator) Digital() *generator {
	g.bytes = numbers

	return g
}

func (g *generator) Generate() (value string) {
	return g._string.rand(g.length, g.bytes)
}
