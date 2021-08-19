package setter

// String 设置字符串
func String(setter func(value string), value string, null string) {
	if "" != value {
		setter(value)
	} else {
		setter(null)
	}
}
