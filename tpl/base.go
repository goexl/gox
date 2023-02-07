package tpl

import (
	"os"
)

type base struct{}

func (b *base) renderToFile(filepath string, fun fun) (err error) {
	var file *os.File
	if _, se := os.Stat(filepath); nil != se && os.IsNotExist(se) {
		file, err = os.Create(filepath)
	} else {
		file, err = os.Open(filepath)
	}
	if nil != err {
		return
	}

	defer b.close(file, &err)
	err = fun(file)

	return
}

func (b *base) close(file *os.File, err *error) {
	if ce := file.Close(); nil != ce {
		*err = ce
	}
}
