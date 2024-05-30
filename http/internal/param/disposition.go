package param

import (
	"github.com/goexl/gox/http/internal/internal/constant"
)

type Disposition struct {
	Filename    string
	Disposition string
}

func NewDisposition(filename string) *Disposition {
	return &Disposition{
		Filename:    filename,
		Disposition: constant.Attachment,
	}
}
