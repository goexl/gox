package gox

// EmailConfig 邮箱配置
type EmailConfig struct {
	// 主机
	Host string `validate:"required"`
	// 端口
	Port int `validate:"required"`
	// 用户名
	Username string `validate:"required"`
	// 密码
	Password string `validate:"required"`
}
