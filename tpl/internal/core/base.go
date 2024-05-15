package core

import (
	"os"
	"path/filepath"
)

type Base struct {
	// 空结构
}

func (b *Base) RenderToFile(path string, fun Function) (err error) {
	if file, gfe := b.getFile(path); nil != gfe {
		b.close(file)
		err = gfe
	} else {
		err = b.render(file, fun)
	}

	return
}

func (b *Base) render(file *os.File, fun Function) (err error) {
	defer b.close(file)
	err = fun(file)

	return
}

func (b *Base) getFile(path string) (file *os.File, err error) {
	if _, se := os.Stat(path); nil != se && os.IsNotExist(se) {
		file, err = b.create(path)
	} else {
		file, err = os.OpenFile(path, os.O_WRONLY, os.ModePerm)
	}

	return
}

func (b *Base) create(path string) (file *os.File, err error) {
	dir := filepath.Dir(path)
	if _, se := os.Stat(dir); nil != se && os.IsNotExist(se) {
		err = os.MkdirAll(dir, os.ModePerm)
	}
	if nil == err {
		file, err = os.Create(path)
	}

	return
}

func (b *Base) close(file *os.File) {
	_ = file.Close()
}
