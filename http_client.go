package gox

import (
	`net/http`
	`time`
)

// HttpClient Http客户端配置
type (
	HttpClient struct {
		// Timeout 超时
		Timeout time.Duration `json:"timeout" yaml:"timeout"`
		// Proxy 代理
		Proxy Proxy `json:"proxy" yaml:"proxy"`
		// Auth 授权配置
		Auth Auth `json:"auth" yaml:"auth"`
		// AllowGetPayload 是否允许Get方法传输数据
		AllowGetPayload bool `default:"true" json:"allowGetPayload" yaml:"allowGetPayload"`
		// QueryParams 通用的查询参数
		QueryParams map[string]string `json:"queryParams" yaml:"queryParams"`
		// FormData 表单参数，只对POST和PUT方法有效
		FormData map[string]string `json:"formData" yaml:"formData"`
		// Headers 通用头信息
		Headers map[string]string `json:"headers" yaml:"headers"`
		// Cookies 通用Cookie
		Cookies []http.Cookie `json:"cookies" yaml:"cookies"`
	}
)
