package gox

type (
	// Field 用于描述一个字段（即键值对）
	Field interface {
		Key
		Value
	}

	// Fields 字段列表
	Fields []Field
)

func (f Fields) Connects(fields Fields) (new Fields) {
	new = make([]Field, 0, len(f)+len(fields))
	for _, _f := range f {
		new = append(new, _f)
	}
	for _, _f := range fields {
		new = append(new, _f)
	}

	return
}

func (f Fields) Connect(field Field) (new Fields) {
	new = make([]Field, 0, len(f)+1)
	new = append(new, field)

	return
}
