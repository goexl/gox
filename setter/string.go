package setter

type stringSetter func(value string)

// String 设置字符串
func String(setter stringSetter, value string, null string) {
	if "" != value {
		setter(value)
	} else {
		setter(null)
	}
}

// StringX 当有值时设置字符串
func StringX(setter stringSetter, value string) {
	if "" != value {
		setter(value)
	}
}
