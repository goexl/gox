package gox

// Features 一个通用的开关
type Features struct {
	bitmap *Bitmap
	_      CannotCopy
}

// NewFeatures 创建开关
func NewFeatures() *Features {
	return &Features{
		bitmap: NewBitmap(),
	}
}

func (f *Features) Enable(feature Feature, features ...Feature) *Features {
	f.bitmap.Set(uint(feature))
	for _, _feature := range features {
		f.bitmap.Set(uint(_feature))
	}

	return f
}

func (f *Features) Disable(feature Feature, features ...Feature) *Features {
	f.bitmap.Unset(uint(feature))
	for _, _feature := range features {
		f.bitmap.Unset(uint(_feature))
	}

	return f
}

func (f *Features) Any(feature Feature, features ...Feature) (enabled bool) {
	all := make([]Feature, 0, len(features)+1)
	all = append(all, feature)
	all = append(all, features...)

	for _, _feature := range all {
		enabled = f.bitmap.Contains(uint(_feature))
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
		enabled = f.bitmap.Contains(uint(_feature))
		if !enabled {
			return
		}
	}

	return
}

func (f *Features) MarshalJSON() ([]byte, error) {
	return f.bitmap.MarshalJSON()
}

func (f *Features) UnmarshalJSON(data []byte) error {
	return f.bitmap.UnmarshalJSON(data)
}
