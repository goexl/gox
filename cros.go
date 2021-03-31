package gox

// CROS 跨域配置
type CROS struct {
	AllowOrigins     []string `default:"['*']" yaml:"allowOrigins"`
	AllowCredentials bool     `default:"true" yaml:"allowCredentials"`
}
