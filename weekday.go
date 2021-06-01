package gox

import (
	`encoding/json`
	`errors`
	`time`
)

// Weekday 星期几
type Weekday time.Weekday

// WeekdayCN 中文的星期几
var WeekdayCN = []string{
	"周日",
	"周一",
	"周二",
	"周三",
	"周四",
	"周五",
	"周六",
}

// String 返回中文的周几
func (w Weekday) String() string {
	if 0 <= w && 6 >= w {
		return WeekdayCN[w]
	}

	return time.Weekday(w).String()
}

// MarshalJSON 编码成JSON
func (w Weekday) MarshalJSON() ([]byte, error) {
	return json.Marshal(w)
}

// UnmarshalJSON 解析成JSON
func (w *Weekday) UnmarshalJSON(b []byte) (err error) {
	var v interface{}
	if err = json.Unmarshal(b, &v); err != nil {
		return err
	}

	switch value := v.(type) {
	case float64:
		*w = Weekday(time.Weekday(value))
	default:
		err = errors.New("无效的Weekday格式")
	}

	return
}

// WeekDay 转换为time.Weekday
func (w *Weekday) WeekDay() time.Weekday {
	return time.Weekday(*w)
}
