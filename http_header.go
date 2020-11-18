package gox

// HttpHeader Http请求头
type HttpHeader struct {
	// key 键
	key string
	// value 值
	value string
}

// NewHttpHeader 创建一个Http请求头
func NewHttpHeader(key string, value string) *HttpHeader {
	return &HttpHeader{
		key:   key,
		value: value,
	}
}

func (hh *HttpHeader) Type() HttpParameterType {
	return HttpParameterTypeHeader
}

func (hh *HttpHeader) Key() string {
	return hh.key
}

func (hh *HttpHeader) Value() string {
	return hh.value
}
