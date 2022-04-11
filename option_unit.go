package gox

var (
	_            = Unit
	_            = KB
	_            = MB
	_            = GB
	_            = TB
	_            = PB
	_            = EB
	_ sizeOption = (*optionUnit)(nil)
)

type optionUnit struct {
	unit Size
}

// KB 单位
func KB() *optionUnit {
	return Unit(SizeKB)
}

// MB 单位
func MB() *optionUnit {
	return Unit(SizeMB)
}

// GB 单位
func GB() *optionUnit {
	return Unit(SizeGB)
}

// TB 单位
func TB() *optionUnit {
	return Unit(SizeTB)
}

// PB 单位
func PB() *optionUnit {
	return Unit(SizePB)
}

// EB 单位
func EB() *optionUnit {
	return Unit(SizeEB)
}

// Unit 单位
func Unit(unit Size) *optionUnit {
	return &optionUnit{
		unit: unit,
	}
}

func (s *optionUnit) applySize(options *sizeOptions) {
	options.unit = s.unit
}
