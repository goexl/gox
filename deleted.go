package gox

import (
	"time"
)

// Deleted 软删除模型
type Deleted struct {
	// 删除时间，用户软删除
	Deleted time.Time `xorm:"deleted default('2020-02-04 09:55:52')" json:"deleted"`
}
