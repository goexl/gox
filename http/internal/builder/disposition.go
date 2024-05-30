package builder

import (
	"github.com/goexl/gox/http/internal/core"
	"github.com/goexl/gox/http/internal/internal/constant"
	"github.com/goexl/gox/http/internal/param"
)

type Disposition struct {
	params *param.Disposition
}

func NewDisposition(filename string) *Disposition {
	return &Disposition{
		params: param.NewDisposition(filename),
	}
}

func (d *Disposition) Inline() (disposition *Disposition) {
	d.params.Disposition = constant.Inline
	disposition = d

	return
}

func (d *Disposition) Attachment() (disposition *Disposition) {
	d.params.Disposition = constant.Attachment

	return d
}

func (d *Disposition) Download() *Disposition {
	return d.Attachment()
}

func (d *Disposition) Build() *core.Disposition {
	return core.NewDisposition(d.params)
}
