package gox

import (
	"time"
)

// Updated 带修改时间模型
type Updated struct {
	// 最后更新时间
	Updated time.Time `xorm:"updated default('2020-02-04 09:55:52')" json:"updated"`
}
