package gox

const (
	// AuthTypeBasic 基本Http授权验证
	AuthTypeBasic AuthType = "basic"
	// AuthTypeToken 基本传Token的授权验证
	AuthTypeToken AuthType = "token"
)

var (
	_ = AuthTypeBasic
	_ = AuthTypeToken
)

type (
	// AuthType 授权类型
	AuthType string

	// AuthConfig 授权信息
	AuthConfig struct {
		// 授权类型
		Type AuthType `default:"basic" json:"type" yaml:"type" xml:"type" toml:"type" validate:"oneof=basic token"`
		// 用户名
		Username string `json:"username" yaml:"username" xml:"username" toml:"username"`
		// 密码
		Password string `json:"password" yaml:"password" xml:"password" toml:"password"`
		// 授权码
		Token string `json:"token" yaml:"token" xml:"token" toml:"token"`
		// 身份验证方案类型
		Scheme URIScheme `json:"scheme" yaml:"scheme" xml:"scheme" toml:"scheme"`
	}
)
