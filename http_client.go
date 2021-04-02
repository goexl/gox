package gox

import (
	`time`
)

// HttpClient Http客户端配置
type (
	HttpClient struct {
		// Timeout 超时
		Timeout time.Duration `json:"timeout" yaml:"timeout"`
		// Proxy 代理
		Proxy Proxy `json:"proxy" yaml:"proxy"`
	}
)
