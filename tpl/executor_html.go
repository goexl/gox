package tpl

import (
	"bytes"
	"html/template"
	"io"

	"github.com/goexl/gox/rand"
)

var _ executor = (*executorHtml)(nil)

type executorHtml struct {
	*base

	template *template.Template
}

func newHtmlExecutor() *executorHtml {
	return new(executorHtml)
}

func (eh *executorHtml) toFile(input []string, inputType inputType, data any, filepath string) (err error) {
	tpl := template.New(filepath)
	switch inputType {
	case inputTypeFile:
		eh.template, err = tpl.ParseFiles(input...)
	case inputTypeString:
		eh.template, err = tpl.Parse(input[0])
	default:
		eh.template, err = tpl.Parse(input[0])
	}

	if nil == err {
		err = eh.renderToFile(filepath, eh.execute(data))
	}

	return
}

func (eh *executorHtml) toString(input []string, inputType inputType, data any) (result string, err error) {
	tpl := template.New(rand.New().String().Generate())
	switch inputType {
	case inputTypeFile:
		eh.template, err = tpl.ParseFiles(input...)
	case inputTypeString:
		eh.template, err = tpl.Parse(input[0])
	default:
		eh.template, err = tpl.Parse(input[0])
	}
	if nil != err {
		return
	}

	buffer := new(bytes.Buffer)
	if err = eh.template.Execute(buffer, data); nil == err {
		result = buffer.String()
	}

	return
}

func (eh *executorHtml) execute(data any) fun {
	return func(wr io.Writer) (err error) {
		return eh.template.Execute(wr, data)
	}
}
