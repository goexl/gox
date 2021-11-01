package gox

type (
	// Field 用于描述一个字段（即键值对）
	Field interface {
		Key
		Value
	}

	// Fields 字段列表
	Fields []Field

	fielder interface {
		// Fields 生成字段列表
		Fields() Fields
	}
)

func (f Fields) Connects(fielder fielder) (new Fields) {
	new = make([]Field, 0, len(f)+len(fielder.Fields()))
	for _, _f := range f {
		new = append(new, _f)
	}
	for _, _f := range fielder.Fields() {
		new = append(new, _f)
	}

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
