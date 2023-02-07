package tpl

import (
	"os"
	"text/template"
)

var _ executor = (*executorText)(nil)

type executorText struct {
	template *template.Template
}

func newTextExecutor() *executorText {
	return new(executorText)
}

func (et *executorText) toFile(template string, data any, filename string) (err error) {
	if isFile {
		t = template.Must(template.New(filename).ParseFiles(input))
	} else {
		t = template.Must(template.New(filename).Parse(input))
	}

	if file, err = os.Create(filename); nil != err {
		return
	}
	defer func() {
		err = file.Close()
	}()

	err = t.Execute(file, data)

	return

	return
}
