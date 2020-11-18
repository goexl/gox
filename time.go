package gox

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	// DefaultTimeLayout 默认的时间布局
	DefaultTimeLayout = "2006-01-02 15:04:05"

	// 时间常量
	// Day 一天
	Day = 24 * time.Hour
	// Week 一周
	Week = 7 * Day
	// Month 一个月
	Month = 30 * Day
	// Year 一年
	Year = 365 * Day
)

// Timestamp 时间戳
type Timestamp time.Time

// ParseTimestamp 从Time对象生成Timestamp
func ParseTimestamp(time time.Time) Timestamp {
	return Timestamp(time)
}

// ZeroTimestamp 0值
func ZeroTimestamp() Timestamp {
	return Timestamp(time.Time{})
}

// Now 当前时间
func Now() Timestamp {
	return Timestamp(time.Now())
}

// MarshalJSON 序列化成JSON时调用
func (t Timestamp) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		t = Timestamp(time.Now())
	}

	return []byte(`"` + time.Time(t).Format(DefaultTimeLayout) + `"`), nil
}

// UnmarshalJSON 反序列化成JSON时调用
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := time.ParseInLocation(`"`+DefaultTimeLayout+`"`, string(b), time.Local)
	*t = Timestamp(ts)

	return err
}

// UnmarshalParam 从Echo参数转换
func (t *Timestamp) UnmarshalParam(src string) error {
	ts, err := time.ParseInLocation(DefaultTimeLayout, src, time.Local)
	*t = Timestamp(ts)

	return err
}

// IsZero 是否是零值
func (t Timestamp) IsZero() bool {
	return time.Time(t).IsZero()
}

// Time 取得真实的Time对象
func (t Timestamp) Time() time.Time {
	return time.Time(t)
}

// DefaultTimeLayout 按默认时间布局转换为字符串
func (t Timestamp) DefaultTimeLayout() string {
	return t.Time().Format(DefaultTimeLayout)
}

// DayStart 转换为当天0点0分0秒的Timestamp
func (t Timestamp) DayStart() Timestamp {
	ts, _ := time.ParseInLocation(DefaultTimeLayout, t.DefaultTimeLayout()[:10]+" 00:00:00", time.Local)

	return Timestamp(ts)
}

// DayEnd 转换为当天23点59分59秒的Timestamp
func (t Timestamp) DayEnd() Timestamp {
	ts, _ := time.ParseInLocation(DefaultTimeLayout, t.DefaultTimeLayout()[:10]+" 23:59:59", time.Local)

	return Timestamp(ts)
}

// Duration 弥补标准库不能使用ParseDuration
type Duration time.Duration

// MarshalJSON 编码成JSON
func (d Duration) MarshalJSON() ([]byte, error) {
	td := time.Duration(d)

	return json.Marshal(td.String())
}

// UnmarshalJSON 解析JSON
func (d *Duration) UnmarshalJSON(b []byte) (err error) {
	var v interface{}
	if err = json.Unmarshal(b, &v); err != nil {
		return err
	}

	switch value := v.(type) {
	case float64:
		*d = Duration(time.Duration(value))
	case string:
		var td time.Duration
		td, err = time.ParseDuration(value)
		*d = Duration(td)
	default:
		err = errors.New("无效的Duration格式")
	}

	return
}

// Duration 转换成time.Duration
func (d *Duration) Duration() time.Duration {
	return time.Duration(*d)
}

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

// Weekday 转换为time.Weekday
func (w *Weekday) WeekDay() time.Weekday {
	return time.Weekday(*w)
}
