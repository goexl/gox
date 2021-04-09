package gox

import (
	`fmt`
)

// Grpc gRPC配置
type Grpc struct {
	// Host 监听的地址，可以是IP也可以是域名
	Host string `json:"ip" yaml:"ip"`
	// Port 监听的端口
	Port int `default:"8081" json:"port" yaml:"port" validate:"required"`
}

// Addr 返回gRPC需要的地址
func (g *Grpc) Addr() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}
