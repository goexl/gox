package check

var _ = New

type builder struct {
	typ checkerType
}

// New 创建新的检查器
func New() *builder {
	return &builder{
		typ: hasTypeAny,
	}
}

func (b *builder) Any() *builder {
	b.typ = hasTypeAny

	return b
}

func (b *builder) All() *builder {
	b.typ = hasTypeAll

	return b
}

func (b *builder) String(check string) *_string {
	return newString(b.typ, check)
}
