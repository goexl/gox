package gox

// Field 用于描述一个字段（即键值对）
type Field interface {
	Key
	Value
}
