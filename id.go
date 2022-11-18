package gox

import (
	"strconv"
)

// Id 带序列号的模型
type Id struct {
	// Id 编号
	Id int64 `xorm:"pk notnull unique index('idx_id') default(0)" json:"id,string"`
}

func (i *Id) IdString() string {
	return strconv.FormatInt(i.Id, 10)
}

func (i *Id) Exists() bool {
	return 0 != i.Id
}
