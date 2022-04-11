package gox

type (
	sizeOption interface {
		applySize(options *sizeOptions)
	}

	sizeOptions struct {
		unit      Size
		separator rune
	}
)

func defaultOptions() *sizeOptions {
	return &sizeOptions{
		unit:      SizeMB,
		separator: RuneSpace,
	}
}
