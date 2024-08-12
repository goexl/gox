package gox_test

import (
	"testing"

	"github.com/goexl/gox"
)

const (
	Switch1 gox.Switch = 1 << 1
	Switch2 gox.Switch = 1 << 2
	Switch3 gox.Switch = 1 << 3
	Switch4 gox.Switch = 1 << 4
)

func TestSwitch(t *testing.T) {
	tests := []struct {
		in       gox.Switch
		enabled  gox.Switch
		expected bool
	}{
		{Switch1, Switch1, true},
		{Switch1, Switch4, false},
		{Switch1 | Switch3, Switch3, true},
		{Switch1 | Switch4, Switch1, true},
		{Switch1 | Switch2 | Switch4, Switch1 | Switch2, true},
		{Switch1 | Switch2 | Switch4, Switch3, false},
	}
	for index, test := range tests {
		got := test.in.Enabled(test.enabled)
		if test.expected != got {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}
