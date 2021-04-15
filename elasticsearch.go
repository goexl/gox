package gox

// ElasticsearchConfig 配置
type ElasticsearchConfig struct {
	// Address 地址
	Address string `json:"address" yaml:"address" validate:"required,uri"`
	// Username 用户名
	Username string `json:"username" yaml:"username"`
	// Password 密码
	Password string `json:"password" yaml:"password"`
}
