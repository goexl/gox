package setter

type int8Setter func(value int8)

// Int8 设置字符串
func Int8(setter int8Setter, value int8, null int8) {
	if 0 != value {
		setter(value)
	} else {
		setter(null)
	}
}

// Int8X 当有值时设置字符串
func Int8X(setter int8Setter, value int8) {
	if 0 != value {
		setter(value)
	}
}
