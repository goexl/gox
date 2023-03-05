package gox_test

import (
	"testing"

	"github.com/goexl/gox"
)

func TestCamel(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"a b", "AB"},
		{"a_b", "AB"},
		{"a-b", "AB"},
	}
	for index, test := range tests {
		_case := gox.Case(test.in)
		if got := _case.Camel(gox.CasePositionHead).String(); test.expected != got {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}
