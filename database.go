package gox

import (
	"fmt"
	"strings"
	`time`
)

type (
	// ConnectionConfig 连接池配置
	ConnectionConfig struct {
		// MaxOpen 最大打开连接数
		MaxOpen int `default:"150" yaml:"maxOpen" json:"maxOpen"`
		// MaxIdle 最大休眠连接数
		MaxIdle int `default:"30" yaml:"maxIdle" json:"maxIdle"`
		// MaxLifetime 每个连接最大存活时间
		MaxLifetime time.Duration `default:"5s" yaml:"maxLifetime" json:"maxLifetime"`
	}

	// DatabaseConfig 数据库配置
	DatabaseConfig struct {
		// Type 数据库类型，支持
		Type string `default:"sqlite3" json:"type" yaml:"type"`

		// Address 地址，填写服务器地址
		Address string `default:"127.0.0.1:3306" json:"address" validate:"required"`
		// Username 授权，用户名
		Username string `json:"username,omitempty" yaml:"username"`
		// Password 授权，密码
		Password string `json:"password,omitempty" yaml:"password"`
		// Protocol 协议，默认用Tcp
		Protocol string `default:"tcp" json:"protocol" yaml:"protocol"`
		// Connection 连接池配置
		Connection ConnectionConfig `json:"connection" yaml:"connection"`

		// Suffix 表名的前缀
		Suffix string `json:"suffix" yaml:"suffix"`
		// Prefix 表名后缀
		Prefix string `json:"prefix" yaml:"prefix"`
		// Schema 连接的数据库名
		Schema             string `default:"schema" json:"schema" yaml:"schema"`
		MigrationTableName string `default:"migration" json:"migration_table_name"`
		// Parameters 额外参数
		Parameters string `json:"parameters"`
		// Path SQLite填写数据库文件的路径
		Path string `default:"data.db" json:"path"`
	}
)

// Dsn 取得数据库的DSN
func (dc *DatabaseConfig) Dsn() string {
	var dsn string

	switch strings.ToUpper(dc.Type) {
	case "MYSQL":
		dsn = fmt.Sprintf(
			"%s:%s@%s(%s)",
			dc.Username, dc.Password,
			dc.Protocol, dc.Address,
		)
		if "" != strings.TrimSpace(dc.Schema) {
			dsn = fmt.Sprintf("%s/%s", dsn, strings.TrimSpace(dc.Schema))
		}
	case "SQLITE3":
		dsn = dc.Path
	default:
		panic("不支持的数据库类型")
	}

	// 增加参数
	if "" != strings.TrimSpace(dc.Parameters) {
		dsn = fmt.Sprintf("%s?%s", dsn, strings.TrimSpace(dc.Parameters))
	}

	return dsn
}
