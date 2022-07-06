package gox

var (
	_ mapstructOption = (*optionZeroFields)(nil)
	_                 = ZeroFields
)

type optionZeroFields struct {
	zeroFields bool
}

// ZeroFields 配置处理零值
func ZeroFields(zeroFields bool) *optionZeroFields {
	return &optionZeroFields{
		zeroFields: zeroFields,
	}
}

func (zf *optionZeroFields) applyMapstruct(options *mapstructOptions) {
	options.zeroFields = zf.zeroFields
}
