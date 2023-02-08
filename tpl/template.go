package tpl

type _template struct {
	input     []string
	inputType inputType
	data      any
	executor  executor
}

func newTemplate(input []string, inputType inputType, data any, executor executor) *_template {
	return &_template{
		input:     input,
		inputType: inputType,
		data:      data,
		executor:  executor,
	}
}

func (t _template) String() (string, error) {
	return t.executor.toString(t.input, t.inputType, t.data)
}

func (t _template) File(name string) error {
	return t.executor.toFile(t.input, t.inputType, t.data, name)
}
