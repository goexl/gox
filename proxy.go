package gox

import (
	`fmt`
)

const (
	// ProxyTypeSocksV4 Socks代码
	ProxyTypeSocksV4 ProxyType = "socks4"
	// ProxyTypeSocksV5 Socks代码
	ProxyTypeSocksV5 ProxyType = "socks4"
	// ProxyTypeSocksV4 Socks代码
	ProxyTypeHttp ProxyType = "http"
	// ProxyTypeSocksV4 Socks代码
	ProxyTypeHttps ProxyType = "https"
)

type (
	// ProxyType
	ProxyType string

	// Proxy 代理配置
	Proxy struct {
		// Host 主机（可以是Ip或者域名）
		Host string `json:"ip" yaml:"ip" validate:"required"`
		// Port 端口
		Port int `default:"80" json:"port" yaml:"port" validate:"required"`
		// Type 代理类型
		Type ProxyType `default:"http" json:"type" yaml:"type" validate:"required,oneof=socks4 socks5 http https"`
		// Username 代理认证用户名
		Username string `json:"username" yaml:"username"`
		// Password 代理认证密码
		Password string `json:"password" yaml:"password"`
	}
)

func (p *Proxy) Addr() (addr string) {
	if "" != p.Username && "" != p.Password {
		addr = fmt.Sprintf("%s://%s:%s@%s:%d", p.Type, p.Username, p.Password, p.Host, p.Port)
	} else {
		addr = fmt.Sprintf("%s://%s:%d", p.Type, p.Host, p.Port)
	}

	return
}
