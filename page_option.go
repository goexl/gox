package gox

type (
	pageOption interface {
		applyPage(options *pageOptions)
	}

	pageOptions struct {
		size   int32
		page   int32
		fields []Field
	}
)

func defaultPageOptions() *pageOptions {
	return &pageOptions{
		size: 20,
		page: 1,
	}
}
