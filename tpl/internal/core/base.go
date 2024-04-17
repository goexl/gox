package core

import (
	"os"
)

type Base struct {
	// 空结构
}

func (b *Base) RenderToFile(filepath string, fun Function) (err error) {
	var file *os.File
	if _, se := os.Stat(filepath); nil != se && os.IsNotExist(se) {
		file, err = os.Create(filepath)
	} else {
		file, err = os.Open(filepath)
	}
	if nil != err {
		return
	}

	defer b.Close(file, &err)
	err = fun(file)

	return
}

func (b *Base) Close(file *os.File, err *error) {
	if ce := file.Close(); nil != ce {
		*err = ce
	}
}
