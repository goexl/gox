package gox

import (
	"fmt"
	"net/url"
)

// ProxyConfig 代理配置
type ProxyConfig struct {
	// 主机，可以是地址或者域名
	Host string `json:"host" yaml:"host" xml:"host" toml:"host" validate:"required"`
	// 端口
	Port int `default:"80" json:"port" yaml:"port" xml:"port" toml:"port" validate:"required"`
	// 代理类型
	Scheme URIScheme `default:"http" json:"scheme" yaml:"scheme" xml:"scheme" toml:"scheme" validate:"required,oneof=socks4 socks5 http https"`
	// 代理认证用户名
	Username string `json:"username" yaml:"username" xml:"username" toml:"username"`
	// 代理认证密码
	Password string `json:"password" yaml:"password" xml:"password" toml:"password"`
}

func (pc *ProxyConfig) Addr() (addr string) {
	if "" != pc.Username && "" != pc.Password {
		addr = fmt.Sprintf(
			"%s://%s:%s@%s:%d",
			pc.Scheme,
			url.QueryEscape(pc.Username), url.QueryEscape(pc.Password),
			pc.Host, pc.Port,
		)
	} else {
		addr = fmt.Sprintf("%s://%s:%d", pc.Scheme, pc.Host, pc.Port)
	}

	return
}
