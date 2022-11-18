package gox

// SoftDeleted 带软删除功能的数据库模型
type SoftDeleted struct {
	Model   `xorm:"extends"`
	Deleted `xorm:"extends"`
}
