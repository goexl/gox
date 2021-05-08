package gox

const (
	// EnvironmentDev 开发
	EnvironmentDev Environment = "dev"
	// EnvironmentTest 测试
	EnvironmentTest Environment = "test"
	// EnvironmentQa 测试
	EnvironmentQa Environment = "qa"
	// EnvironmentRelease 发布环境
	EnvironmentRelease Environment = "release"
	// EnvironmentProd 生产
	EnvironmentProd Environment = "prod"
	// EnvironmentLocal 本地环境
	EnvironmentLocal Environment = "local"
	// EnvironmentSimulation 模拟请求（不发真实请求到服务器）
	EnvironmentSimulation Environment = "simulation"
)

// Environment 环境类型
type Environment string
