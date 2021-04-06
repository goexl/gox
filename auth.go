package gox

const (
	// AuthTypeBasic 基本Http授权验证
	AuthTypeBasic AuthType = "basic"
	// AuthTypeToken 基本传Token的授权验证
	AuthTypeToken AuthType = "token"
)

type (
	// AuthType 授权类型
	AuthType string

	// Auth 授权信息
	Auth struct {
		// Type 授权类型
		Type AuthType `default:"type" json:"type" yaml:"type" validate:"oneof=basic token"`
		// Username 用户名
		Username string `json:"username" yaml:"username"`
		// Password 密码
		Password string `json:"password" yaml:"password"`
		// Token 授权码
		Token string `json:"token" yaml:"token"`
		// Scheme 身份验证方案类型
		Scheme string `json:"scheme" yaml:"scheme"`
	}
)
