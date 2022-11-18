package gox

import (
	"strconv"
	"time"
)

type (
	// IdStruct 带序列号的模型
	IdStruct struct {
		// Id 编号
		Id int64 `xorm:"pk notnull unique index('idx_id') default(0)" json:"id,string"`
	}

	// Created 带创建时间模型
	Created struct {
		// 创建时间
		Created time.Time `xorm:"created default('2020-02-04 09:55:52')" json:"created"`
	}

	// Updated 带修改时间模型
	Updated struct {
		// 最后更新时间
		Updated time.Time `xorm:"updated default('2020-02-04 09:55:52')" json:"updated"`
	}

	// Deleted 软删除模型
	Deleted struct {
		// 删除时间，用户软删除
		Deleted time.Time `xorm:"deleted default('2020-02-04 09:55:52')" json:"deleted"`
	}

	// Base 基础数据库模型
	Base struct {
		IdStruct `xorm:"extends"`
		Created  `xorm:"extends"`
		Updated  `xorm:"extends"`
	}

	// SoftDelete 带软删除功能的数据库模型
	SoftDelete struct {
		Base    `xorm:"extends"`
		Deleted `xorm:"extends"`
	}
)

// IdString Id的字符串形式
func (is *IdStruct) IdString() string {
	return strconv.FormatInt(is.Id, 10)
}

// Exists 对象是否存在
func (is *IdStruct) Exists() bool {
	return 0 != is.Id
}
