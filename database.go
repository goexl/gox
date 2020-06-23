package gox

import (
	"fmt"
	"strings"
)

// Database 数据库配置
type Database struct {
	Type string `default:"sqlite3"`

	Address  string `default:"127.0.0.1:3306"`
	Username string
	Password string
	Protocol string `default:"tcp"`

	Suffix             string
	Prefix             string
	Schema             string `default:"schema"`
	MigrationTableName string `default:"migration"`
	Parameters         string
	Path               string `default:"data.db"`
}

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
