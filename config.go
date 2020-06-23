package gox

type (
	// CROS 跨域配置
	CROS struct {
		AllowOrigins     []string `default:"['*']" yaml:"allowOrigins"`
		AllowCredentials bool     `default:"true" yaml:"allowCredentials"`
	}
)
