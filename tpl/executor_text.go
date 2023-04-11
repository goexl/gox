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

func (et *executorText) toFile(input []string, it inputType, name string, data any, filepath string) (err error) {
	if err = et.init(input, it); nil == err {
		err = et.renderToFile(filepath, et.execute(name, data))
	}

	return
}

func (et *executorText) toString(input []string, it inputType, name string, data any) (result string, err error) {
	if err = et.init(input, it); nil != err {
		return
	}

	buffer := new(bytes.Buffer)
	if "" == name {
		err = et.template.Execute(buffer, data)
	} else {
		err = et.template.ExecuteTemplate(buffer, name, data)
	}
	if nil == err {
		result = buffer.String()
	}

	return
}

func (et *executorText) init(input []string, it inputType) (err error) {
	switch it {
	case inputTypeFile:
		et.template, err = template.ParseFiles(input...)
	case inputTypeString:
		fallthrough
	default:
		et.template, err = template.New(rand.New().String().Build().Generate()).Parse(input[0])
	}

	return
}

func (et *executorText) execute(name string, data any) fun {
	return func(wr io.Writer) (err error) {
		if "" == name {
			err = et.template.Execute(wr, data)
		} else {
			err = et.template.ExecuteTemplate(wr, name, data)
		}

		return
	}
}
