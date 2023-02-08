package tpl

import (
	"io"
)

type fun func(wr io.Writer) (err error)
