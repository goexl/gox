package gox

type pageOptions struct {
	size  int
	page  int
	extra map[string]interface{}
}

func defaultPageOptions() *pageOptions {
	return &pageOptions{
		size: 20,
		page: 1,
	}
}
