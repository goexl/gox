package http

var _ = Header

type header struct{}

// Header 头操作
func Header() *header {
	return new(header)
}

func (h *header) Disposition(filename string) *dispositionBuilder {
	return newDispositionBuilder(filename)
}
