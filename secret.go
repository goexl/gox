package gox

import (
	"fmt"
)

// Secret 应用授权
type Secret struct {
	// 授权，相当于用户名
	Ak string `json:"ak" yaml:"ak" xml:"ak" toml:"ak" validate:"required"`
	// 授权，相当于密码
	Sk string `json:"sk" yaml:"sk" xml:"sk" toml:"sk" validate:"required"`
}

func (s *Secret) Key() string {
	return fmt.Sprintf("%s-%s", s.Ak, s.Sk)
}
