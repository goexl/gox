package tpl

import (
	"bytes"
	"io"
	"text/template"

	"github.com/goexl/gox/rand"
)

var _ executor = (*executorText)(nil)

type executorText struct {
	*base

	template *template.Template
}

func newTextExecutor() *executorText {
	return new(executorText)
}

func (et *executorText) toFile(input []string, inputType inputType, data any, filepath string) (err error) {
	tpl := template.New(filepath)
	switch inputType {
	case inputTypeFile:
		et.template, err = tpl.ParseFiles(input...)
	case inputTypeString:
		et.template, err = tpl.Parse(input[0])
	default:
		et.template, err = tpl.Parse(input[0])
	}

	if nil == err {
		err = et.renderToFile(filepath, et.execute(data))
	}

	return
}

func (et *executorText) toString(input []string, inputType inputType, data any) (result string, err error) {
	tpl := template.New(rand.New().String().Build().Generate())
	switch inputType {
	case inputTypeFile:
		et.template, err = tpl.ParseFiles(input...)
	case inputTypeString:
		et.template, err = tpl.Parse(input[0])
	default:
		et.template, err = tpl.Parse(input[0])
	}
	if nil != err {
		return
	}

	buffer := new(bytes.Buffer)
	if err = et.template.Execute(buffer, data); nil == err {
		result = buffer.String()
	}

	return
}

func (et *executorText) execute(data any) fun {
	return func(wr io.Writer) (err error) {
		return et.template.Execute(wr, data)
	}
}
