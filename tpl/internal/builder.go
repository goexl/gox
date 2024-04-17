package internal

import (
	"github.com/goexl/gox/tpl/internal/core"
	"github.com/goexl/gox/tpl/internal/executor"
)

type Builder struct {
	data      any
	input     []string
	inputType core.InputType
	executor  core.Executor
}

func NewBuilder(input string) *Builder {
	return &Builder{
		input:     []string{input},
		inputType: core.InputTypeString,
		executor:  executor.NewText(),
	}
}

func (b *Builder) Data(data any) *Builder {
	b.data = data

	return b
}

func (b *Builder) File(paths ...string) *Builder {
	b.inputType = core.InputTypeFile
	b.input = append(b.input, paths...)

	return b
}

func (b *Builder) String() *Builder {
	b.inputType = core.InputTypeString

	return b
}

func (b *Builder) Text() *Builder {
	b.executor = executor.NewText()

	return b
}

func (b *Builder) Html() *Builder {
	b.executor = executor.NewHtml()

	return b
}

func (b *Builder) Build() *Render {
	return NewRender(b.input, b.inputType, b.data, b.executor)
}
