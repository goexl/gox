package tpl

import (
	"os"
	"text/template"
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
	switch inputType {
	case inputTypeFile:
		et.template = template.Must(template.New(filepath).ParseFiles(input...))
	case inputTypeString:
		et.template = template.Must(template.New(filepath).Parse(input[0]))
	default:
		et.template = template.Must(template.New(filepath).Parse(input[0]))
	}
	err = et.renderToFile(filepath, func(file *os.File) (err error) {
		return et.template.Execute(file, data)
	})

	return
}

func (et *executorText) toString(input []string, inputType inputType, data any) (result string, err error) {
	return
}
