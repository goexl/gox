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

func (d *disposition) Fill(headers Headers) {
	headers[HeaderContentDisposition] = d.String()
}
