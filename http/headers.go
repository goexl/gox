package http

const (
	// HeaderAcceptLanguage 可接受的语言
	HeaderAcceptLanguage = "Accept-Language"
	// HeaderContentType 请求数据类型
	HeaderContentType = "Content-Type"
	// HeaderAuthorization 授权
	HeaderAuthorization = "Authorization"
	// HeaderContentDisposition 下载时的附件名指定
	HeaderContentDisposition = "Content-Disposition"
)

// 禁止IDE提示未使用
var (
	_ = HeaderAcceptLanguage
	_ = HeaderContentType
	_ = HeaderAuthorization
	_ = HeaderContentDisposition
)
