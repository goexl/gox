package gox

type (
	// ClientCertificate 客户端秘钥
	ClientCertificate struct {
		// 公钥文件路径
		Public string `json:"public" yaml:"public" validate:"required,file"`
		// 私钥文件路径
		Private string `json:"private" yaml:"private" validate:"required,file"`
	}

	// CertificateConfig 秘钥
	CertificateConfig struct {
		// 是否跳过TLS检查
		Skip bool `default:"true" json:"skip" yaml:"skip"`
		// 根秘钥文件路径
		Root string `json:"root" yaml:"root" validate:"required,file"`
		// 客户端
		Clients []ClientCertificate `json:"clients" yaml:"clients" validate:"structonly"`
	}
)
