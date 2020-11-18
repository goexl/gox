package gox

import (
	`fmt`
	`net/url`
)

const (
	// Http常用头
	// HeaderAcceptLanguage 可接受的语言
	HeaderAcceptLanguage = "Accept-Language"
	// HeaderContentType 请求数据类型
	HeaderContentType = "Content-Type"
	// HeaderAuthorization 授权
	HeaderAuthorization = "Authorization"
	// HeaderContentDisposition
	HeaderContentDisposition = "Content-disposition"

	// Http方法集合
	// HttpMethodGet GET方法
	HttpMethodGet HttpMethod = "GET"
	// HttpMethodPost POST方法
	HttpMethodPost HttpMethod = "POST"
	// HttpMethodPut PUT方法
	HttpMethodPut HttpMethod = "PUT"
	// HttpMethodDelete DELETE方法
	HttpMethodDelete HttpMethod = "DELETE"
	// MethodPatch PATCH方法
	HttpMethodPatch HttpMethod = "PATCH"
	// MethodHead HEAD方法
	HttpMethodHead HttpMethod = "HEAD"
	// MethodOptions OPTIONS方法
	HttpMethodOptions HttpMethod = "OPTIONS"
)

const (
	// ContentDispositionTypeAttachment 附件下载
	ContentDispositionTypeAttachment ContentDispositionType = "attachment"
	// ContentDispositionTypeInline 浏览器直接打开
	ContentDispositionTypeInline ContentDispositionType = "inline"
)

const (
	// ContentDispositionTypeHeader 请求头
	HttpParameterTypeHeader HttpParameterType = "header"
	// ContentDispositionTypePathParameter 路径参数
	HttpParameterTypePathParameter HttpParameterType = "path"
)

type (
	// HttpMethod Http方法
	HttpMethod string

	// ContentDispositionType 下载类型
	ContentDispositionType string

	// HttpParameterType Http额外参数类型
	HttpParameterType string

	// HttpParameter Http额外参数接口
	HttpParameter interface {
		// Type 类型
		Type() HttpParameterType
		// Key 键
		Key() string
		// Value 值
		Value() string
	}
)

// ContentDisposition 解决附件下载乱码
func ContentDisposition(filename string, dispositionType ContentDispositionType) (disposition string) {
	// 文件名需要编码
	filename = url.QueryEscape(filename)
	disposition = fmt.Sprintf("%s; filename=%s;filename*=utf-8''%s", dispositionType, filename, filename)

	return
}
