package gox

import (
	`time`
)

// CheckConfig 检查器配置，可以用于多种场景：健康检查、嗅探器等
type CheckConfig struct {
	// Interval 间隔时间
	Interval time.Duration `json:"interval" yaml:"interval" validate:"required"`
	// Timeout 超时时间
	Timeout time.Duration `json:"timeout" yaml:"timeout" validate:"required"`
	// StartupTimeout 启动时的超时时间
	StartupTimeout time.Duration `json:"startupTimeout" yaml:"startupTimeout" validate:"required"`
}
