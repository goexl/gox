package gox

var _ pageOption = (*optionSize)(nil)

type optionSize struct {
	size int
}

// Size 配置大小
func Size(size int) *optionSize {
	return &optionSize{
		size: size,
	}
}

func (s *optionSize) applyPage(options *pageOptions) {
	options.size = s.size
}
