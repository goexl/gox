package gox

var _ pageOption = (*optionSize)(nil)

type optionSize struct {
	size int32
}

// Size 配置大小
func Size(size int32) *optionSize {
	return &optionSize{
		size: size,
	}
}

func (s *optionSize) applyPage(options *pageOptions) {
	options.size = s.size
}
