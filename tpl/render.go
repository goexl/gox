package tpl

type render struct {
	input     []string
	inputType inputType
	name      string
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

func (r *render) Template(name string) *render {
	r.name = name

	return r
}

func (r *render) ToString() (string, error) {
	return r.executor.toString(r.input, r.inputType, r.name, r.data)
}

func (r *render) ToFile(path string) error {
	return r.executor.toFile(r.input, r.inputType, r.name, r.data, path)
}
