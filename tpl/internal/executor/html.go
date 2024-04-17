package executor

import (
	"bytes"
	"html/template"
	"io"

	"github.com/goexl/gox/rand"
	"github.com/goexl/gox/tpl/internal/core"
)

var _ core.Executor = (*Html)(nil)

type Html struct {
	*core.Base

	template *template.Template
}

func NewHtml() *Html {
	return new(Html)
}

func (h *Html) ToFile(input []string, it core.InputType, name string, data any, filepath string) (err error) {
	if err = h.init(input, it); nil == err {
		err = h.RenderToFile(filepath, h.execute(name, data))
	}

	return
}

func (h *Html) ToString(input []string, it core.InputType, name string, data any) (result string, err error) {
	if err = h.init(input, it); nil != err {
		return
	}

	buffer := new(bytes.Buffer)
	if "" == name {
		err = h.template.Execute(buffer, data)
	} else {
		err = h.template.ExecuteTemplate(buffer, name, data)
	}
	if nil == err {
		result = buffer.String()
	}

	return
}

func (h *Html) init(input []string, it core.InputType) (err error) {
	switch it {
	case core.InputTypeFile:
		h.template, err = template.ParseFiles(input...)
	case core.InputTypeString:
		fallthrough
	default:
		h.template, err = template.New(rand.New().String().Build().Generate()).Parse(input[0])
	}

	return
}

func (h *Html) execute(name string, data any) core.Function {
	return func(wr io.Writer) (err error) {
		if "" == name {
			err = h.template.Execute(wr, data)
		} else {
			err = h.template.ExecuteTemplate(wr, name, data)
		}

		return
	}
}
