package setter

type boolSetter func(value bool)

// Bool 设置布尔值
func Bool(setter boolSetter, value bool, null bool) {
	if value {
		setter(value)
	} else {
		setter(null)
	}
}

// BoolX 当有值时设置布尔值
func BoolX(setter boolSetter, value bool) {
	if value {
		setter(value)
	}
}
