package gox

import (
	"strconv"
)

type (
	// IdStruct 带序列号的模型
	IdStruct struct {
		// 编号
		Id int64 `xorm:"pk notnull unique index('idx_id') default(0)" json:"id,string"`
	}

	// BaseStruct Xorm基础模型
	BaseStruct struct {
		// 编号
		Id int64 `xorm:"pk notnull unique index('idx_id') default(0)" json:"id,string"`
		// 创建时间
		CreatedAt Timestamp `xorm:"created default('2020-02-04 09:55:52')" json:"createdAt"`
		// 最后更新时间
		UpdatedAt Timestamp `xorm:"updated default('2020-02-04 09:55:52')" json:"updatedAt"`
	}

	// CreateStruct 带创建时间模型
	CreateStruct struct {
		// 编号
		Id int64 `xorm:"pk notnull unique index('idx_id') default(0)" json:"id,string"`
		// 创建时间
		CreatedAt Timestamp `xorm:"created default('2020-02-04 09:55:52')" json:"createdAt"`
	}

	// UpdateStruct 带修改时间模型
	UpdateStruct struct {
		// 编号
		Id int64 `xorm:"pk notnull unique index('idx_id') default(0)" json:"id,string"`
		// 最后更新时间
		UpdatedAt Timestamp `xorm:"updated default('2020-02-04 09:55:52')" json:"updatedAt"`
	}

	// DeleteStruct 软删除模型
	DeleteStruct struct {
		// 编号
		Id int64 `xorm:"pk notnull unique index('idx_id') default(0)" json:"id,string"`
		// 创建时间
		CreatedAt Timestamp `xorm:"created default('2020-02-04 09:55:52')" json:"createdAt"`
		// 最后更新时间
		UpdatedAt Timestamp `xorm:"updated default('2020-02-04 09:55:52')" json:"updatedAt"`
		// 删除时间
		// 用户软删除
		DeletedAt Timestamp `xorm:"deleted default('2020-02-04 09:55:52')" json:"deletedAt"`
	}
)

// IdString Id的字符串形式
func (b IdStruct) IdString() string {
	return strconv.FormatInt(b.Id, 10)
}

// Exists 对象是否存在
func (b BaseStruct) Exists() bool {
	return 0 != b.Id
}

// IdString Id的字符串形式
func (b BaseStruct) IdString() string {
	return strconv.FormatInt(b.Id, 10)
}
