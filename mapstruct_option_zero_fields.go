package gox

var _ mapstructOption = (*mapstructOptionZeroFields)(nil)

type mapstructOptionZeroFields struct {
	zeroFields bool
}

// ZeroFields 配置处理零值
func ZeroFields(zeroFields bool) *mapstructOptionZeroFields {
	return &mapstructOptionZeroFields{
		zeroFields: zeroFields,
	}
}

func (s *mapstructOptionZeroFields) apply(options *mapstructOptions) {
	options.zeroFields = s.zeroFields
}
