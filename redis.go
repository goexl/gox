package gox

// RedisConfig Redis配置
type RedisConfig struct {
	// Addr 通信地址
	Addr string `default:"127.0.0.1:6379" json:"addr" yaml:"addr" validate:"required"`
	// Password 授权密码
	Password string `json:"password" yaml:"password"`
	// DB 数据库
	DB int `default:"0" json:"db" yaml:"db"`
}
