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

func (eh *executorHtml) toFile(input []string, it inputType, name string, data any, filepath string) (err error) {
	if err = eh.init(input, it); nil == err {
		err = eh.renderToFile(filepath, eh.execute(name, data))
	}

	return
}

func (eh *executorHtml) toString(input []string, it inputType, name string, data any) (result string, err error) {
	if err = eh.init(input, it); nil != err {
		return
	}

	buffer := new(bytes.Buffer)
	if "" == name {
		err = eh.template.Execute(buffer, data)
	} else {
		err = eh.template.ExecuteTemplate(buffer, name, data)
	}
	if nil == err {
		result = buffer.String()
	}

	return
}

func (eh *executorHtml) init(input []string, it inputType) (err error) {
	switch it {
	case inputTypeFile:
		eh.template, err = template.ParseFiles(input...)
	case inputTypeString:
		fallthrough
	default:
		eh.template, err = template.New(rand.New().String().Build().Generate()).Parse(input[0])
	}

	return
}

func (eh *executorHtml) execute(name string, data any) fun {
	return func(wr io.Writer) (err error) {
		if "" == name {
			err = eh.template.Execute(wr, data)
		} else {
			err = eh.template.ExecuteTemplate(wr, name, data)
		}

		return
	}
}
