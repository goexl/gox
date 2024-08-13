package gox

// Features 一个通用的开关
type Features struct {
	features uint64
	_        CannotCopy
}

// NewFeatures 创建开关
func NewFeatures() *Features {
	return &Features{
		features: 0,
	}
}

func (f *Features) Enable(feature Feature, features ...Feature) *Features {
	f.features |= 1 << feature
	for _, _feature := range features {
		f.features |= 1 << _feature
	}

	return f
}

func (f *Features) Disable(feature Feature, features ...Feature) *Features {
	f.features &^= 1 << feature
	for _, _feature := range features {
		f.features &^= 1 << _feature
	}

	return f
}

func (f *Features) Any(feature Feature, features ...Feature) (enabled bool) {
	all := make([]Feature, 0, len(features)+1)
	all = append(all, feature)
	all = append(all, features...)

	for _, _feature := range all {
		enabled = 0 != f.features&(1<<_feature)
		if enabled {
			return
		}
	}

	return
}

func (f *Features) All(feature Feature, features ...Feature) (enabled bool) {
	all := make([]Feature, 0, len(features)+1)
	all = append(all, feature)
	all = append(all, features...)

	for _, _feature := range all {
		enabled = 0 != f.features&(1<<_feature)
		if !enabled {
			return
		}
	}

	return
}
