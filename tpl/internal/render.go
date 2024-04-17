package internal

import (
	"github.com/goexl/gox/tpl/internal/core"
)

type Render struct {
	input     []string
	inputType core.InputType
	name      string
	data      any
	executor  core.Executor
}

func NewRender(input []string, inputType core.InputType, data any, executor core.Executor) *Render {
	return &Render{
		input:     input,
		inputType: inputType,
		data:      data,
		executor:  executor,
	}
}

func (r *Render) Template(name string) *Render {
	r.name = name

	return r
}

func (r *Render) ToString() (string, error) {
	return r.executor.ToString(r.input, r.inputType, r.name, r.data)
}

func (r *Render) ToFile(path string) error {
	return r.executor.ToFile(r.input, r.inputType, r.name, r.data, path)
}
