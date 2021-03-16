package gox

// Redis Redis配置
type Redis struct {
	// Addr 通信地址
	Addr string `default:"127.0.0.1:6379" validate:"required"`
	// Password 授权密码
	Password string
	// DB 数据库
	DB int `default:"0"`
}
