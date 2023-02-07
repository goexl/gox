package tpl

type builder struct {
	data      any
	input     []string
	inputType inputType
}

func newBuilder(input string) *builder {
	return &builder{
		input:     []string{input},
		inputType: inputTypeString,
	}
}

func (b *builder) Data(data any) *builder {
	b.data = data

	return b
}

func (b *builder) File(paths ...string) *builder {
	b.inputType = inputTypeFile
	b.input = append(b.input, paths...)

	return b
}

func (b *builder) String() *builder {
	b.inputType = inputTypeString

	return b
}
