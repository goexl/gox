package rand

var _ = New

type builder struct {
	_string *_string
	length  int
	bytes   string
}

// New 创建随机字符串
func New() *builder {
	return &builder{
		_string: newString(),
		length:  8,
		bytes:   letters,
	}
}

func (b *builder) Length(length int) (builder *builder) {
	b.length = length

	return
}

func (b *builder) Code() (builder *builder) {
	b.length = 6
	b.bytes = numbers

	return
}

func (b *builder) Digital() (builder *builder) {
	b.bytes = numbers

	return
}

func (b *builder) Build() (value string) {
	return b._string.rand(b.length, b.bytes)
}
