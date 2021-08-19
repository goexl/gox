package setter

import (
	`time`
)

// Duration 设置字符串
func Duration(setter func(value time.Duration), value time.Duration, null time.Duration) {
	if 0 != value {
		setter(value)
	} else {
		setter(null)
	}
}
