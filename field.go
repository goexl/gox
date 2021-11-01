package gox

type (
	// Field 用于描述一个字段（即键值对）
	Field interface {
		Key
		Value
	}

	// Fields 字段列表
	Fields []Field

	fields interface {
		// Fields 生成字段列表
		Fields() Fields
	}
)

func (f Fields) Connects(fields ...fields) (new Fields) {
	// 默认创建16个元素，然后再做精简
	new = make([]Field, 0, 16)
	for _, _f := range f {
		new = append(new, _f)
	}
	for _, field := range fields {
		for _, _f := range field.Fields() {
			new = append(new, _f)
		}
	}
	fields = fields[:0]

	return
}

func (f Fields) Connect(field Field) (new Fields) {
	new = make([]Field, 0, len(f)+1)
	for _, _f := range f {
		new = append(new, _f)
	}
	new = append(new, field)

	return
}
