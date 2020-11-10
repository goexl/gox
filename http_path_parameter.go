package gox

// HttpPathParameter Http路径参数
type HttpPathParameter struct {
	// key 键
	key string
	// value 值
	value string
}

// NewHttpPathParameter 创建一个Http路径参数
func NewHttpPathParameter(key string, value string) *HttpPathParameter {
	return &HttpPathParameter{
		key:   key,
		value: value,
	}
}

func (hh *HttpPathParameter) Type() HttpParameterType {
	return HttpParameterTypePathParameter
}

func (hh *HttpPathParameter) Key() string {
	return hh.key
}

func (hh *HttpPathParameter) Value() string {
	return hh.value
}
