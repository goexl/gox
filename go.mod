module github.com/goexl/gox

go 1.17

require (
	github.com/mitchellh/mapstructure v1.4.3
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29
	golang.org/x/text v0.3.7
)

// v1 项目从原来的storezhang/gox迁移过来，原来的版本号不再使用，直到最新发布到该版本
retract [v1.0.0, v1.8.4]
