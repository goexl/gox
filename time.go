package gox

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	DefaultTimeLayout = "2006-01-02 15:04:05"
)

// Timestamp 时间戳
type Timestamp time.Time

// MarshalJSON 序列化成JSON时调用
func (t Timestamp) MarshalJSON() ([]byte, error) {
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

// ParseTimestamp 从Time对象生成Timestamp
func ParseTimestamp(time time.Time) Timestamp {
	return Timestamp(time)
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
