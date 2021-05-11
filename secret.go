package gox

// Secret 描述一个第三方应用授权
type Secret struct {
	// 授权，相当于用户名
	Id string `json:"id" yaml:"id" validate:"required"`
	// 授权，相当于密码
	Key string `json:"key" yaml:"key" validate:"required"`
}
