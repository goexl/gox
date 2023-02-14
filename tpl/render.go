package tpl

type render struct {
	input     []string
	inputType inputType
	data      any
	executor  executor
}

func newRender(input []string, inputType inputType, data any, executor executor) *render {
	return &render{
		input:     input,
		inputType: inputType,
		data:      data,
		executor:  executor,
	}
}

func (r render) ToString() (string, error) {
	return r.executor.toString(r.input, r.inputType, r.data)
}

func (r render) ToFile(name string) error {
	return r.executor.toFile(r.input, r.inputType, r.data, name)
}
