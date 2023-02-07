package tpl

import (
	"os"
)

type fun func(file *os.File) (err error)
