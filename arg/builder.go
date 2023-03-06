package arg

type builder struct {
	params *params
}

func newBuilder() *builder {
	return &builder{
		params: newParams(),
	}
}

func (b *builder) Size(size int) *builder {
	b.params.size = size

	return b
}

func (b *builder) Short(placeholder string) *builder {
	b.params.short = placeholder

	return b
}

func (b *builder) Long(placeholder string) *builder {
	b.params.long = placeholder

	return b
}

func (b *builder) Equal(placeholder string) *builder {
	b.params.equal = placeholder

	return b
}

func (b *builder) Build() *argsBuilder {
	return newArgsBuilder(b.params)
}
