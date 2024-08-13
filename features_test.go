package gox_test

import (
	"testing"

	"github.com/goexl/gox"
)

const (
	Feature1 = iota + 1
	Feature2
	Feature3
	Feature4
)

func TestSwitchEnableAny(t *testing.T) {
	tests := []struct {
		enabled  []gox.Feature
		checks   []gox.Feature
		expected bool
	}{
		{[]gox.Feature{Feature1}, []gox.Feature{Feature1}, true},
		{[]gox.Feature{Feature1}, []gox.Feature{Feature4}, false},
		{[]gox.Feature{Feature1, Feature3}, []gox.Feature{Feature3}, true},
		{[]gox.Feature{Feature1, Feature4}, []gox.Feature{Feature1}, true},
		{[]gox.Feature{Feature1, Feature2, Feature4}, []gox.Feature{Feature1, Feature2}, true},
		{[]gox.Feature{Feature1, Feature2, Feature4}, []gox.Feature{Feature3}, false},
	}
	for index, test := range tests {
		features := gox.NewFeatures()
		features.Enable(test.enabled[0], test.enabled[1:]...)
		got := features.Any(test.checks[0], test.checks[1:]...)
		if test.expected != got {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}

func TestSwitchEnableAll(t *testing.T) {
	tests := []struct {
		enabled  []gox.Feature
		checks   []gox.Feature
		expected bool
	}{
		{[]gox.Feature{Feature1}, []gox.Feature{Feature1}, true},
		{[]gox.Feature{Feature1}, []gox.Feature{Feature4}, false},
		{[]gox.Feature{Feature1, Feature3}, []gox.Feature{Feature3}, true},
		{[]gox.Feature{Feature1, Feature4}, []gox.Feature{Feature1}, true},
		{[]gox.Feature{Feature1, Feature2, Feature4}, []gox.Feature{Feature1, Feature2}, true},
		{[]gox.Feature{Feature1, Feature2, Feature4}, []gox.Feature{Feature3}, false},
	}
	for index, test := range tests {
		features := gox.NewFeatures()
		features.Enable(test.enabled[0], test.enabled[1:]...)
		got := features.All(test.checks[0], test.checks[1:]...)
		if test.expected != got {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}
