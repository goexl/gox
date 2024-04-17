package executor

import (
	"bytes"
	"io"
	"text/template"

	"github.com/goexl/gox/rand"
	"github.com/goexl/gox/tpl/internal/core"
)

var _ core.Executor = (*Text)(nil)

type Text struct {
	*core.Base

	template *template.Template
}

func NewText() *Text {
	return new(Text)
}

func (t *Text) ToFile(input []string, it core.InputType, name string, data any, filepath string) (err error) {
	if err = t.init(input, it); nil == err {
		err = t.RenderToFile(filepath, t.execute(name, data))
	}

	return
}

func (t *Text) ToString(input []string, it core.InputType, name string, data any) (result string, err error) {
	if err = t.init(input, it); nil != err {
		return
	}

	buffer := new(bytes.Buffer)
	if "" == name {
		err = t.template.Execute(buffer, data)
	} else {
		err = t.template.ExecuteTemplate(buffer, name, data)
	}
	if nil == err {
		result = buffer.String()
	}

	return
}

func (t *Text) init(input []string, it core.InputType) (err error) {
	switch it {
	case core.InputTypeFile:
		t.template, err = template.ParseFiles(input...)
	case core.InputTypeString:
		fallthrough
	default:
		t.template, err = template.New(rand.New().String().Build().Generate()).Parse(input[0])
	}

	return
}

func (t *Text) execute(name string, data any) core.Function {
	return func(wr io.Writer) (err error) {
		if "" == name {
			err = t.template.Execute(wr, data)
		} else {
			err = t.template.ExecuteTemplate(wr, name, data)
		}

		return
	}
}
