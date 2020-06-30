package gox

// 基础用户数据
type BaseUser struct {
	BaseStruct `xorm:"extends"`

	// 用户名
	Username DBString `xorm:"varchar(32) notnull default('') unique(uidx_name)" json:"username" validate:"omitempty,min=1,max=32,email"`
	// 手机号
	// 类似于8617089792784
	Phone DBString `xorm:"varchar(15) notnull default('') unique(uidx_phone)" json:"phone" validate:"omitempty,alphanum,max=15"`
	// 密码
	Pwd string `xorm:"varchar(512) notnull default('')" json:"-"`
}
