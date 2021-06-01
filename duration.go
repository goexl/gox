package gox

import (
	`encoding/json`
	`errors`
	`time`
)

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
