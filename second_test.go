package gox_test

import (
	"testing"
	"time"

	"github.com/goexl/gox"
)

func TestParseSecond(t *testing.T) {
	tests := []struct {
		in       string
		expected time.Time
	}{
		{"0", time.Unix(0, 0)},
		{"1000", time.Unix(1000, 0)},
		{"-1", time.Unix(-1, 0)},
	}
	for index, test := range tests {
		got, err := gox.ParseSecond(test.in)
		if nil != err {
			t.Errorf("第%d个测试未通过，原因：%v", index+1, err)
		} else if test.expected != got.Time {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}
