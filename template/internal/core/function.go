package core

import (
	"io"
)

type Function func(wr io.Writer) (err error)
