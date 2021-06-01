package gox

import (
	"time"
)

const (
	// DefaultTimeLayout 默认的时间布局
	DefaultTimeLayout = "2006-01-02 15:04:05"

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

// Format 格式化字符串
func (t *Timestamp) Format() string {
	return time.Time(*t).Format(DefaultTimeLayout)
}

// GobEncode Gob序列化编码
func (t Timestamp) GobEncode() ([]byte, error) {
	return time.Time(t).MarshalBinary()
}

// GobDecode Gob序列化解码
func (t *Timestamp) GobDecode(data []byte) (err error) {
	rt := time.Time(*t)
	if err = rt.UnmarshalBinary(data); nil != err {
		return
	}
	*t = Timestamp(rt)

	return
}

// MarshalJSON 序列化成JSON时调用
func (t Timestamp) MarshalJSON() ([]byte, error) {
	tt := t
	if t.IsZero() {
		tt = Timestamp(time.Now())
	}

	return []byte(`"` + time.Time(tt).Format(DefaultTimeLayout) + `"`), nil
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

func TruncateDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func TruncateUnix(timestamp int64) int64 {
	return TruncateDay(time.Unix(timestamp, 0)).Unix()
}
