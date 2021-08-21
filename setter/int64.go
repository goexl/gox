package setter

type int64Setter func(value int64)

// Int64 设置整数值
func Int64(setter int64Setter, value int64, null int64) {
	if 0 != value {
		setter(value)
	} else {
		setter(null)
	}
}

// Int64X 当有值时设置整数值
func Int64X(setter int64Setter, value int64) {
	if 0 != value {
		setter(value)
	}
}
