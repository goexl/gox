package gox

var (
	_            = Separator
	_ sizeOption = (*optionSeparator)(nil)
)

type optionSeparator struct {
	separator rune
}

// Separator 分隔符
func Separator(separator rune) *optionSeparator {
	return &optionSeparator{
		separator: separator,
	}
}

func (s *optionSeparator) applySize(options *sizeOptions) {
	options.separator = s.separator
}
