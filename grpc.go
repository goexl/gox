package gox

// Grpc gRPC配置
type Grpc struct {
	// Ip 监听的地址
	Ip string `default:"localhost" json:"ip" yaml:"ip" validate:"required"`
	// Port 监听的端口
	Port int `default:"8081" json:"port" yaml:"port" validate:"required"`
}
