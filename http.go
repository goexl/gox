package gox

const (
	HeaderContentType = "Content-Type"

	// Http方法集合
	// GET方法
	HttpMethodGet HttpMethod = "get"
	// POST方法
	HttpMethodPost HttpMethod = "post"
	// PUT方法
	HttpMethodPut HttpMethod = "put"
	// DELETE方法
	HttpMethodDelete HttpMethod = "delete"
)

type (
	// HttpMethod Http方法
	HttpMethod string
)
