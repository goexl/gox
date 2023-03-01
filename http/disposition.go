package http

import (
	"fmt"
	"net/url"
)

type disposition struct {
	params *dispositionParams
}

func newDisposition(params *dispositionParams) *disposition {
	return &disposition{
		params: params,
	}
}

func (d *disposition) String() (disposition string) {
	// 文件名需要编码，不然中文会乱码
	filename := url.QueryEscape(d.params.filename)
	disposition = fmt.Sprintf("%s; filename=%s;filename*=utf-8''%s", d.params.disposition, filename, filename)

	return
}

func (d *disposition) Set(headers Headers) {
	headers.Set(HeaderContentDisposition, d.String())
}

func (d *disposition) Headers() (headers Headers) {
	headers = make(Headers)
	d.Set(headers)

	return
}
