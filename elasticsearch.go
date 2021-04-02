package gox

// Elasticsearch 配置
type Elasticsearch struct {
	// Address 地址
	Address string `yaml:"address" validate:"required,uri"`
	// Username 用户名
	Username string `yaml:"username"`
	// Password 密码
	Password string `yaml:"password"`
}
