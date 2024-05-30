package core

import (
	"fmt"
	"net/url"

	"github.com/goexl/gox/http"
	"github.com/goexl/gox/http/internal/param"
)

type Disposition struct {
	params *param.Disposition
}

func NewDisposition(params *param.Disposition) *Disposition {
	return &Disposition{
		params: params,
	}
}

func (d *Disposition) String() (disposition string) {
	// 文件名需要编码，不然中文会乱码
	filename := url.QueryEscape(d.params.Filename)
	disposition = fmt.Sprintf("%s; Filename=%s;Filename*=utf-8''%s", d.params.Disposition, filename, filename)

	return
}

func (d *Disposition) Set(headers http.Header) {
	headers.Set(http.HeaderContentDisposition, d.String())
}

func (d *Disposition) Headers() (headers http.Header) {
	headers = make(http.Header)
	d.Set(headers)

	return
}
