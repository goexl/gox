package gox

const (
	// Http常用头
	// 请求数据类型
	HeaderContentType = "Content-Type"
	// 授权
	HeaderAuthorization = "Authorization"

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
