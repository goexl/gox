package gox

type pageOptions struct {
	size   int
	page   int
	fields []Field
}

func defaultPageOptions() *pageOptions {
	return &pageOptions{
		size: 20,
		page: 1,
	}
}
