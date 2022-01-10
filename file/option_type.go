package file

var (
	_ option     = (*optionType)(nil)
	_ nameOption = (*optionType)(nil)
)

type optionType struct {
	_type _type
}

// Type 配置类型
func Type(_type _type) *optionType {
	return &optionType{
		_type: _type,
	}
}

// File 配置类型为文件
func File() *optionType {
	return Type(TypeFile)
}

func Dir() *optionType {
	return Type(TypeDir)
}

func (t *optionType) apply(options *options) {
	options._type = t._type
}

func (t *optionType) applyName(options *nameOptions) {
	options._type = t._type
}
