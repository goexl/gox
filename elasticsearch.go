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
	Password string `json:"password" yaml:"password"`
	// Health 健康检查配置
	Health CheckConfig `json:"health" yaml:"health" validate:"structonly"`
	// Sniffer 嗅探器配置
	Sniffer CheckConfig `json:"sniffer" yaml:"sniffer" validate:"structonly"`
	// Gzip 配置是否启用压缩
	Gzip bool `default:"false" json:"gzip" yaml:"gzip"`
	// Headers 传输头
	Headers http.Header `json:"headers" yaml:"headers"`
	// Plugins 配置启动时的必备插件
	Plugins []string `json:"plugins" yaml:"plugins"`
}
