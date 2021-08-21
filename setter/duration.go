package setter

import (
	`time`
)

type durationSetter func(value time.Duration)

// Duration 设置时间段
func Duration(setter durationSetter, value time.Duration, null time.Duration) {
	if 0 != value {
		setter(value)
	} else {
		setter(null)
	}
}

// DurationX 设置时间段
func DurationX(setter durationSetter, value time.Duration) {
	if 0 != value {
		setter(value)
	}
}
