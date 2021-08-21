package setter

type int32Setter func(value int32)

// Int32 设置字符串
func Int32(setter int32Setter, value int32, null int32) {
	if 0 != value {
		setter(value)
	} else {
		setter(null)
	}
}

// Int32X 当有值时设置字符串
func Int32X(setter int32Setter, value int32) {
	if 0 != value {
		setter(value)
	}
}
