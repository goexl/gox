package gox_test

import (
	"testing"
	"time"

	"github.com/goexl/gox"
)

func TestParseMillisecond(t *testing.T) {
	tests := []struct {
		in       string
		expected time.Time
	}{
		{"0", time.UnixMilli(0)},
		{"1000", time.UnixMilli(1000)},
		{"-1", time.UnixMilli(-1)},
	}
	for index, test := range tests {
		got, err := gox.ParseMillisecond(test.in)
		if nil != err {
			t.Errorf("第%d个测试未通过，原因：%v", index+1, err)
		} else if test.expected != got.Time {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}
