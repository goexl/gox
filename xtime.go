package gox

import (
	`time`
)

func TruncateDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func TruncateUnix(timestamp int64) int64 {
	tm := time.Unix(timestamp, 0)
	return TruncateDay(tm).Unix()
}
