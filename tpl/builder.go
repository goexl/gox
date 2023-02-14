package tpl

type builder struct {
	data      any
	input     []string
	inputType inputType
	executor  executor
}

func newBuilder(input string) *builder {
	return &builder{
		input:     []string{input},
		inputType: inputTypeString,
		executor:  newTextExecutor(),
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

func (b *builder) Text() *builder {
	b.executor = newTextExecutor()

	return b
}

func (b *builder) Html() *builder {
	b.executor = newHtmlExecutor()

	return b
}

func (b *builder) Build() *render {
	return newRender(b.input, b.inputType, b.data, b.executor)
}
