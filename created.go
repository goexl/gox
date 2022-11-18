package gox

import (
	"time"
)

// Created 带创建时间模型
type Created struct {
	// 创建时间
	Created time.Time `xorm:"created default('2020-02-04 09:55:52')" json:"created"`
}
