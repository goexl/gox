package gox

import (
	"fmt"
	"strings"
	`time`
)

type (
	// Connection 连接池配置
	Connection struct {
		// MaxOpen 最大打开连接数
		MaxOpen int `default:"150" yaml:"maxOpen" json:"maxOpen"`
		// MaxIdle 最大休眠连接数
		MaxIdle int `default:"30" yaml:"maxIdle" json:"maxIdle"`
		// MaxLifetime 每个连接最大存活时间
		MaxLifetime time.Duration `default:"5s" yaml:"maxLifetime" json:"maxLifetime"`
	}

	// Database 数据库配置
	Database struct {
		// Type 数据库类型，支持
		Type string `default:"sqlite3"`

		Address  string `default:"127.0.0.1:3306"`
		Username string
		Password string
		Protocol string `default:"tcp"`
		// Connection 连接池配置
		Connection Connection

		Suffix             string
		Prefix             string
		Schema             string `default:"schema"`
		MigrationTableName string `default:"migration"`
		Parameters         string
		Path               string `default:"data.db"`
	}
)

// Dsn 取得数据库的DSN
func (db *Database) Dsn() string {
	var dsn string

	switch strings.ToUpper(db.Type) {
	case "MYSQL":
		dsn = fmt.Sprintf(
			"%s:%s@%s(%s)",
			db.Username, db.Password,
			db.Protocol, db.Address,
		)
		if "" != strings.TrimSpace(db.Schema) {
			dsn = fmt.Sprintf("%s/%s", dsn, strings.TrimSpace(db.Schema))
		}
	case "SQLITE3":
		dsn = db.Path
	default:
		panic("不支持的数据库类型")
	}

	// 增加参数
	if "" != strings.TrimSpace(db.Parameters) {
		dsn = fmt.Sprintf("%s?%s", dsn, strings.TrimSpace(db.Parameters))
	}

	return dsn
}
