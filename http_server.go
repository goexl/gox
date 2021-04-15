package gox

// HttpServerConfig Http服务器配置
type (
	HttpServerConfig struct {
		// Host 监听的地址，可以是IP也可以是域名
		Host string `default:"localhost" json:"ip" yaml:"ip" validate:"required"`
		// Port 监听的端口
		Port int `default:"8081" json:"port" yaml:"port" validate:"required"`
		// BasePath 前缀地址
		BasePath string `yaml:"basePath"`
		// CROS 跨域配置
		CROS CROS
	}
)
