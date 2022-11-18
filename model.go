package gox

// Model 基础数据库模型
type Model struct {
	Id      `xorm:"extends"`
	Created `xorm:"extends"`
	Updated `xorm:"extends"`
}
