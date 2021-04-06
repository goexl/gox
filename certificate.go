package gox

type (
	// ClientCertificate 客户端秘钥
	ClientCertificate struct {
		// Public 公钥文件路径
		Public string `json:"public" yaml:"public" validate:"required,file"`
		// Private 私钥文件路径
		Private string `json:"private" yaml:"private" validate:"required,file"`
	}

	// Certificate 秘钥
	Certificate struct {
		// Skip 是否跳过TLS检查
		Skip bool `default:"true" json:"skip" yaml:"skip"`
		// Root 根秘钥文件路径
		Root string `json:"root" yaml:"root" validate:"required,file"`
		// Clients 客户端
		Clients []ClientCertificate `json:"clients" yaml:"clients" validate:"structonly"`
	}
)
