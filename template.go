package gox

import (
	`os`
	`text/template`
)

// RenderToFile 渲染模板到文件
func RenderToFile(filename string, input string, data interface{}, isFile bool) (err error) {
	var (
		file *os.File
		t    *template.Template
	)

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
}
