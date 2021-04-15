package gox

import (
	`net/http`
)

// ElasticsearchConfig 配置
type ElasticsearchConfig struct {
	// Address 地址
	Address string `json:"address" yaml:"address" validate:"required,uri"`
	// Username 用户名
	Username string `json:"username" yaml:"username"`
	// Password 密码
	Password    string `json:"password" yaml:"password"`
	HealthCheck bool
	Sniff       bool
	// Gzip 配置是否启用压缩
	Gzip bool `default:"false" json:"gzip" yaml:"gzip"`
	// Headers 传输头
	Headers http.Header `json:"headers" yaml:"headers"`
	// Plugins 配置启动时的必备插件
	Plugins []string `json:"plugins" yaml:"plugins"`
}
